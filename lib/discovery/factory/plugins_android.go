// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build android

// The paypal gatt library, which the default ble plugin depends on doesn't work on
// android.  Instead we use a jni version of the plugin.

package factory

import (
	"fmt"

	"v.io/x/ref/lib/discovery"
	"v.io/x/ref/lib/discovery/plugins/mdns"
)

var pluginFactories = map[string]func(host string) (discovery.Plugin, error){
	"mdns": mdns.New,
	"ble": func(string) (discovery.Plugin, error) {
		return nil, fmt.Errorf("ble factory not initalized")
	},
}

// SetBleFactory sets the plugin factory for ble.  This needs to be called before the first time
// the discovery api is used.
func SetBleFactory(creator func(string) (discovery.Plugin, error)) {
	pluginFactories["ble"] = creator
}