package rt

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"veyron.io/veyron/veyron2"
	"veyron.io/veyron/veyron2/config"
	"veyron.io/veyron/veyron2/i18n"
	"veyron.io/veyron/veyron2/ipc"
	"veyron.io/veyron/veyron2/ipc/stream"
	"veyron.io/veyron/veyron2/naming"
	"veyron.io/veyron/veyron2/rt"

	"veyron.io/veyron/veyron/lib/exec"
	"veyron.io/veyron/veyron/lib/flags"
	_ "veyron.io/veyron/veyron/lib/stats/sysstats"
	"veyron.io/veyron/veyron/runtimes/google/naming/namespace"
	"veyron.io/veyron/veyron2/options"
	"veyron.io/veyron/veyron2/security"
	"veyron.io/veyron/veyron2/vlog"
)

// TODO(caprita): Verrorize this, and maybe move it in the API.
var errCleaningUp = fmt.Errorf("operation rejected: runtime is being cleaned up")

type vrt struct {
	mu sync.Mutex

	profile        veyron2.Profile
	publisher      *config.Publisher
	sm             []stream.Manager // GUARDED_BY(mu)
	ns             naming.Namespace
	signals        chan os.Signal
	principal      security.Principal
	client         ipc.Client
	mgmt           *mgmtImpl
	flags          flags.RuntimeFlags
	reservedDisp   ipc.Dispatcher
	reservedPrefix string
	reservedOpts   []ipc.ServerOpt
	nServers       int  // GUARDED_BY(mu)
	cleaningUp     bool // GUARDED_BY(mu)

	lang    i18n.LangID // Language, from environment variables.
	program string      // Program name, from os.Args[0].
}

var _ veyron2.Runtime = (*vrt)(nil)

var flagsOnce sync.Once
var runtimeFlags *flags.Flags

func init() {
	rt.RegisterRuntime(veyron2.GoogleRuntimeName, New)
	runtimeFlags = flags.CreateAndRegister(flag.CommandLine, flags.Runtime)
}

// Implements veyron2/rt.New
func New(opts ...veyron2.ROpt) (veyron2.Runtime, error) {
	rt := &vrt{mgmt: new(mgmtImpl), lang: i18n.LangIDFromEnv(), program: filepath.Base(os.Args[0])}
	flagsOnce.Do(func() {
		runtimeFlags.Parse(os.Args[1:])
	})
	rt.flags = runtimeFlags.RuntimeFlags()
	for _, o := range opts {
		switch v := o.(type) {
		case options.RuntimePrincipal:
			rt.principal = v.Principal
		case options.Profile:
			rt.profile = v.Profile
		default:
			return nil, fmt.Errorf("option has wrong type %T", o)
		}
	}

	rt.initLogging()
	rt.initSignalHandling()

	if rt.profile == nil {
		return nil, fmt.Errorf("No profile configured!")
	} else {
		vlog.VI(1).Infof("Using profile %q", rt.profile.Name())
	}

	if ns, err := namespace.New(rt, rt.flags.NamespaceRoots...); err != nil {
		return nil, fmt.Errorf("Couldn't create mount table: %v", err)
	} else {
		rt.ns = ns
	}
	vlog.VI(2).Infof("Namespace Roots: %s", rt.ns.Roots())

	// Create the default stream manager.
	if _, err := rt.NewStreamManager(); err != nil {
		return nil, fmt.Errorf("failed to create stream manager: %s", err)
	}

	if err := rt.initSecurity(rt.flags.Credentials); err != nil {
		return nil, fmt.Errorf("failed to init sercurity: %s", err)
	}

	var err error
	if rt.client, err = rt.NewClient(); err != nil {
		return nil, fmt.Errorf("failed to create new client: %s", err)
	}

	handle, err := exec.GetChildHandle()
	switch err {
	case exec.ErrNoVersion:
		// The process has not been started through the veyron exec
		// library. No further action is needed.
	case nil:
		// The process has been started through the veyron exec
		// library. Signal the parent the the child is ready.
		handle.SetReady()
	default:
		return nil, fmt.Errorf("failed to get child handle: %s", err)
	}

	rt.publisher = config.NewPublisher()
	if err := rt.profile.Init(rt, rt.publisher); err != nil {
		return nil, err
	}

	// TODO(caprita, cnicolaou): how is this to be configured?
	// Can it ever be anything other than a localhost/loopback address?
	listenSpec := ipc.ListenSpec{Protocol: "tcp", Address: "127.0.0.1:0"}
	if err := rt.mgmt.initMgmt(rt, listenSpec); err != nil {
		return nil, err
	}

	vlog.VI(2).Infof("rt.Init done")
	return rt, nil
}

func (rt *vrt) Publisher() *config.Publisher {
	return rt.publisher
}

func (rt *vrt) Profile() veyron2.Profile {
	return rt.profile
}

func (rt *vrt) ConfigureReservedName(prefix string, server ipc.Dispatcher, opts ...ipc.ServerOpt) {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	rt.reservedDisp = server
	prefix = strings.TrimLeft(prefix, "_")
	rt.reservedPrefix = naming.ReservedNamePrefix + prefix
	rt.reservedOpts = make([]ipc.ServerOpt, 0, len(opts))
	copy(rt.reservedOpts, opts)
}

func (rt *vrt) Cleanup() {
	rt.mu.Lock()
	if rt.cleaningUp {
		rt.mu.Unlock()
		// TODO(caprita): Should we actually treat extra Cleanups as
		// silent no-ops?  For now, it's better not to, in order to
		// expose programming errors.
		vlog.Errorf("rt.Cleanup done")
		return
	}
	rt.cleaningUp = true
	rt.mu.Unlock()

	// TODO(caprita): Consider shutting down mgmt later in the runtime's
	// shutdown sequence, to capture some of the runtime internal shutdown
	// tasks in the task tracker.
	rt.mgmt.shutdown()
	// It's ok to access rt.sm out of lock below, since a Mutex acts as a
	// barrier in Go and hence we're guaranteed that cleaningUp is true at
	// this point.  The only code that mutates rt.sm is NewStreamManager in
	// ipc.go, which respects the value of cleaningUp.
	//
	// TODO(caprita): It would be cleaner if we do something like this
	// inside the lock (and then iterate over smgrs below):
	//   smgrs := rt.sm
	//   rt.sm = nil
	// However, to make that work we need to prevent NewClient and NewServer
	// from using rt.sm after cleaningUp has been set.  Hence the TODO.
	for _, sm := range rt.sm {
		sm.Shutdown()
	}
	rt.shutdownSignalHandling()
	rt.shutdownLogging()
}
