// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: watchable

package watchable

import (
	"fmt"
	"v.io/v23/vdl"
	"v.io/x/ref/services/syncbase/server/interfaces"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// SyncgroupOp represents a change in the set of prefixes that should be tracked
// by sync, i.e. the union of prefix sets across all syncgroups. Note that an
// individual syncgroup's prefixes cannot be changed; this record type is used
// to track changes due to syncgroup create/join/leave/destroy.
type SyncgroupOp struct {
	SgId     interfaces.GroupId
	Prefixes []string
	Remove   bool
}

func (SyncgroupOp) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/watchable.SyncgroupOp"`
}) {
}

func (m *SyncgroupOp) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.SgId == interfaces.GroupId(""))
	if var4 {
		if err := fieldsTarget1.ZeroField("SgId"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("SgId")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.SgId.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Prefixes) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Prefixes"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Prefixes")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			listTarget8, err := fieldTarget6.StartList(tt.NonOptional().Field(1).Type, len(m.Prefixes))
			if err != nil {
				return err
			}
			for i, elem10 := range m.Prefixes {
				elemTarget9, err := listTarget8.StartElem(i)
				if err != nil {
					return err
				}
				if err := elemTarget9.FromString(string(elem10), tt.NonOptional().Field(1).Type.Elem()); err != nil {
					return err
				}
				if err := listTarget8.FinishElem(elemTarget9); err != nil {
					return err
				}
			}
			if err := fieldTarget6.FinishList(listTarget8); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var13 := (m.Remove == false)
	if var13 {
		if err := fieldsTarget1.ZeroField("Remove"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("Remove")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget12.FromBool(bool(m.Remove), tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *SyncgroupOp) MakeVDLTarget() vdl.Target {
	return &SyncgroupOpTarget{Value: m}
}

type SyncgroupOpTarget struct {
	Value          *SyncgroupOp
	sgIdTarget     interfaces.GroupIdTarget
	prefixesTarget vdl.StringSliceTarget
	removeTarget   vdl.BoolTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *SyncgroupOpTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*SyncgroupOp)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *SyncgroupOpTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "SgId":
		t.sgIdTarget.Value = &t.Value.SgId
		target, err := &t.sgIdTarget, error(nil)
		return nil, target, err
	case "Prefixes":
		t.prefixesTarget.Value = &t.Value.Prefixes
		target, err := &t.prefixesTarget, error(nil)
		return nil, target, err
	case "Remove":
		t.removeTarget.Value = &t.Value.Remove
		target, err := &t.removeTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/watchable.SyncgroupOp", name)
	}
}
func (t *SyncgroupOpTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *SyncgroupOpTarget) ZeroField(name string) error {
	switch name {
	case "SgId":
		t.Value.SgId = interfaces.GroupId("")
		return nil
	case "Prefixes":
		t.Value.Prefixes = []string(nil)
		return nil
	case "Remove":
		t.Value.Remove = false
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/watchable.SyncgroupOp", name)
	}
}
func (t *SyncgroupOpTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x SyncgroupOp) VDLIsZero() bool {
	if x.SgId != "" {
		return false
	}
	if len(x.Prefixes) != 0 {
		return false
	}
	if x.Remove {
		return false
	}
	return true
}

