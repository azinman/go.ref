// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vclock

package vclock

import (
	"fmt"
	"time"
	"v.io/v23/vdl"
	vdltime "v.io/v23/vdlroot/time"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// VClockData is the persistent state of the Syncbase virtual clock.
// All times are UTC.
type VClockData struct {
	// System time at boot.
	SystemTimeAtBoot time.Time
	// Current estimate of NTP time minus system clock time.
	Skew time.Duration
	// Elapsed time since boot, as seen by VClockD. Used for detecting reboots.
	ElapsedTimeSinceBoot time.Duration
	// NTP server timestamp from the most recent NTP sync, or zero value if none.
	// Note, the NTP sync may have been performed by some peer device.
	LastNtpTs time.Time
	// Number of reboots since last NTP sync, accumulated across all hops of p2p
	// clock sync. E.g. if LastNtpTs came from some peer device, NumReboots will
	// equal that device's NumReboots at the time of sync plus the number of
	// reboots on this device since then.
	NumReboots uint16
	// Number of sync hops between this device and the source of LastNtpTs.
	NumHops uint16
}

func (VClockData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vclock.VClockData"`
}) {
}

func (m *VClockData) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var wireValue2 vdltime.Time
	if err := vdltime.TimeFromNative(&wireValue2, m.SystemTimeAtBoot); err != nil {
		return err
	}

	var5 := (wireValue2 == vdltime.Time{})
	if var5 {
		if err := fieldsTarget1.ZeroField("SystemTimeAtBoot"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("SystemTimeAtBoot")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue2.FillVDLTarget(fieldTarget4, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
				return err
			}
		}
	}
	var wireValue6 vdltime.Duration
	if err := vdltime.DurationFromNative(&wireValue6, m.Skew); err != nil {
		return err
	}

	var9 := (wireValue6 == vdltime.Duration{})
	if var9 {
		if err := fieldsTarget1.ZeroField("Skew"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget7, fieldTarget8, err := fieldsTarget1.StartField("Skew")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue6.FillVDLTarget(fieldTarget8, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget7, fieldTarget8); err != nil {
				return err
			}
		}
	}
	var wireValue10 vdltime.Duration
	if err := vdltime.DurationFromNative(&wireValue10, m.ElapsedTimeSinceBoot); err != nil {
		return err
	}

	var13 := (wireValue10 == vdltime.Duration{})
	if var13 {
		if err := fieldsTarget1.ZeroField("ElapsedTimeSinceBoot"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("ElapsedTimeSinceBoot")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue10.FillVDLTarget(fieldTarget12, tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
				return err
			}
		}
	}
	var wireValue14 vdltime.Time
	if err := vdltime.TimeFromNative(&wireValue14, m.LastNtpTs); err != nil {
		return err
	}

	var17 := (wireValue14 == vdltime.Time{})
	if var17 {
		if err := fieldsTarget1.ZeroField("LastNtpTs"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget15, fieldTarget16, err := fieldsTarget1.StartField("LastNtpTs")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue14.FillVDLTarget(fieldTarget16, tt.NonOptional().Field(3).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget15, fieldTarget16); err != nil {
				return err
			}
		}
	}
	var20 := (m.NumReboots == uint16(0))
	if var20 {
		if err := fieldsTarget1.ZeroField("NumReboots"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget18, fieldTarget19, err := fieldsTarget1.StartField("NumReboots")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget19.FromUint(uint64(m.NumReboots), tt.NonOptional().Field(4).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget18, fieldTarget19); err != nil {
				return err
			}
		}
	}
	var23 := (m.NumHops == uint16(0))
	if var23 {
		if err := fieldsTarget1.ZeroField("NumHops"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget21, fieldTarget22, err := fieldsTarget1.StartField("NumHops")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget22.FromUint(uint64(m.NumHops), tt.NonOptional().Field(5).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget21, fieldTarget22); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *VClockData) MakeVDLTarget() vdl.Target {
	return &VClockDataTarget{Value: m}
}

type VClockDataTarget struct {
	Value                      *VClockData
	systemTimeAtBootTarget     vdltime.TimeTarget
	skewTarget                 vdltime.DurationTarget
	elapsedTimeSinceBootTarget vdltime.DurationTarget
	lastNtpTsTarget            vdltime.TimeTarget
	numRebootsTarget           vdl.Uint16Target
	numHopsTarget              vdl.Uint16Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *VClockDataTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*VClockData)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *VClockDataTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "SystemTimeAtBoot":
		t.systemTimeAtBootTarget.Value = &t.Value.SystemTimeAtBoot
		target, err := &t.systemTimeAtBootTarget, error(nil)
		return nil, target, err
	case "Skew":
		t.skewTarget.Value = &t.Value.Skew
		target, err := &t.skewTarget, error(nil)
		return nil, target, err
	case "ElapsedTimeSinceBoot":
		t.elapsedTimeSinceBootTarget.Value = &t.Value.ElapsedTimeSinceBoot
		target, err := &t.elapsedTimeSinceBootTarget, error(nil)
		return nil, target, err
	case "LastNtpTs":
		t.lastNtpTsTarget.Value = &t.Value.LastNtpTs
		target, err := &t.lastNtpTsTarget, error(nil)
		return nil, target, err
	case "NumReboots":
		t.numRebootsTarget.Value = &t.Value.NumReboots
		target, err := &t.numRebootsTarget, error(nil)
		return nil, target, err
	case "NumHops":
		t.numHopsTarget.Value = &t.Value.NumHops
		target, err := &t.numHopsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vclock.VClockData", name)
	}
}
func (t *VClockDataTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *VClockDataTarget) ZeroField(name string) error {
	switch name {
	case "SystemTimeAtBoot":
		t.Value.SystemTimeAtBoot = time.Time{}
		return nil
	case "Skew":
		t.Value.Skew = time.Duration(0)
		return nil
	case "ElapsedTimeSinceBoot":
		t.Value.ElapsedTimeSinceBoot = time.Duration(0)
		return nil
	case "LastNtpTs":
		t.Value.LastNtpTs = time.Time{}
		return nil
	case "NumReboots":
		t.Value.NumReboots = uint16(0)
		return nil
	case "NumHops":
		t.Value.NumHops = uint16(0)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vclock.VClockData", name)
	}
}
func (t *VClockDataTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x VClockData) VDLIsZero() bool {
	if !x.SystemTimeAtBoot.IsZero() {
		return false
	}
	if x.Skew != time.Duration(0) {
		return false
	}
	if x.ElapsedTimeSinceBoot != time.Duration(0) {
		return false
	}
	if !x.LastNtpTs.IsZero() {
		return false
	}
	if x.NumReboots != 0 {
		return false
	}
	if x.NumHops != 0 {
		return false
	}
	return true
}

