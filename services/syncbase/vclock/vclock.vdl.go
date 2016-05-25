// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vclock

package vclock

import (
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

func (x VClockData) VDLIsZero() bool {
	if !x.SystemTimeAtBoot.IsZero() {
		return false
	}
	if x.Skew != 0 {
		return false
	}
	if x.ElapsedTimeSinceBoot != 0 {
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
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
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
	if x.Skew != 0 {
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
	if x.ElapsedTimeSinceBoot != 0 {
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
		if err := enc.NextFieldValueUint("NumReboots", vdl.Uint16Type, uint64(x.NumReboots)); err != nil {
			return err
		}
	}
	if x.NumHops != 0 {
		if err := enc.NextFieldValueUint("NumHops", vdl.Uint16Type, uint64(x.NumHops)); err != nil {
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
	if err := dec.StartValue(__VDLType_struct_1); err != nil {
		return err
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
			switch value, err := dec.ReadValueUint(16); {
			case err != nil:
				return err
			default:
				x.NumReboots = uint16(value)
			}
		case "NumHops":
			switch value, err := dec.ReadValueUint(16); {
			case err != nil:
				return err
			default:
				x.NumHops = uint16(value)
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1 *vdl.Type
	__VDLType_struct_2 *vdl.Type
	__VDLType_struct_3 *vdl.Type
)

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

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*VClockData)(nil)).Elem()
	__VDLType_struct_2 = vdl.TypeOf((*vdltime.Time)(nil)).Elem()
	__VDLType_struct_3 = vdl.TypeOf((*vdltime.Duration)(nil)).Elem()

	return struct{}{}
}
