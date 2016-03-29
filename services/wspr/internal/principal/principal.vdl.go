// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: principal

package principal

import (
	"fmt"
	"v.io/v23/security"
	"v.io/v23/vdl"
	"v.io/v23/vdl/vdlconv"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Identifier of a blessings cache entry.
type BlessingsId uint32

func (BlessingsId) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsId"`
}) {
}

func (m *BlessingsId) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromUint(uint64((*m)), tt); err != nil {
		return err
	}
	return nil
}

func (m *BlessingsId) MakeVDLTarget() vdl.Target {
	return &BlessingsIdTarget{Value: m}
}

type BlessingsIdTarget struct {
	Value *BlessingsId
	vdl.TargetBase
}

func (t *BlessingsIdTarget) FromUint(src uint64, tt *vdl.Type) error {

	val, err := vdlconv.Uint64ToUint32(src)
	if err != nil {
		return err
	}
	*t.Value = BlessingsId(val)

	return nil
}
func (t *BlessingsIdTarget) FromInt(src int64, tt *vdl.Type) error {

	val, err := vdlconv.Int64ToUint32(src)
	if err != nil {
		return err
	}
	*t.Value = BlessingsId(val)

	return nil
}
func (t *BlessingsIdTarget) FromFloat(src float64, tt *vdl.Type) error {

	val, err := vdlconv.Float64ToUint32(src)
	if err != nil {
		return err
	}
	*t.Value = BlessingsId(val)

	return nil
}

type BlessingsCacheAddMessage struct {
	CacheId   BlessingsId
	Blessings security.Blessings
}

func (BlessingsCacheAddMessage) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsCacheAddMessage"`
}) {
}

func (m *BlessingsCacheAddMessage) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("CacheId")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.CacheId.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	var wireValue4 security.WireBlessings
	if err := security.WireBlessingsFromNative(&wireValue4, m.Blessings); err != nil {
		return err
	}

	keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Blessings")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue4.FillVDLTarget(fieldTarget6, tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlessingsCacheAddMessage) MakeVDLTarget() vdl.Target {
	return &BlessingsCacheAddMessageTarget{Value: m}
}

type BlessingsCacheAddMessageTarget struct {
	Value           *BlessingsCacheAddMessage
	cacheIdTarget   BlessingsIdTarget
	blessingsTarget security.WireBlessingsTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlessingsCacheAddMessageTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlessingsCacheAddMessage)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlessingsCacheAddMessageTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "CacheId":
		t.cacheIdTarget.Value = &t.Value.CacheId
		target, err := &t.cacheIdTarget, error(nil)
		return nil, target, err
	case "Blessings":
		t.blessingsTarget.Value = &t.Value.Blessings
		target, err := &t.blessingsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/wspr/internal/principal.BlessingsCacheAddMessage", name)
	}
}
func (t *BlessingsCacheAddMessageTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlessingsCacheAddMessageTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// Message from Blessings Cache GC to delete a cache entry in Javascript.
type BlessingsCacheDeleteMessage struct {
	CacheId BlessingsId
	// Number of references expected. Javascript should wait until this number
	// has been received before deleting the entry because up until that point
	// messages with further references are expected.
	DeleteAfter uint32
}

func (BlessingsCacheDeleteMessage) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsCacheDeleteMessage"`
}) {
}

