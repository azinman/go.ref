// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: uuid.vdl

package testdata

import (
	// VDL system imports
	"v.io/v23/vdl"
)

// UuidTestData represents the inputs and outputs for a uuid test.
type UuidTestData struct {
	// In is the input string.
	In string
	// Want is the expected uuid's human-readable string form.
	Want string
}

func (UuidTestData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/discovery/testdata.UuidTestData"`
}) {
}

func (m *UuidTestData) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_uuid_v_io_x_ref_lib_discovery_testdata_UuidTestData == nil || __VDLTypeuuid0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("In")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.In), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Want")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromString(string(m.Want), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *UuidTestData) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*UuidTestData)(nil))
}

var __VDLTypeuuid0 *vdl.Type = vdl.TypeOf((*UuidTestData)(nil))
var __VDLType_uuid_v_io_x_ref_lib_discovery_testdata_UuidTestData *vdl.Type = vdl.TypeOf(UuidTestData{})

func __VDLEnsureNativeBuilt_uuid() {
}

var ServiceUuidTest = []UuidTestData{
	{
		In:   "v.io",
		Want: "2101363c-688d-548a-a600-34d506e1aad0",
	},
	{
		In:   "v.io/v23/abc",
		Want: "6726c4e5-b6eb-5547-9228-b2913f4fad52",
	},
	{
		In:   "v.io/v23/abc/xyz",
		Want: "be8a57d7-931d-5ee4-9243-0bebde0029a5",
	},
}

var AttributeUuidTest = []UuidTestData{
	{
		In:   "name",
		Want: "217a496d-3aae-5748-baf0-a77555f8f4f4",
	},
	{
		In:   "_attr",
		Want: "6c020e4b-9a59-5c7f-92e7-45954a16a402",
	},
	{
		In:   "xyz",
		Want: "c10b25a2-2d4d-5a19-bb7c-1ee1c4972b4c",
	},
}