func (x SyncgroupOp) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*SyncgroupOp)(nil)).Elem()); err != nil {
		return err
	}
	if x.SgId != "" {
		if err := enc.NextField("SgId"); err != nil {
			return err
		}
		if err := x.SgId.VDLWrite(enc); err != nil {
			return err
		}
	}
	if len(x.Prefixes) != 0 {
		if err := enc.NextField("Prefixes"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.Prefixes); err != nil {
			return err
		}
	}
	if x.Remove {
		if err := enc.NextField("Remove"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.BoolType); err != nil {
			return err
		}
		if err := enc.EncodeBool(x.Remove); err != nil {
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

func __VDLWriteAnon_list_1(enc vdl.Encoder, x []string) error {
	if err := enc.StartValue(vdl.TypeOf((*[]string)(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for i := 0; i < len(x); i++ {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x[i]); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *SyncgroupOp) VDLRead(dec vdl.Decoder) error {
	*x = SyncgroupOp{}
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
		case "SgId":
			if err := x.SgId.VDLRead(dec); err != nil {
				return err
			}
		case "Prefixes":
			if err := __VDLReadAnon_list_1(dec, &x.Prefixes); err != nil {
				return err
			}
		case "Remove":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Remove, err = dec.DecodeBool(); err != nil {
				return err
			}
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

func __VDLReadAnon_list_1(dec vdl.Decoder, x *[]string) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible list %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len > 0:
		*x = make([]string, 0, len)
	default:
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var elem string
		if err := dec.StartValue(); err != nil {
			return err
		}
		var err error
		if elem, err = dec.DecodeString(); err != nil {
			return err
		}
		if err := dec.FinishValue(); err != nil {
			return err
		}
		*x = append(*x, elem)
	}
}

// SyncSnapshotOp represents a snapshot operation when creating and joining a
// syncgroup. The sync watcher needs to get a snapshot of the Database at the
// point of creating/joining a syncgroup. A SyncSnapshotOp entry is written to
// the log for each Database key that falls within the syncgroup prefixes. This
// allows sync to initialize its metadata at the correct versions of the objects
// when they become syncable. These log entries should be filtered by the
// client-facing Watch interface because the user data did not actually change.
type SyncSnapshotOp struct {
	Key     []byte
	Version []byte
}

func (SyncSnapshotOp) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/watchable.SyncSnapshotOp"`
}) {
}

func (m *SyncSnapshotOp) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var var4 bool
	if len(m.Key) == 0 {
		var4 = true
	}
	if var4 {
		if err := fieldsTarget1.ZeroField("Key"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Key")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := fieldTarget3.FromBytes([]byte(m.Key), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Version) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Version"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Version")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := fieldTarget6.FromBytes([]byte(m.Version), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *SyncSnapshotOp) MakeVDLTarget() vdl.Target {
	return &SyncSnapshotOpTarget{Value: m}
}

type SyncSnapshotOpTarget struct {
	Value         *SyncSnapshotOp
	keyTarget     vdl.BytesTarget
	versionTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *SyncSnapshotOpTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*SyncSnapshotOp)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *SyncSnapshotOpTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Key":
		t.keyTarget.Value = &t.Value.Key
		target, err := &t.keyTarget, error(nil)
		return nil, target, err
	case "Version":
		t.versionTarget.Value = &t.Value.Version
		target, err := &t.versionTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/watchable.SyncSnapshotOp", name)
	}
}
func (t *SyncSnapshotOpTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *SyncSnapshotOpTarget) ZeroField(name string) error {
	switch name {
	case "Key":
		t.Value.Key = []byte(nil)
		return nil
	case "Version":
		t.Value.Version = []byte(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/watchable.SyncSnapshotOp", name)
	}
}
func (t *SyncSnapshotOpTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x SyncSnapshotOp) VDLIsZero() bool {
	if len(x.Key) != 0 {
		return false
	}
	if len(x.Version) != 0 {
		return false
	}
	return true
}

func (x SyncSnapshotOp) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*SyncSnapshotOp)(nil)).Elem()); err != nil {
		return err
	}
	if len(x.Key) != 0 {
		if err := enc.NextField("Key"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*[]byte)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBytes(x.Key); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if len(x.Version) != 0 {
		if err := enc.NextField("Version"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*[]byte)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBytes(x.Version); err != nil {
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

func (x *SyncSnapshotOp) VDLRead(dec vdl.Decoder) error {
	*x = SyncSnapshotOp{}
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
		case "Key":
			if err := dec.StartValue(); err != nil {
				return err
			}
			if err := dec.DecodeBytes(-1, &x.Key); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Version":
			if err := dec.StartValue(); err != nil {
				return err
			}
			if err := dec.DecodeBytes(-1, &x.Version); err != nil {
				return err
			}
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

// StateChange represents the set of types of state change requests possible.
type StateChange int

const (
	StateChangePauseSync StateChange = iota
	StateChangeResumeSync
)

// StateChangeAll holds all labels for StateChange.
var StateChangeAll = [...]StateChange{StateChangePauseSync, StateChangeResumeSync}

// StateChangeFromString creates a StateChange from a string label.
func StateChangeFromString(label string) (x StateChange, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *StateChange) Set(label string) error {
	switch label {
	case "PauseSync", "pausesync":
		*x = StateChangePauseSync
		return nil
	case "ResumeSync", "resumesync":
		*x = StateChangeResumeSync
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in watchable.StateChange", label)
}

// String returns the string label of x.
func (x StateChange) String() string {
	switch x {
	case StateChangePauseSync:
		return "PauseSync"
	case StateChangeResumeSync:
		return "ResumeSync"
	}
	return ""
}

func (StateChange) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/watchable.StateChange"`
	Enum struct{ PauseSync, ResumeSync string }
}) {
}

func (m *StateChange) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromEnumLabel((*m).String(), tt); err != nil {
		return err
	}
	return nil
}

func (m *StateChange) MakeVDLTarget() vdl.Target {
	return &StateChangeTarget{Value: m}
}

type StateChangeTarget struct {
	Value *StateChange
	vdl.TargetBase
}

func (t *StateChangeTarget) FromEnumLabel(src string, tt *vdl.Type) error {

	if ttWant := vdl.TypeOf((*StateChange)(nil)); !vdl.Compatible(tt, ttWant) {
		return fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	switch src {
	case "PauseSync":
		*t.Value = 0
	case "ResumeSync":
		*t.Value = 1
	default:
		return fmt.Errorf("label %s not in enum StateChange", src)
	}

	return nil
}

func (x StateChange) VDLIsZero() bool {
	return x == StateChangePauseSync
}

func (x StateChange) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*StateChange)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeString(x.String()); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *StateChange) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	enum, err := dec.DecodeString()
	if err != nil {
		return err
	}
	if err := x.Set(enum); err != nil {
		return err
	}
	return dec.FinishValue()
}