func (m *BlessingsCacheDeleteMessage) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("CacheId")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.CacheId.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("DeleteAfter")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromUint(uint64(m.DeleteAfter), tt.NonOptional().Field(1).Type); err != nil {
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

func (m *BlessingsCacheDeleteMessage) MakeVDLTarget() vdl.Target {
	return &BlessingsCacheDeleteMessageTarget{Value: m}
}

type BlessingsCacheDeleteMessageTarget struct {
	Value             *BlessingsCacheDeleteMessage
	cacheIdTarget     BlessingsIdTarget
	deleteAfterTarget vdl.Uint32Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlessingsCacheDeleteMessageTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlessingsCacheDeleteMessage)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlessingsCacheDeleteMessageTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "CacheId":
		t.cacheIdTarget.Value = &t.Value.CacheId
		target, err := &t.cacheIdTarget, error(nil)
		return nil, target, err
	case "DeleteAfter":
		t.deleteAfterTarget.Value = &t.Value.DeleteAfter
		target, err := &t.deleteAfterTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/wspr/internal/principal.BlessingsCacheDeleteMessage", name)
	}
}
func (t *BlessingsCacheDeleteMessageTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlessingsCacheDeleteMessageTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

type (
	// BlessingsCacheMessage represents any single field of the BlessingsCacheMessage union type.
	BlessingsCacheMessage interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the BlessingsCacheMessage union type.
		__VDLReflect(__BlessingsCacheMessageReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
	}
	// BlessingsCacheMessageAdd represents field Add of the BlessingsCacheMessage union type.
	BlessingsCacheMessageAdd struct{ Value BlessingsCacheAddMessage }
	// BlessingsCacheMessageDelete represents field Delete of the BlessingsCacheMessage union type.
	BlessingsCacheMessageDelete struct{ Value BlessingsCacheDeleteMessage }
	// __BlessingsCacheMessageReflect describes the BlessingsCacheMessage union type.
	__BlessingsCacheMessageReflect struct {
		Name               string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsCacheMessage"`
		Type               BlessingsCacheMessage
		UnionTargetFactory blessingsCacheMessageTargetFactory
		Union              struct {
			Add    BlessingsCacheMessageAdd
			Delete BlessingsCacheMessageDelete
		}
	}
)

func (x BlessingsCacheMessageAdd) Index() int                                  { return 0 }
func (x BlessingsCacheMessageAdd) Interface() interface{}                      { return x.Value }
func (x BlessingsCacheMessageAdd) Name() string                                { return "Add" }
func (x BlessingsCacheMessageAdd) __VDLReflect(__BlessingsCacheMessageReflect) {}

func (m BlessingsCacheMessageAdd) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Add")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlessingsCacheMessageAdd) MakeVDLTarget() vdl.Target {
	return nil
}

func (x BlessingsCacheMessageDelete) Index() int                                  { return 1 }
func (x BlessingsCacheMessageDelete) Interface() interface{}                      { return x.Value }
func (x BlessingsCacheMessageDelete) Name() string                                { return "Delete" }
func (x BlessingsCacheMessageDelete) __VDLReflect(__BlessingsCacheMessageReflect) {}

func (m BlessingsCacheMessageDelete) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Delete")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(1).Type); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlessingsCacheMessageDelete) MakeVDLTarget() vdl.Target {
	return nil
}

type BlessingsCacheMessageTarget struct {
	Value     *BlessingsCacheMessage
	fieldName string

	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlessingsCacheMessageTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if ttWant := vdl.TypeOf((*BlessingsCacheMessage)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}

	return t, nil
}
func (t *BlessingsCacheMessageTarget) StartField(name string) (key, field vdl.Target, _ error) {
	t.fieldName = name
	switch name {
	case "Add":
		val := BlessingsCacheAddMessage{}
		return nil, &BlessingsCacheAddMessageTarget{Value: &val}, nil
	case "Delete":
		val := BlessingsCacheDeleteMessage{}
		return nil, &BlessingsCacheDeleteMessageTarget{Value: &val}, nil
	default:
		return nil, nil, fmt.Errorf("field %s not in union v.io/x/ref/services/wspr/internal/principal.BlessingsCacheMessage", name)
	}
}
func (t *BlessingsCacheMessageTarget) FinishField(_, fieldTarget vdl.Target) error {
	switch t.fieldName {
	case "Add":
		*t.Value = BlessingsCacheMessageAdd{*(fieldTarget.(*BlessingsCacheAddMessageTarget)).Value}
	case "Delete":
		*t.Value = BlessingsCacheMessageDelete{*(fieldTarget.(*BlessingsCacheDeleteMessageTarget)).Value}
	}
	return nil
}
func (t *BlessingsCacheMessageTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

type blessingsCacheMessageTargetFactory struct{}

func (t blessingsCacheMessageTargetFactory) VDLMakeUnionTarget(union interface{}) (vdl.Target, error) {
	if typedUnion, ok := union.(*BlessingsCacheMessage); ok {
		return &BlessingsCacheMessageTarget{Value: typedUnion}, nil
	}
	return nil, fmt.Errorf("got %T, want *BlessingsCacheMessage", union)
}

// Create zero values for each type.
var (
	__VDLZeroBlessingsId                 = BlessingsId(0)
	__VDLZeroBlessingsCacheAddMessage    = BlessingsCacheAddMessage{}
	__VDLZeroBlessingsCacheDeleteMessage = BlessingsCacheDeleteMessage{}
	__VDLZeroBlessingsCacheMessage       = BlessingsCacheMessage(BlessingsCacheMessageAdd{})
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

	// Register types.
	vdl.Register((*BlessingsId)(nil))
	vdl.Register((*BlessingsCacheAddMessage)(nil))
	vdl.Register((*BlessingsCacheDeleteMessage)(nil))
	vdl.Register((*BlessingsCacheMessage)(nil))

	return struct{}{}
}
