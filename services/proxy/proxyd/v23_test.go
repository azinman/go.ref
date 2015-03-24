// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated via go generate.
// DO NOT UPDATE MANUALLY
package main_test

import "fmt"
import "testing"
import "os"

import "v.io/x/ref/test"
import "v.io/x/ref/test/modules"
import "v.io/x/ref/test/v23tests"

func init() {
	modules.RegisterChild("runServer", ``, runServer)
	modules.RegisterChild("runClient", ``, runClient)
}

func TestMain(m *testing.M) {
	test.Init()
	if modules.IsModulesChildProcess() {
		if err := modules.Dispatch(); err != nil {
			fmt.Fprintf(os.Stderr, "modules.Dispatch failed: %v\n", err)
			os.Exit(1)
		}
		return
	}
	cleanup := v23tests.UseSharedBinDir()
	r := m.Run()
	cleanup()
	os.Exit(r)
}

func TestV23Proxyd(t *testing.T) {
	v23tests.RunTest(t, V23TestProxyd)
}