// DbStateChangeRequestOp represents a database state change request.
// Specifically there are two events that create this op:
// PauseSync, indicating a client request to pause sync on this db.
// ResumeSync, indicating a client request to resume sync on this db.
// Client watcher will ignore this op.
type DbStateChangeRequestOp struct {
	RequestType StateChange
}

func (DbStateChangeRequestOp) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/watchable.DbStateChangeRequestOp"`
}) {
}

func (m *DbStateChangeRequestOp) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.RequestType == StateChangePauseSync)
	if var4 {
		if err := fieldsTarget1.ZeroField("RequestType"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("RequestType")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.RequestType.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *DbStateChangeRequestOp) MakeVDLTarget() vdl.Target {
	return &DbStateChangeRequestOpTarget{Value: m}
}

type DbStateChangeRequestOpTarget struct {
	Value             *DbStateChangeRequestOp
	requestTypeTarget StateChangeTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *DbStateChangeRequestOpTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*DbStateChangeRequestOp)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *DbStateChangeRequestOpTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "RequestType":
		t.requestTypeTarget.Value = &t.Value.RequestType
		target, err := &t.requestTypeTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/watchable.DbStateChangeRequestOp", name)
	}
}
func (t *DbStateChangeRequestOpTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *DbStateChangeRequestOpTarget) ZeroField(name string) error {
	switch name {
	case "RequestType":
		t.Value.RequestType = StateChangePauseSync
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/watchable.DbStateChangeRequestOp", name)
	}
}
func (t *DbStateChangeRequestOpTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x DbStateChangeRequestOp) VDLIsZero() bool {
	return x == DbStateChangeRequestOp{}
}

func (x DbStateChangeRequestOp) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*DbStateChangeRequestOp)(nil)).Elem()); err != nil {
		return err
	}
	if x.RequestType != StateChangePauseSync {
		if err := enc.NextField("RequestType"); err != nil {
			return err
		}
		if err := x.RequestType.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *DbStateChangeRequestOp) VDLRead(dec vdl.Decoder) error {
	*x = DbStateChangeRequestOp{}
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
		case "RequestType":
			if err := x.RequestType.VDLRead(dec); err != nil {
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
	vdl.Register((*SyncgroupOp)(nil))
	vdl.Register((*SyncSnapshotOp)(nil))
	vdl.Register((*StateChange)(nil))
	vdl.Register((*DbStateChangeRequestOp)(nil))

	return struct{}{}
}
