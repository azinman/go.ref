// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testutil_test

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"v.io/x/ref/lib/v23test"
	_ "v.io/x/ref/runtime/factories/generic"
	"v.io/x/ref/test/expect"
	"v.io/x/ref/test/testutil"
)

func TestFormatLogline(t *testing.T) {
	line, want := testutil.FormatLogLine(2, "test"), "testing.go:.*"
	if ok, err := regexp.MatchString(want, line); !ok || err != nil {
		t.Errorf("got %v, want %v", line, want)
	}
}

func panicHelper(ch chan string) {
	defer func() {
		if r := recover(); r != nil {
			ch <- r.(string)
		}
	}()
	testutil.RandomInt()
}

func TestPanic(t *testing.T) {
	testutil.Rand = nil
	ch := make(chan string)
	go panicHelper(ch)
	str := <-ch
	if got, want := str, "It looks like the singleton random number generator has not been initialized, please call InitRandGenerator."; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func start(t *testing.T, c *v23test.Cmd) *expect.Session {
	s := expect.NewSession(t, c.StdoutPipe(), time.Minute)
	c.Start()
	return s
}

func TestRandSeed(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{SuppressChildOutput: true})
	defer sh.Cleanup()

	s := start(t, sh.Cmd("jiri", "go", "test", "./testdata"))
	s.ExpectRE("FAIL: TestRandSeedInternal.*", 1)
	parts := s.ExpectRE(`Seeded pseudo-random number generator with (\d+)`, -1)
	if len(parts) != 1 || len(parts[0]) != 2 {
		t.Fatalf("failed to match regexp")
	}
	seed := parts[0][1]
	parts = s.ExpectRE(`rand: (\d+)`, -1)
	if len(parts) != 1 || len(parts[0]) != 2 {
		t.Fatalf("failed to match regexp")
	}
	randInt := parts[0][1]

	// Rerun the test, this time with the seed that we want to use.
	cmd := sh.Cmd("jiri", "go", "test", "./testdata")
	cmd.Vars["V23_RNG_SEED"] = seed
	s = start(t, cmd)
	s.ExpectRE("FAIL: TestRandSeedInternal.*", 1)
	s.ExpectRE("Seeded pseudo-random number generator with "+seed, -1)
	s.ExpectRE("rand: "+randInt, 1)
}

func TestFileTreeEqual(t *testing.T) {
	tests := []struct {
		A, B, Err, Debug         string
		FileA, DirA, FileB, DirB *regexp.Regexp
	}{
		{"./testdata/NOEXIST", "./testdata/A", "no such file", "", nil, nil, nil, nil},
		{"./testdata/A", "./testdata/NOEXIST", "no such file", "", nil, nil, nil, nil},

		{"./testdata/A", "./testdata/A", "", "", nil, nil, nil, nil},
		{"./testdata/A/subdir", "./testdata/A/subdir", "", "", nil, nil, nil, nil},

		{"./testdata/A/subdir", "./testdata/SameSubdir/subdir", "", "", nil, nil, nil, nil},
		{"./testdata/SameSubdir/subdir", "./testdata/A/subdir", "", "", nil, nil, nil, nil},

		{"./testdata/A/subdir", "./testdata/DiffSubdirFileName/subdir", "", "relative path doesn't match", nil, nil, nil, nil},
		{"./testdata/DiffSubdirFileName/subdir", "./testdata/A/subdir", "", "relative path doesn't match", nil, nil, nil, nil},

		{"./testdata/A/subdir", "./testdata/DiffSubdirFileBytes/subdir", "", "bytes don't match", nil, nil, nil, nil},
		{"./testdata/DiffSubdirFileBytes/subdir", "./testdata/A/subdir", "", "bytes don't match", nil, nil, nil, nil},

		{"./testdata/A/subdir", "./testdata/ExtraFile/subdir", "", "node count mismatch", nil, nil, nil, nil},
		{"./testdata/ExtraFile/subdir", "./testdata/A/subdir", "", "node count mismatch", nil, nil, nil, nil},

		{"./testdata/A/subdir", "./testdata/ExtraFile/subdir", "", "", nil, nil, regexp.MustCompile(`file3`), nil},
		{"./testdata/ExtraFile/subdir", "./testdata/A/subdir", "", "", regexp.MustCompile(`file3`), nil, nil, nil},

		{"./testdata/A/subdir", "./testdata/ExtraSubdir/subdir", "", "node count mismatch", nil, nil, nil, nil},
		{"./testdata/ExtraSubdir/subdir", "./testdata/A/subdir", "", "node count mismatch", nil, nil, nil, nil},

		{"./testdata/A/subdir", "./testdata/ExtraSubdir/subdir", "", "", nil, nil, regexp.MustCompile(`file3`), regexp.MustCompile(`subdir$`)},
		{"./testdata/ExtraSubdir/subdir", "./testdata/A/subdir", "", "", regexp.MustCompile(`file3`), regexp.MustCompile(`subdir$`), nil, nil},
	}
	for _, test := range tests {
		name := fmt.Sprintf("(%v,%v)", test.A, test.B)
		var debug bytes.Buffer
		opts := testutil.FileTreeOpts{
			Debug: &debug,
			FileA: test.FileA, DirA: test.DirA,
			FileB: test.FileB, DirB: test.DirB,
		}
		equal, err := testutil.FileTreeEqual(test.A, test.B, opts)
		if got, want := err == nil, test.Err == ""; got != want {
			t.Errorf("%v got success %v, want %v", name, got, want)
		}
		if got, want := fmt.Sprint(err), test.Err; err != nil && !strings.Contains(got, want) {
			t.Errorf("%v got error str %v, want substr %v", name, got, want)
		}
		if got, want := equal, test.Err == "" && test.Debug == ""; got != want {
			t.Errorf("%v got %v, want %v", name, got, want)
		}
		if got, want := debug.String(), test.Debug; !strings.Contains(got, want) || got != "" && want == "" {
			t.Errorf("%v got debug %v, want substr %v", name, got, want)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(v23test.Run(m.Run))
}
