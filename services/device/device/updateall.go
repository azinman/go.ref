// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/services/device"
	"v.io/v23/verror"
	"v.io/x/lib/cmdline2"
	"v.io/x/ref/lib/v23cmd"
	deviceimpl "v.io/x/ref/services/device/internal/impl"
)

// TODO(caprita): Re-implement this with Glob, so that one can say instead,
// update <devicename>/apps/... or <devicename>/apps/appname/* etc.

// TODO(caprita): Add unit test.

var cmdUpdateAll = &cmdline2.Command{
	Runner:   v23cmd.RunnerFunc(runUpdateAll),
	Name:     "updateall",
	Short:    "Update all installations/instances of an application",
	Long:     "Given a name that can refer to an app instance or app installation or app or all apps on a device, updates all installations and instances under that name",
	ArgsName: "<object name>",
	ArgsLong: `
<object name> is the vanadium object name to update, as follows:

<devicename>/apps/apptitle/installationid/instanceid: updates the given instance, killing/restarting it if running

<devicename>/apps/apptitle/installationid: updates the given installation and then all its instances

<devicename>/apps/apptitle: updates all installations for the given app

<devicename>/apps: updates all apps on the device
`,
}

type updater func(ctx *context.T, env *cmdline2.Env, von string) error

func updateChildren(ctx *context.T, env *cmdline2.Env, von string, updateChild updater) error {
	ns := v23.GetNamespace(ctx)
	pattern := naming.Join(von, "*")
	c, err := ns.Glob(ctx, pattern)
	if err != nil {
		return fmt.Errorf("ns.Glob(%q) failed: %v", pattern, err)
	}
	var (
		pending     sync.WaitGroup
		numErrors   int
		numErrorsMu sync.Mutex
	)
	for res := range c {
		switch v := res.(type) {
		case *naming.GlobReplyEntry:
			pending.Add(1)
			go func() {
				if err := updateChild(ctx, env, v.Value.Name); err != nil {
					numErrorsMu.Lock()
					numErrors++
					numErrorsMu.Unlock()
				}
				pending.Done()
			}()
		case *naming.GlobReplyError:
			fmt.Fprintf(env.Stderr, "Glob error for %q: %v\n", v.Value.Name, v.Value.Error)
			numErrorsMu.Lock()
			numErrors++
			numErrorsMu.Unlock()
		}
	}
	pending.Wait()
	if numErrors > 0 {
		return fmt.Errorf("%d error(s) encountered while updating children", numErrors)
	}
	return nil
}

func instanceIsRunning(ctx *context.T, von string) (bool, error) {
	status, err := device.ApplicationClient(von).Status(ctx)
	if err != nil {
		return false, fmt.Errorf("Failed to get status for instance %q: %v", von, err)
	}
	s, ok := status.(device.StatusInstance)
	if !ok {
		return false, fmt.Errorf("Status for instance %q of wrong type (%T)", von, status)
	}
	return s.Value.State == device.InstanceStateRunning, nil
}

func updateInstance(ctx *context.T, env *cmdline2.Env, von string) (retErr error) {
	defer func() {
		if retErr == nil {
			fmt.Fprintf(env.Stdout, "Successfully updated instance %q.\n", von)
		} else {
			retErr = fmt.Errorf("failed to update instance %q: %v", von, retErr)
			fmt.Fprintf(env.Stderr, "ERROR: %v.\n", retErr)
		}
	}()
	running, err := instanceIsRunning(ctx, von)
	if err != nil {
		return err
	}
	if running {
		// Try killing the app.
		if err := device.ApplicationClient(von).Kill(ctx, killDeadline); err != nil {
			// Check the app's state again in case we killed it,
			// nevermind any errors.  The sleep is because Kill
			// currently (4/29/15) returns asynchronously with the
			// device manager shooting the app down.
			time.Sleep(time.Second)
			running, rerr := instanceIsRunning(ctx, von)
			if rerr != nil {
				return rerr
			}
			if running {
				return fmt.Errorf("failed to kill instance %q: %v", von, err)
			}
			fmt.Fprintf(env.Stderr, "Kill(%s) returned an error (%s) but app is now not running.\n", von, err)
		}
		// App was running, and we killed it.
		defer func() {
			// Re-start the instance.
			if err := device.ApplicationClient(von).Run(ctx); err != nil {
				err = fmt.Errorf("failed to run instance %q: %v", von, err)
				if retErr == nil {
					retErr = err
				} else {
					fmt.Fprintf(env.Stderr, "ERROR: %v.\n", err)
				}
			}
		}()
	}
	// Update the instance.
	switch err := device.ApplicationClient(von).Update(ctx); {
	case err == nil:
		return nil
	case verror.ErrorID(err) == deviceimpl.ErrUpdateNoOp.ID:
		// TODO(caprita): Ideally, we wouldn't even attempt a kill /
		// restart if there's no newer version of the application.
		fmt.Fprintf(env.Stdout, "Instance %q already up to date.\n", von)
		return nil
	default:
		return err
	}
}

