package mounttable

import (
	"errors"
	"net"
	"strconv"
	"strings"
	"time"

	"v.io/x/lib/netconfig"
	"v.io/x/ref/lib/glob"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/services/mounttable"
	"v.io/v23/services/security/access"
	vdltime "v.io/v23/vdlroot/time"
	"v.io/v23/verror"
	"v.io/x/lib/vlog"

	mdns "github.com/presotto/go-mdns-sd"
)

const addressPrefix = "address:"

// neighborhood defines a set of machines on the same multicast media.
type neighborhood struct {
	mdns   *mdns.MDNS
	nelems int
	nw     netconfig.NetConfigWatcher
}

var _ rpc.Dispatcher = (*neighborhood)(nil)

type neighborhoodService struct {
	name  string
	elems []string
	nh    *neighborhood
}

func getPort(address string) uint16 {
	epAddr, _ := naming.SplitAddressName(address)

	ep, err := v23.NewEndpoint(epAddr)
	if err != nil {
		return 0
	}
	addr := ep.Addr()
	if addr == nil {
		return 0
	}
	switch addr.Network() {
	case "tcp", "tcp4", "tcp6", "ws", "ws4", "ws6", "wsh", "wsh4", "wsh6":
	default:
		return 0
	}
	_, pstr, err := net.SplitHostPort(addr.String())
	if err != nil {
		return 0
	}
	port, err := strconv.ParseUint(pstr, 10, 16)
	if err != nil || port == 0 {
		return 0
	}
	return uint16(port)
}

func newNeighborhood(host string, addresses []string, loopback bool) (*neighborhood, error) {
	// Create the TXT contents with addresses to announce. Also pick up a port number.
	var txt []string
	var port uint16
	for _, addr := range addresses {
		txt = append(txt, addressPrefix+addr)
		if port == 0 {
			port = getPort(addr)
		}
	}
	if txt == nil {
		return nil, errors.New("neighborhood passed no useful addresses")
	}
	if port == 0 {
		return nil, errors.New("neighborhood couldn't determine a port to use")
	}

	// Start up MDNS, subscribe to the vanadium service, and add us as a vanadium service provider.
	mdns, err := mdns.NewMDNS(host, "", "", loopback, false)
	if err != nil {
		vlog.Errorf("mdns startup failed: %s", err)
		return nil, err
	}
	vlog.VI(2).Infof("listening for service vanadium on port %d", port)
	mdns.SubscribeToService("vanadium")
	mdns.AddService("vanadium", "", port, txt...)

	nh := &neighborhood{
		mdns: mdns,
	}

	// Watch the network configuration so that we can make MDNS reattach to
	// interfaces when the network changes.
	nh.nw, err = netconfig.NewNetConfigWatcher()
	if err != nil {
		vlog.Errorf("nighborhood can't watch network: %s", err)
		return nh, nil
	}
	go func() {
		if _, ok := <-nh.nw.Channel(); !ok {
			return
		}
		if _, err := nh.mdns.ScanInterfaces(); err != nil {
			vlog.Errorf("nighborhood can't scan interfaces: %s", err)
		}
	}()

	return nh, nil
}

// NewLoopbackNeighborhoodDispatcher creates a new instance of a dispatcher for
// a neighborhood service provider on loopback interfaces (meant for testing).
func NewLoopbackNeighborhoodDispatcher(host string, addresses ...string) (rpc.Dispatcher, error) {
	return newNeighborhood(host, addresses, true)
}

// NewNeighborhoodDispatcher creates a new instance of a dispatcher for a
// neighborhood service provider.
func NewNeighborhoodDispatcher(host string, addresses ...string) (rpc.Dispatcher, error) {
	return newNeighborhood(host, addresses, false)
}

// Lookup implements rpc.Dispatcher.Lookup.
func (nh *neighborhood) Lookup(name string) (interface{}, security.Authorizer, error) {
	vlog.VI(1).Infof("*********************LookupServer '%s'\n", name)
	elems := strings.Split(name, "/")[nh.nelems:]
	if name == "" {
		elems = nil
	}
	ns := &neighborhoodService{
		name:  name,
		elems: elems,
		nh:    nh,
	}
	return mounttable.MountTableServer(ns), nh, nil
}

func (nh *neighborhood) Authorize(*context.T) error {
	// TODO(rthellend): Figure out whether it's OK to accept all requests
	// unconditionally.
	return nil
}

// Stop performs cleanup.
func (nh *neighborhood) Stop() {
	if nh.nw != nil {
		nh.nw.Stop()
	}
	nh.mdns.Stop()
}

