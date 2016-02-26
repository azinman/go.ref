// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package v23test defines Shell, a wrapper around gosh.Shell that provides
// Vanadium-specific functionality such as credentials management,
// StartRootMountTable, and StartSyncbase.
package v23test

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"syscall"
	"testing"
	"time"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/x/lib/envvar"
	"v.io/x/lib/gosh"
	"v.io/x/ref"
	"v.io/x/ref/test"
	"v.io/x/ref/test/expect"
)

const (
	envBinDir           = "V23_BIN_DIR"
	envChildOutputDir   = "TMPDIR"
	envShellTestProcess = "V23_SHELL_TEST_PROCESS"
	useAgent            = true
)

var envPrincipal string

func init() {
	if useAgent {
		envPrincipal = ref.EnvAgentPath
	} else {
		envPrincipal = ref.EnvCredentials
	}
}

// TODO(sadovsky):
// - Make StartRootMountTable and StartSyncbase fast, and change tests that
//   build no other binaries to be normal (non-integration) tests.
// - Eliminate test.V23Init() and either add v23test.Init() or have v23.Init()
//   check for an env var and perform test-specific configuration.
// - Switch to using the testing package's -test.short flag and eliminate
//   SkipUnlessRunningIntegrationTests, the -v23.tests flag, and the "jiri test"
//   implementation that parses test code to identify integration tests.

// Cmd wraps gosh.Cmd and provides Vanadium-specific functionality.
type Cmd struct {
	*gosh.Cmd
	S  *expect.Session
	sh *Shell
}

// Clone returns a new Cmd with a copy of this Cmd's configuration.
func (c *Cmd) Clone() *Cmd {
	res := &Cmd{Cmd: c.Cmd.Clone(), sh: c.sh}
	initSession(c.sh.tb, res)
	return res
}

// WithCredentials returns a clone of this command, configured to use the given
// credentials.
func (c *Cmd) WithCredentials(cr *Credentials) *Cmd {
	res := c.Clone()
	res.Vars[envPrincipal] = cr.Handle
	return res
}

// Shell wraps gosh.Shell and provides Vanadium-specific functionality.
type Shell struct {
	*gosh.Shell
	Ctx *context.T
	tb  testing.TB
	pm  principalManager
}

var calledInitTestMain = false

// NewShell creates a new Shell. Tests and benchmarks should pass their
// testing.TB; non-tests should pass nil. Ctx is the Vanadium context to use; if
// it's nil, NewShell will call v23.Init to create a context.
func NewShell(tb testing.TB, ctx *context.T) *Shell {
	// Note: gosh only requires gosh.InitMain to be called if the user uses
	// gosh.Shell.FuncCmd, while we require InitTestMain to be called whenever a
	// user calls NewShell with a non-nil testing.TB. We're stricter than gosh
	// because we want to make sure that V23_BIN_DIR gets shared across tests.
	if tb != nil && !calledInitTestMain {
		tb.Fatal("must call v23test.TestMain or v23test.InitTestMain from TestMain")
		return nil
	}

	sh := &Shell{
		Shell: gosh.NewShell(tb),
		tb:    tb,
	}
	if sh.Err != nil {
		return sh
	}
	sh.ChildOutputDir = os.Getenv(envChildOutputDir)

	// Filter out any v23test or credentials-related env vars coming from outside.
	// Note, we intentionally retain envChildOutputDir ("TMPDIR") and
	// envShellTestProcess, as these should be propagated downstream.
	if envChildOutputDir != "TMPDIR" { // sanity check
		panic(envChildOutputDir)
	}
	for _, key := range []string{ref.EnvCredentials, ref.EnvAgentPath} {
		delete(sh.Vars, key)
	}
	if sh.tb != nil {
		sh.Vars[envShellTestProcess] = "1"
	}

	cleanup := true
	defer func() {
		if cleanup {
			sh.Cleanup()
		}
	}()

	if err := sh.initPrincipalManager(); err != nil {
		if _, ok := err.(errAlreadyHandled); !ok {
			sh.HandleError(err)
		}
		return sh
	}
	if err := sh.initCtx(ctx); err != nil {
		if _, ok := err.(errAlreadyHandled); !ok {
			sh.HandleError(err)
		}
		return sh
	}

	cleanup = false
	return sh
}

type errAlreadyHandled struct {
	error
}

