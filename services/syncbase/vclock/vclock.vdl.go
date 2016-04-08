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
	time_2 "v.io/v23/vdlroot/time"
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
	var wireValue2 time_2.Time
	if err := time_2.TimeFromNative(&wireValue2, m.SystemTimeAtBoot); err != nil {
		return err
	}

	var5 := (wireValue2 == time_2.Time{})
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
	var wireValue6 time_2.Duration
	if err := time_2.DurationFromNative(&wireValue6, m.Skew); err != nil {
		return err
	}

	var9 := (wireValue6 == time_2.Duration{})
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
	var wireValue10 time_2.Duration
	if err := time_2.DurationFromNative(&wireValue10, m.ElapsedTimeSinceBoot); err != nil {
		return err
	}

	var13 := (wireValue10 == time_2.Duration{})
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
	var wireValue14 time_2.Time
	if err := time_2.TimeFromNative(&wireValue14, m.LastNtpTs); err != nil {
		return err
	}

	var17 := (wireValue14 == time_2.Time{})
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
	systemTimeAtBootTarget     time_2.TimeTarget
	skewTarget                 time_2.DurationTarget
	elapsedTimeSinceBootTarget time_2.DurationTarget
	lastNtpTsTarget            time_2.TimeTarget
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
		t.Value.SystemTimeAtBoot = func() time.Time {
			var native time.Time
			if err := vdl.Convert(&native, time_2.Time{}); err != nil {
				panic(err)
			}
			return native
		}()
		return nil
	case "Skew":
		t.Value.Skew = func() time.Duration {
			var native time.Duration
			if err := vdl.Convert(&native, time_2.Duration{}); err != nil {
				panic(err)
			}
			return native
		}()
		return nil
	case "ElapsedTimeSinceBoot":
		t.Value.ElapsedTimeSinceBoot = func() time.Duration {
			var native time.Duration
			if err := vdl.Convert(&native, time_2.Duration{}); err != nil {
				panic(err)
			}
			return native
		}()
		return nil
	case "LastNtpTs":
		t.Value.LastNtpTs = func() time.Time {
			var native time.Time
			if err := vdl.Convert(&native, time_2.Time{}); err != nil {
				panic(err)
			}
			return native
		}()
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

func (x *VClockData) VDLRead(dec vdl.Decoder) error {
	*x = VClockData{}
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	match := 0
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			if match == 0 && dec.Type().NumField() > 0 {
				return fmt.Errorf("no matching fields in struct %T, from %v", *x, dec.Type())
			}
			return dec.FinishValue()
		case "SystemTimeAtBoot":
			match++
			var wire time_2.Time
			if err = wire.VDLRead(dec); err != nil {
				return err
			}
			if err = time_2.TimeToNative(wire, &x.SystemTimeAtBoot); err != nil {
				return err
			}
		case "Skew":
			match++
			var wire time_2.Duration
			if err = wire.VDLRead(dec); err != nil {
				return err
			}
			if err = time_2.DurationToNative(wire, &x.Skew); err != nil {
				return err
			}
		case "ElapsedTimeSinceBoot":
			match++
			var wire time_2.Duration
			if err = wire.VDLRead(dec); err != nil {
				return err
			}
			if err = time_2.DurationToNative(wire, &x.ElapsedTimeSinceBoot); err != nil {
				return err
			}
		case "LastNtpTs":
			match++
			var wire time_2.Time
			if err = wire.VDLRead(dec); err != nil {
				return err
			}
			if err = time_2.TimeToNative(wire, &x.LastNtpTs); err != nil {
				return err
			}
		case "NumReboots":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			tmp, err := dec.DecodeUint(16)
			if err != nil {
				return err
			}
			x.NumReboots = uint16(tmp)
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "NumHops":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			tmp, err := dec.DecodeUint(16)
			if err != nil {
				return err
			}
			x.NumHops = uint16(tmp)
			if err = dec.FinishValue(); err != nil {
				return err
			}
		default:
			if err = dec.SkipValue(); err != nil {
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