// neighbor returns the MountedServers for a particular neighbor.
func (nh *neighborhood) neighbor(instance string) []naming.MountedServer {
	now := time.Now()
	var reply []naming.MountedServer
	si := nh.mdns.ResolveInstance(instance, "vanadium")

	// Use a map to dedup any addresses seen
	addrMap := make(map[string]vdltime.Deadline)

	// Look for any TXT records with addresses.
	for _, rr := range si.TxtRRs {
		for _, s := range rr.Txt {
			if !strings.HasPrefix(s, addressPrefix) {
				continue
			}
			addr := s[len(addressPrefix):]
			ttl := time.Second * time.Duration(rr.Header().Ttl)
			addrMap[addr] = vdltime.Deadline{now.Add(ttl)}
		}
	}
	for addr, deadline := range addrMap {
		reply = append(reply, naming.MountedServer{addr, nil, deadline})
	}

	if reply != nil {
		return reply
	}

	// If we didn't get any direct endpoints, make some up from the target and port.
	// TODO(rthellend): Do we need the code below? If we haven't received
	// any results at this point, it would seem to indicate a bug somewhere
	// because NeighborhoodServer won't start without at least one address.
	for _, rr := range si.SrvRRs {
		ips, ttl := nh.mdns.ResolveAddress(rr.Target)
		for _, ip := range ips {
			addr := net.JoinHostPort(ip.String(), strconv.Itoa(int(rr.Port)))
			ep := naming.FormatEndpoint("tcp", addr)
			deadline := vdltime.Deadline{now.Add(time.Second * time.Duration(ttl))}
			reply = append(reply, naming.MountedServer{naming.JoinAddressName(ep, ""), nil, deadline})
		}
	}
	return reply
}

// neighbors returns all neighbors and their MountedServer structs.
func (nh *neighborhood) neighbors() map[string][]naming.MountedServer {
	neighbors := make(map[string][]naming.MountedServer, 0)
	members := nh.mdns.ServiceDiscovery("vanadium")
	for _, m := range members {
		if neighbor := nh.neighbor(m.Name); neighbor != nil {
			neighbors[m.Name] = neighbor
		}
	}
	vlog.VI(2).Infof("members %v neighbors %v", members, neighbors)
	return neighbors
}

// ResolveStepX implements ResolveStepX
func (ns *neighborhoodService) ResolveStepX(call rpc.ServerCall) (entry naming.MountEntry, err error) {
	return ns.ResolveStep(call)
}

// ResolveStep implements ResolveStep
func (ns *neighborhoodService) ResolveStep(call rpc.ServerCall) (entry naming.MountEntry, err error) {
	nh := ns.nh
	vlog.VI(2).Infof("ResolveStep %v\n", ns.elems)
	if len(ns.elems) == 0 {
		//nothing can be mounted at the root
		err = verror.New(naming.ErrNoSuchNameRoot, call.Context(), ns.elems)
		return
	}

	// We can only resolve the first element and it always refers to a mount table (for now).
	neighbor := nh.neighbor(ns.elems[0])
	if neighbor == nil {
		err = verror.New(naming.ErrNoSuchName, call.Context(), ns.elems)
		entry.Name = ns.name
		return
	}
	entry.ServesMountTable = true
	entry.Name = naming.Join(ns.elems[1:]...)
	entry.Servers = neighbor
	return
}

// Mount not implemented.
func (ns *neighborhoodService) Mount(call rpc.ServerCall, server string, ttlsecs uint32, opts naming.MountFlag) error {
	return ns.MountX(call, server, nil, ttlsecs, opts)
}
func (ns *neighborhoodService) MountX(_ rpc.ServerCall, _ string, _ []security.BlessingPattern, _ uint32, _ naming.MountFlag) error {
	return errors.New("this server does not implement Mount")
}

// Unmount not implemented.
func (*neighborhoodService) Unmount(_ rpc.ServerCall, _ string) error {
	return errors.New("this server does not implement Unmount")
}

// Delete not implemented.
func (*neighborhoodService) Delete(_ rpc.ServerCall, _ bool) error {
	return errors.New("this server does not implement Delete")
}

// Glob__ implements rpc.AllGlobber
func (ns *neighborhoodService) Glob__(call rpc.ServerCall, pattern string) (<-chan naming.GlobReply, error) {
	g, err := glob.Parse(pattern)
	if err != nil {
		return nil, err
	}

	// return all neighbors that match the first element of the pattern.
	nh := ns.nh

	switch len(ns.elems) {
	case 0:
		ch := make(chan naming.GlobReply)
		go func() {
			defer close(ch)
			for k, n := range nh.neighbors() {
				if ok, _, _ := g.MatchInitialSegment(k); !ok {
					continue
				}
				ch <- naming.GlobReplyEntry{naming.MountEntry{Name: k, Servers: n, ServesMountTable: true}}
			}
		}()
		return ch, nil
	case 1:
		neighbor := nh.neighbor(ns.elems[0])
		if neighbor == nil {
			return nil, verror.New(naming.ErrNoSuchName, call.Context(), ns.elems[0])
		}
		ch := make(chan naming.GlobReply, 1)
		ch <- naming.GlobReplyEntry{naming.MountEntry{Name: "", Servers: neighbor, ServesMountTable: true}}
		close(ch)
		return ch, nil
	default:
		return nil, verror.New(naming.ErrNoSuchName, call.Context(), ns.elems)
	}
}

func (*neighborhoodService) SetPermissions(call rpc.ServerCall, acl access.Permissions, etag string) error {
	return errors.New("this server does not implement SetPermissions")
}

func (*neighborhoodService) GetPermissions(call rpc.ServerCall) (acl access.Permissions, etag string, err error) {
	return nil, "", nil
}
