// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"testing"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/x/lib/cmdline"
	"v.io/x/ref/services/ben"
	"v.io/x/ref/test"
)

func TestGo(t *testing.T) {
	// go test -bench . testdata_test.go
	goTestOutput := runGoTest(t)
	ctx, shutdown := test.V23Init()
	defer shutdown()

	// Start an archival server
	var archiver archiver
	_, server, err := v23.WithNewServer(ctx, "", &archiver, security.AllowEveryone())
	if err != nil {
		t.Fatal(err)
	}

	// Pipe goTestOutput through the "benup" command
	env := cmdline.EnvFromOS()
	env.Stdin = bytes.NewBufferString(goTestOutput)
	out := bytes.NewBuffer(nil)
	env.Stdout = out
	flagArchiver = server.Status().Endpoints[0].Name()
	if err := runUpload(ctx, env, nil); err != nil {
		t.Fatal(err)
	}
	gotS, gotC, gotR := archiver.lastCall()
	// Some sanity checks on the Scenario
	if got, want := gotS.Cpu.Architecture, runtime.GOARCH; got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
	if got, want := gotS.Os.Name, runtime.GOOS; got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
	// And the state of the code
	if len(gotC) == 0 {
		t.Errorf("SourceCode not detected?")
	}
	// All the runs
	// Since goTestOutput was generated by running "go test" directly on a
	// file, the "package name" is set to "command-line-arguments" (by the
	// "go" tool).
	const (
		bmGood       = "command-line-arguments.BenchmarkGood"
		bmThroughput = "command-line-arguments.BenchmarkThroughput"
		bmAllocs     = "command-line-arguments.BenchmarkAllocs"
		bmAll        = "command-line-arguments.BenchmarkAllMetrics"
	)
	var gotNames []string
	for _, run := range gotR {
		gotNames = append(gotNames, run.Name)
		// Some fields should be set for all benchmarks
		if run.Iterations == 0 {
			t.Errorf("Got 0 iterations for %q", run.Name)
		}
		if run.NanoSecsPerOp == 0 {
			t.Errorf("Got 0 NanoSecsPerOp for %q", run.Name)
		}
		if run.Parallelism == 0 {
			t.Errorf("Got 0 parallelism for %q", run.Name)
		}
		switch run.Name {
		case bmGood:
			// No other metrics set
		case bmThroughput:
			if run.MegaBytesPerSec == 0 {
				t.Errorf("Got 0 MegaBytesPerSec for %q", run.Name)
			}
		case bmAllocs:
			if run.AllocsPerOp == 0 {
				t.Errorf("Got 0 AllocsPerOp for %q", run.Name)
			}
			if run.AllocedBytesPerOp == 0 {
				t.Errorf("Got 0 AllocedBytesPerOp for %q", run.Name)
			}
		case bmAll:
			if run.MegaBytesPerSec == 0 {
				t.Errorf("Got 0 MegaBytesPerSec for %q", run.Name)
			}
			if run.AllocsPerOp == 0 {
				t.Errorf("Got 0 AllocsPerOp for %q", run.Name)
			}
			if run.AllocedBytesPerOp == 0 {
				t.Errorf("Got 0 AllocedBytesPerOp for %q", run.Name)
			}
		}
	}
	sort.Strings(gotNames)
	if got, want := gotNames, []string{bmAll, bmAllocs, bmGood, bmThroughput}; !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
	// And benup output should said that it uploaded 4 benchmarks.
	if want := "Uploaded 4 benchmark(s)"; !strings.Contains(out.String(), want) {
		t.Errorf("Did not find %q in benup output:\n%v", want, out.String())
	}
	if want := "someprotocol://someurl"; !strings.Contains(out.String(), want) {
		t.Errorf("Did not find %q in benup output:\n%v", want, out.String())
	}
}

func TestNoBenchmarks(t *testing.T) {
	// When the input does not contain any results, no RPCs will be send.
	tests := []string{
		"",
		"There is nothing here",
		`
BENDROIDOS_VERSION="Foo bar baz"
Gotcha!`,
	}
	ctx, shutdown := test.V23Init()
	defer shutdown()
	env := cmdline.EnvFromOS()
	for idx, test := range tests {
		env.Stdin = bytes.NewBufferString(test)
		out := bytes.NewBuffer(nil)
		env.Stdout = out
		if err := runUpload(ctx, env, nil); err != nil {
			t.Errorf("%v -- with input #%d %q", err, idx, test)
		}
		if want := "No benchmarks found to upload"; !strings.Contains(out.String(), want) {
			t.Errorf("Did not find %q in output for input #%d (output = %q)", want, idx, out.String())
		}
	}
}

func TestDetectScenario(t *testing.T) {
	scn := detectScenario()
	testAllFieldsSet(t, scn.Cpu)
	testAllFieldsSet(t, scn.Os)
}

func runGoTest(t *testing.T) string {
	// Generates the output of "go test" on testdata_test.go
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("failed to extract filename of current source file")
	}
	gobin, err := exec.LookPath("go")
	if err != nil {
		t.Fatalf("go compiler not found in PATH: %v", err)
	}
	versionBytes, err := exec.Command(gobin, "version").CombinedOutput()
	if err != nil {
		t.Fatalf("%v version failed: %v (%s)", gobin, err, versionBytes)
	}
	t.Logf("Go compiler: %s", versionBytes)
	cmd := exec.Command(gobin, "test", "-v", "-bench", ".", filepath.Join(filepath.Dir(file), "testdata_test.go"))
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%v failed: %v (%s)", cmd.Args, err, output)
	}
	return string(output)
}

type archiver struct {
	mu   sync.Mutex
	scn  ben.Scenario
	code string
	runs []ben.Run
}

func (a *archiver) Archive(ctx *context.T, call rpc.ServerCall, scenario ben.Scenario, code string, runs []ben.Run) (string, error) {
	a.mu.Lock()
	a.scn = scenario
	a.code = code
	a.runs = runs
	a.mu.Unlock()
	return "someprotocol://someurl", nil
}

func (a *archiver) lastCall() (ben.Scenario, string, []ben.Run) {
	a.mu.Lock()
	scn, code, runs := a.scn, a.code, a.runs
	a.scn = ben.Scenario{}
	a.code = ""
	a.runs = nil
	a.mu.Unlock()
	return scn, code, runs
}

func testAllFieldsSet(t *testing.T, value interface{}) {
	rv := reflect.ValueOf(value)
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		v := f.Interface()
		z := reflect.Zero(f.Type()).Interface()
		if reflect.DeepEqual(v, z) {
			t.Errorf("%s not set", rv.Type().Field(i).Name)
		}
	}
}