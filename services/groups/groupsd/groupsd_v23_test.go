// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/security/access"
	"v.io/v23/services/groups"
	"v.io/v23/verror"
	"v.io/x/lib/gosh"
	"v.io/x/lib/set"
	"v.io/x/ref/lib/signals"
	"v.io/x/ref/lib/v23test"
	"v.io/x/ref/services/groups/groupsd/testdata/kvstore"
	"v.io/x/ref/test/expect"
)

type relateResult struct {
	Remainder      map[string]struct{}
	Approximations []groups.Approximation
	Version        string
}

// TestV23GroupServerIntegration tests the integration between the "groups"
// command-line client and the "groupsd" server.
func TestV23GroupServerIntegration(t *testing.T) {
	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()
	sh.StartRootMountTable()

	// Build binaries for the client and server.
	var (
		clientBin  = sh.JiriBuildGoPkg("v.io/x/ref/services/groups/groups")
		serverBin  = sh.JiriBuildGoPkg("v.io/x/ref/services/groups/groupsd", "-tags=leveldb")
		serverName = "groups-server"
		groupA     = naming.Join(serverName, "groupA")
		groupB     = naming.Join(serverName, "groupB")
	)

	// Start the groups server.
	sh.Cmd(serverBin, "-name="+serverName, "-v23.tcp.address=127.0.0.1:0").Start()

	// Create a couple of groups.
	sh.Cmd(clientBin, "create", groupA).Run()
	sh.Cmd(clientBin, "create", groupB, "a", "a:b").Run()

	// Add a couple of blessing patterns.
	sh.Cmd(clientBin, "add", groupA, "<grp:groups-server/groupB>").Run()
	sh.Cmd(clientBin, "add", groupA, "a").Run()
	sh.Cmd(clientBin, "add", groupB, "a:b:c").Run()

	// Remove a blessing pattern.
	sh.Cmd(clientBin, "remove", groupB, "a").Run()

	// Test simple group resolution.
	{
		stdout, stderr := sh.Cmd(clientBin, "relate", groupB, "a:b:c:d").Output()

		var got relateResult
		if err := json.Unmarshal([]byte(stdout), &got); err != nil {
			t.Fatalf("Unmarshal(%v) failed: %v", stdout, err)
		}
		want := relateResult{
			Remainder:      set.String.FromSlice([]string{"c:d", "d"}),
			Approximations: nil,
			Version:        "2",
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
		if got, want := stderr, ""; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	// Test recursive group resolution.
	{
		stdout, stderr := sh.Cmd(clientBin, "relate", groupA, "a:b:c:d").Output()

		var got relateResult
		if err := json.Unmarshal([]byte(stdout), &got); err != nil {
			t.Fatalf("Unmarshal(%v) failed: %v", stdout, err)
		}
		want := relateResult{
			Remainder:      set.String.FromSlice([]string{"b:c:d", "c:d", "d"}),
			Approximations: nil,
			Version:        "2",
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
		if got, want := stderr, ""; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	// Test group resolution failure. Note that under-approximation is
	// used as the default to handle resolution failures.
	{
		sh.Cmd(clientBin, "add", groupB, "<grp:groups-server/groupC>").Run()

		stdout, stderr := sh.Cmd(clientBin, "relate", groupB, "a:b:c:d").Output()

		var got relateResult
		if err := json.Unmarshal([]byte(stdout), &got); err != nil {
			t.Fatalf("Unmarshal(%v) failed: %v", stdout, err)
		}
		want := relateResult{
			Remainder: set.String.FromSlice([]string{"c:d", "d"}),
			Approximations: []groups.Approximation{
				groups.Approximation{
					Reason:  "v.io/v23/verror.NoExist",
					Details: `groupsd:"groupC".Relate: Does not exist: groupC`,
				},
			},
			Version: "3",
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
		if got, want := stderr, ""; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	// Delete the groups.
	sh.Cmd(clientBin, "delete", groupA).Run()
	sh.Cmd(clientBin, "delete", groupB).Run()
}

// store implements the kvstore.Store interface.
type store map[string]string

func (s store) Get(ctx *context.T, call rpc.ServerCall, key string) (string, error) {
	return s[key], nil
}

func (s store) Set(ctx *context.T, call rpc.ServerCall, key string, value string) error {
	s[key] = value
	return nil
}

const (
	kvServerName = "key-value-store"
	getFailed    = "GET FAILED"
	getOK        = "GET OK"
	setFailed    = "SET FAILED"
	setOK        = "SET OK"
)

var runServer = gosh.Register("server", func() error {
	ctx, shutdown := v23.Init()
	defer shutdown()
	// Use a shorter timeout to reduce the test overall runtime as the
	// permission authorizer will attempt to connect to a non-existing
	// groups server at some point in the test.
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	authorizer, err := groups.PermissionsAuthorizer(access.Permissions{
		"Read":  access.AccessList{In: []security.BlessingPattern{"<grp:groups-server/readers>"}},
		"Write": access.AccessList{In: []security.BlessingPattern{"<grp:groups-server/writers>"}},
	}, access.TypicalTagType())
	if err != nil {
		return err
	}
	if _, _, err := v23.WithNewServer(ctx, kvServerName, kvstore.StoreServer(&store{}), authorizer); err != nil {
		return err
	}
	<-signals.ShutdownOnSignals(ctx)
	return nil
})

var runClient = gosh.Register("client", func(command string, args ...string) error {
	ctx, shutdown := v23.Init()
	defer shutdown()

	client := kvstore.StoreClient(kvServerName)
	switch command {
	case "get":
		if got, want := len(args), 1; got != want {
			return fmt.Errorf("unexpected number of arguments: got %v, want %v", got, want)
		}
		key := args[0]
		value, err := client.Get(ctx, key)
		if err != nil {
			fmt.Printf("%v %v\n", getFailed, verror.ErrorID(err))
		} else {
			fmt.Printf("%v %v\n", getOK, value)
		}
	case "set":
		if got, want := len(args), 2; got != want {
			return fmt.Errorf("unexpected number of arguments: got %v, want %v", got, want)
		}
		key, value := args[0], args[1]
		if err := client.Set(ctx, key, value); err != nil {
			fmt.Printf("%v %v\n", setFailed, verror.ErrorID(err))
		} else {
			fmt.Printf("%v\n", setOK)
		}
	}
	return nil
})

func startClient(t *testing.T, sh *v23test.Shell, name string, args ...interface{}) *expect.Session {
	cmd := sh.Fn(runClient, args...).WithCredentials(sh.ForkCredentials(name))
	session := expect.NewSession(t, cmd.StdoutPipe(), time.Minute)
	cmd.Start()
	return session
}

// TestV23GroupServerAuthorization uses an instance of the KeyValueStore server
// with an groups-based authorizer to test the group server implementation.
func TestV23GroupServerAuthorization(t *testing.T) {
	// TODO(sadovsky): Figure out why this test fails on Jenkins.
	t.Skip("Passes locally but fails on Jenkins, see presubmit for v.io/c/18566")

	sh := v23test.NewShell(t, v23test.Opts{Large: true})
	defer sh.Cleanup()
	sh.StartRootMountTable()

	// Build binaries for the groups client and server.
	var (
		clientBin  = sh.JiriBuildGoPkg("v.io/x/ref/services/groups/groups")
		serverBin  = sh.JiriBuildGoPkg("v.io/x/ref/services/groups/groupsd")
		serverName = "groups-server"
		readers    = naming.Join(serverName, "readers")
		writers    = naming.Join(serverName, "writers")
	)

	// Start the groups server.
	server := sh.Cmd(serverBin, "-name="+serverName, "-v23.tcp.address=127.0.0.1:0")
	server.Start()

	// Create a couple of groups. The <readers> and <writers> groups
	// identify blessings that can be used to read from and write to the
	// key value store server respectively.
	sh.Cmd(clientBin, "create", readers, "root:alice", "root:bob").Run()
	sh.Cmd(clientBin, "create", writers, "root:alice").Run()

	// Start an instance of the key value store server.
	sh.Fn(runServer).Start()

	// Test that alice can write.
	startClient(t, sh, "alice", "set", "foo", "bar").Expect(setOK)
	// Test that alice can read.
	startClient(t, sh, "alice", "get", "foo").Expectf("%v %v", getOK, "bar")
	// Test that bob can read.
	startClient(t, sh, "bob", "get", "foo").Expectf("%v %v", getOK, "bar")
	// Test that bob cannot write.
	startClient(t, sh, "bob", "set", "foo", "bar").Expectf("%v %v", setFailed, verror.ErrNoAccess.ID)

	// Stop the groups server and check that as a consequence "alice"
	// cannot read from the key value store server anymore.
	server.Shutdown(os.Interrupt)
	startClient(t, sh, "alice", "get", "foo").Expectf("%v %v", getFailed, verror.ErrNoAccess.ID)
}

func TestMain(m *testing.M) {
	os.Exit(v23test.Run(m.Run))
}