// ForkCredentials creates a new Credentials (with a fresh principal) and
// blesses it with the given extensions and no caveats, using this principal's
// default blessings. Additionally, it calls SetDefaultBlessings.
func (sh *Shell) ForkCredentials(extensions ...string) *Credentials {
	sh.Ok()
	creds, err := newCredentials(sh.pm)
	if err != nil {
		sh.HandleError(err)
		return nil
	}
	if err := addDefaultBlessings(v23.GetPrincipal(sh.Ctx), creds.Principal, extensions...); err != nil {
		sh.HandleError(err)
		return nil
	}
	return creds
}

// ForkContext creates a new context with forked credentials.
func (sh *Shell) ForkContext(extensions ...string) *context.T {
	sh.Ok()
	c := sh.ForkCredentials(extensions...)
	if sh.Err != nil {
		return nil
	}
	ctx, err := v23.WithPrincipal(sh.Ctx, c.Principal)
	sh.HandleError(err)
	return ctx
}

// Cleanup cleans up all resources associated with this Shell.
// See gosh.Shell.Cleanup for detailed description.
func (sh *Shell) Cleanup() {
	// Run sh.Shell.Cleanup even if DebugSystemShell panics.
	defer sh.Shell.Cleanup()
	if sh.tb != nil && sh.tb.Failed() && test.IntegrationTestsDebugShellOnError {
		sh.DebugSystemShell()
	}
}

// BuildGoPkg compiles a Go package using the "go build" command and writes the
// resulting binary to V23_BIN_DIR, or to the -o flag location if specified. If
// -o is relative, it is interpreted as relative to V23_BIN_DIR. If the binary
// already exists at the target location, it is not rebuilt. Returns the
// absolute path to the binary.
func BuildGoPkg(sh *Shell, pkg string, flags ...string) string {
	sh.Ok()
	binDir := os.Getenv(envBinDir)
	if binDir == "" {
		sh.HandleError(errors.New("v23test: missing V23_BIN_DIR"))
		return ""
	}
	return gosh.BuildGoPkg(sh.Shell, binDir, pkg, flags...)
}

// InitTestMain is intended for developers with complex setup or teardown
// requirements; developers should generally use TestMain. InitTestMain must be
// called early on in TestMain, before m.Run is called. It calls gosh.InitMain
// and, if V23_BIN_DIR is not set, sets it to a new temporary directory. Returns
// a cleanup function that should be called after m.Run but before os.Exit.
func InitTestMain() func() {
	gosh.InitMain()
	calledInitTestMain = true
	cleanup := func() {}
	if os.Getenv(envBinDir) == "" {
		dir, err := ioutil.TempDir("", "bin-")
		if err != nil {
			panic(err)
		}
		cleanup = func() {
			os.RemoveAll(dir)
			os.Unsetenv(envBinDir)
		}
		if err := os.Setenv(envBinDir, dir); err != nil {
			cleanup()
			panic(err)
		}
	}
	return cleanup
}

// TestMain calls flag.Parse and does some v23test/gosh setup work, then calls
// os.Exit(m.Run()). Developers with complex setup or teardown requirements may
// need to use InitTestMain instead.
func TestMain(m *testing.M) {
	flag.Parse()
	var code int
	func() {
		defer InitTestMain()()
		code = m.Run()
	}()
	os.Exit(code)
}

// SkipUnlessRunningIntegrationTests should be called first thing inside of the
// test function body of an integration test. It causes this test to be skipped
// unless integration tests are being run, i.e. unless the -v23.tests flag is
// set.
// TODO(sadovsky): Switch to using -test.short. See TODO above.
func SkipUnlessRunningIntegrationTests(tb testing.TB) {
	// Note: The "jiri test run vanadium-integration-test" command looks for test
	// function names that start with "TestV23", and runs "go test" for only those
	// packages containing at least one such test. That's how it avoids passing
	// the -v23.tests flag to packages for which the flag is not registered.
	name, err := callerName()
	if err != nil {
		tb.Fatal(err)
	}
	if !strings.HasPrefix(name, "TestV23") {
		tb.Fatalf("integration test names must start with \"TestV23\": %s", name)
		return
	}
	if !test.IntegrationTestsEnabled {
		tb.SkipNow()
	}
}

////////////////////////////////////////////////////////////////////////////////
// Methods for starting subprocesses

func initSession(tb testing.TB, c *Cmd) {
	c.S = expect.NewSession(tb, c.StdoutPipe(), time.Minute)
	c.S.SetVerbosity(testing.Verbose())
}

func newCmd(sh *Shell, c *gosh.Cmd) *Cmd {
	res := &Cmd{Cmd: c, sh: sh}
	initSession(sh.tb, res)
	res.Vars[envPrincipal] = sh.ForkCredentials("child").Handle
	return res
}