func updateInstallation(ctx *context.T, env *cmdline2.Env, von string) (retErr error) {
	defer func() {
		if retErr == nil {
			fmt.Fprintf(env.Stdout, "Successfully updated installation %q.\n", von)
		} else {
			retErr = fmt.Errorf("failed to update installation %q: %v", von, retErr)
			fmt.Fprintf(env.Stderr, "ERROR: %v.\n", retErr)
		}
	}()

	// First, update the installation.
	switch err := device.ApplicationClient(von).Update(ctx); {
	case err == nil:
		fmt.Fprintf(env.Stdout, "Successfully updated version for installation %q.\n", von)
	case verror.ErrorID(err) == deviceimpl.ErrUpdateNoOp.ID:
		fmt.Fprintf(env.Stdout, "Installation %q already up to date.\n", von)
		// NOTE: we still proceed to update the instances in this case,
		// since it's possible that some instances are still running
		// from older versions.
	default:
		return err
	}
	// Then, update all the instances for the installation.
	return updateChildren(ctx, env, von, updateInstance)
}

func updateApp(ctx *context.T, env *cmdline2.Env, von string) error {
	if err := updateChildren(ctx, env, von, updateInstallation); err != nil {
		err = fmt.Errorf("failed to update app %q: %v", von, err)
		fmt.Fprintf(env.Stderr, "ERROR: %v.\n", err)
		return err
	}
	fmt.Fprintf(env.Stdout, "Successfully updated app %q.\n", von)
	return nil
}

func updateAllApps(ctx *context.T, env *cmdline2.Env, von string) error {
	if err := updateChildren(ctx, env, von, updateApp); err != nil {
		err = fmt.Errorf("failed to update all apps %q: %v", von, err)
		fmt.Fprintf(env.Stderr, "ERROR: %v.\n", err)
		return err
	}
	fmt.Fprintf(env.Stdout, "Successfully updated all apps %q.\n", von)
	return nil
}

func runUpdateAll(ctx *context.T, env *cmdline2.Env, args []string) error {
	if expected, got := 1, len(args); expected != got {
		return env.UsageErrorf("updateall: incorrect number of arguments, expected %d, got %d", expected, got)
	}
	appVON := args[0]
	components := strings.Split(appVON, "/")
	var prefix string
	// TODO(caprita): Trying to figure out what the app suffix is by looking
	// for "/apps/" in the name is hacky and error-prone (e.g., what if an
	// app has the title "apps").  Instead, we should either query the
	// server or use resolution to split up the name into address and
	// suffix.
	for i := len(components) - 1; i >= 0; i-- {
		if components[i] == "apps" {
			prefix = naming.Join(components[:i+1]...)
			components = components[i+1:]
			break
		}
	}
	if prefix == "" {
		return fmt.Errorf("couldn't recognize app name: %q", appVON)
	}
	fmt.Printf("prefix: %q, components: %q\n", prefix, components)
	switch len(components) {
	case 0:
		return updateAllApps(ctx, env, appVON)
	case 1:
		return updateApp(ctx, env, appVON)
	case 2:
		return updateInstallation(ctx, env, appVON)
	case 3:
		return updateInstance(ctx, env, appVON)
	}
	return env.UsageErrorf("updateall: name %q does not refer to a supported app hierarchy object", appVON)
}
