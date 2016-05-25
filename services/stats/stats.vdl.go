// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: stats

// Packages stats defines the non-native types exported by the stats service.
package stats

import (
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// HistogramBucket is one histogram bucket.
type HistogramBucket struct {
	// LowBound is the lower bound of the bucket.
	LowBound int64
	// Count is the number of values in the bucket.
	Count int64
}

func (HistogramBucket) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/stats.HistogramBucket"`
}) {
}

func (x HistogramBucket) VDLIsZero() bool {
	return x == HistogramBucket{}
}

func (x HistogramBucket) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.LowBound != 0 {
		if err := enc.NextFieldValueInt("LowBound", vdl.Int64Type, x.LowBound); err != nil {
			return err
		}
	}
	if x.Count != 0 {
		if err := enc.NextFieldValueInt("Count", vdl.Int64Type, x.Count); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *HistogramBucket) VDLRead(dec vdl.Decoder) error {
	*x = HistogramBucket{}
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
		case "LowBound":
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.LowBound = value
			}
		case "Count":
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.Count = value
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

// HistogramValue is the value of Histogram objects.
type HistogramValue struct {
	// Count is the total number of values added to the histogram.
	Count int64
	// Sum is the sum of all the values added to the histogram.
	Sum int64
	// Min is the minimum of all the values added to the histogram.
	Min int64
	// Max is the maximum of all the values added to the histogram.
	Max int64
	// Buckets contains all the buckets of the histogram.
	Buckets []HistogramBucket
}

func (HistogramValue) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/stats.HistogramValue"`
}) {
}

func (x HistogramValue) VDLIsZero() bool {
	if x.Count != 0 {
		return false
	}
	if x.Sum != 0 {
		return false
	}
	if x.Min != 0 {
		return false
	}
	if x.Max != 0 {
		return false
	}
	if len(x.Buckets) != 0 {
		return false
	}
	return true
}

func (x HistogramValue) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_2); err != nil {
		return err
	}
	if x.Count != 0 {
		if err := enc.NextFieldValueInt("Count", vdl.Int64Type, x.Count); err != nil {
			return err
		}
	}
	if x.Sum != 0 {
		if err := enc.NextFieldValueInt("Sum", vdl.Int64Type, x.Sum); err != nil {
			return err
		}
	}
	if x.Min != 0 {
		if err := enc.NextFieldValueInt("Min", vdl.Int64Type, x.Min); err != nil {
			return err
		}
	}
	if x.Max != 0 {
		if err := enc.NextFieldValueInt("Max", vdl.Int64Type, x.Max); err != nil {
			return err
		}
	}
	if len(x.Buckets) != 0 {
		if err := enc.NextField("Buckets"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.Buckets); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_1(enc vdl.Encoder, x []HistogramBucket) error {
	if err := enc.StartValue(__VDLType_list_3); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *HistogramValue) VDLRead(dec vdl.Decoder) error {
	*x = HistogramValue{}
	if err := dec.StartValue(__VDLType_struct_2); err != nil {
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
		case "Count":
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.Count = value
			}
		case "Sum":
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.Sum = value
			}
		case "Min":
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.Min = value
			}
		case "Max":
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.Max = value
			}
		case "Buckets":
			if err := __VDLReadAnon_list_1(dec, &x.Buckets); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_1(dec vdl.Decoder, x *[]HistogramBucket) error {
	if err := dec.StartValue(__VDLType_list_3); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]HistogramBucket, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem HistogramBucket
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1 *vdl.Type
	__VDLType_struct_2 *vdl.Type
	__VDLType_list_3   *vdl.Type
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
	vdl.Register((*HistogramBucket)(nil))
	vdl.Register((*HistogramValue)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*HistogramBucket)(nil)).Elem()
	__VDLType_struct_2 = vdl.TypeOf((*HistogramValue)(nil)).Elem()
	__VDLType_list_3 = vdl.TypeOf((*[]HistogramBucket)(nil))

	return struct{}{}
}