func (x VClockData) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*VClockData)(nil)).Elem()); err != nil {
		return err
	}
	if !x.SystemTimeAtBoot.IsZero() {
		if err := enc.NextField("SystemTimeAtBoot"); err != nil {
			return err
		}
		var wire vdltime.Time
		if err := vdltime.TimeFromNative(&wire, x.SystemTimeAtBoot); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Skew != time.Duration(0) {
		if err := enc.NextField("Skew"); err != nil {
			return err
		}
		var wire vdltime.Duration
		if err := vdltime.DurationFromNative(&wire, x.Skew); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.ElapsedTimeSinceBoot != time.Duration(0) {
		if err := enc.NextField("ElapsedTimeSinceBoot"); err != nil {
			return err
		}
		var wire vdltime.Duration
		if err := vdltime.DurationFromNative(&wire, x.ElapsedTimeSinceBoot); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.LastNtpTs.IsZero() {
		if err := enc.NextField("LastNtpTs"); err != nil {
			return err
		}
		var wire vdltime.Time
		if err := vdltime.TimeFromNative(&wire, x.LastNtpTs); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.NumReboots != 0 {
		if err := enc.NextField("NumReboots"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.Uint16Type); err != nil {
			return err
		}
		if err := enc.EncodeUint(uint64(x.NumReboots)); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.NumHops != 0 {
		if err := enc.NextField("NumHops"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.Uint16Type); err != nil {
			return err
		}
		if err := enc.EncodeUint(uint64(x.NumHops)); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *VClockData) VDLRead(dec vdl.Decoder) error {
	*x = VClockData{}
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "SystemTimeAtBoot":
			var wire vdltime.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.TimeToNative(wire, &x.SystemTimeAtBoot); err != nil {
				return err
			}
		case "Skew":
			var wire vdltime.Duration
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.DurationToNative(wire, &x.Skew); err != nil {
				return err
			}
		case "ElapsedTimeSinceBoot":
			var wire vdltime.Duration
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.DurationToNative(wire, &x.ElapsedTimeSinceBoot); err != nil {
				return err
			}
		case "LastNtpTs":
			var wire vdltime.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.TimeToNative(wire, &x.LastNtpTs); err != nil {
				return err
			}
		case "NumReboots":
			if err := dec.StartValue(); err != nil {
				return err
			}
			tmp, err := dec.DecodeUint(16)
			if err != nil {
				return err
			}
			x.NumReboots = uint16(tmp)
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "NumHops":
			if err := dec.StartValue(); err != nil {
				return err
			}
			tmp, err := dec.DecodeUint(16)
			if err != nil {
				return err
			}
			x.NumHops = uint16(tmp)
			if err := dec.FinishValue(); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*VClockData)(nil))

	return struct{}{}
}