// Cmd returns a Cmd for an invocation of the named program. The given arguments
// are passed to the child as command-line arguments.
func (sh *Shell) Cmd(name string, args ...string) *Cmd {
	c := sh.Shell.Cmd(name, args...)
	if sh.Err != nil {
		return nil
	}
	return newCmd(sh, c)
}

// FuncCmd returns a Cmd for an invocation of the given registered Func. The
// given arguments are gob-encoded in the parent process, then gob-decoded in
// the child and passed to the Func. To specify command-line arguments for the
// child invocation, append to the returned Cmd's Args.
func (sh *Shell) FuncCmd(f *gosh.Func, args ...interface{}) *Cmd {
	c := sh.Shell.FuncCmd(f, args...)
	if sh.Err != nil {
		return nil
	}
	return newCmd(sh, c)
}

////////////////////////////////////////////////////////////////////////////////
// DebugSystemShell

// DebugSystemShell drops the user into a debug system shell (e.g. bash) that
// includes all environment variables from sh. If there is no controlling TTY,
// DebugSystemShell does nothing.
func (sh *Shell) DebugSystemShell() {
	cwd, err := os.Getwd()
	if err != nil {
		sh.tb.Fatalf("Getwd() failed: %v\n", err)
		return
	}

	// Transfer stdin, stdout, and stderr to the new process, and set target
	// directory for the system shell to start in.
	devtty := "/dev/tty"
	fd, err := syscall.Open(devtty, syscall.O_RDWR, 0)
	if err != nil {
		sh.tb.Logf("WARNING: Open(%q) failed: %v\n", devtty, err)
		return
	}

	file := os.NewFile(uintptr(fd), devtty)
	attr := os.ProcAttr{
		Files: []*os.File{file, file, file},
		Dir:   cwd,
	}
	env := envvar.MergeMaps(envvar.SliceToMap(os.Environ()), sh.Vars)
	env[envPrincipal] = sh.ForkCredentials("debug").Handle
	attr.Env = envvar.MapToSlice(env)

	write := func(s string) {
		if _, err := file.WriteString(s); err != nil {
			sh.tb.Fatalf("WriteString(%q) failed: %v\n", s, err)
			return
		}
	}

	write(">> Starting a new interactive shell\n")
	write(">> Hit Ctrl-D to resume the test\n")

	shellPath := "/bin/sh"
	if shellPathFromEnv := os.Getenv("SHELL"); shellPathFromEnv != "" {
		shellPath = shellPathFromEnv
	}
	proc, err := os.StartProcess(shellPath, []string{}, &attr)
	if err != nil {
		sh.tb.Fatalf("StartProcess(%q) failed: %v\n", shellPath, err)
		return
	}

	// Wait until the user exits the shell.
	state, err := proc.Wait()
	if err != nil {
		sh.tb.Fatalf("Wait() failed: %v\n", err)
		return
	}

	write(fmt.Sprintf(">> Exited shell: %s\n", state.String()))
}

////////////////////////////////////////////////////////////////////////////////
// Internals

func callerName() (string, error) {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "", errors.New("runtime.Caller failed")
	}
	name := runtime.FuncForPC(pc).Name()
	// Strip package path.
	return name[strings.LastIndex(name, ".")+1:], nil
}

func (sh *Shell) initPrincipalManager() error {
	dir := sh.MakeTempDir()
	if sh.Err != nil {
		return errAlreadyHandled{sh.Err}
	}
	var pm principalManager
	if useAgent {
		var err error
		if pm, err = newAgentPrincipalManager(dir); err != nil {
			return err
		}
	} else {
		pm = newFilesystemPrincipalManager(dir)
	}
	sh.pm = pm
	return nil
}

func (sh *Shell) initCtx(ctx *context.T) error {
	if ctx == nil {
		var shutdown func()
		ctx, shutdown = v23.Init()
		if sh.AddCleanupHandler(shutdown); sh.Err != nil {
			return errAlreadyHandled{sh.Err}
		}
		if sh.tb != nil {
			creds, err := newRootCredentials(sh.pm)
			if err != nil {
				return err
			}
			if ctx, err = v23.WithPrincipal(ctx, creds.Principal); err != nil {
				return err
			}
		}
	}
	if sh.tb != nil {
		ctx = v23.WithListenSpec(ctx, rpc.ListenSpec{Addrs: rpc.ListenAddrs{{Protocol: "tcp", Address: "127.0.0.1:0"}}})
	}
	sh.Ctx = ctx
	return nil
}
