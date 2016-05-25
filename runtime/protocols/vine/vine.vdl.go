// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vine

package vine

import (
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/rpc"
	"v.io/v23/vdl"
	"v.io/v23/verror"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// PeerKey is a key that represents a connection from a Dialer tag to an Acceptor tag.
type PeerKey struct {
	Dialer   string
	Acceptor string
}

func (PeerKey) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/runtime/protocols/vine.PeerKey"`
}) {
}

func (x PeerKey) VDLIsZero() bool {
	return x == PeerKey{}
}

func (x PeerKey) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.Dialer != "" {
		if err := enc.NextFieldValueString("Dialer", vdl.StringType, x.Dialer); err != nil {
			return err
		}
	}
	if x.Acceptor != "" {
		if err := enc.NextFieldValueString("Acceptor", vdl.StringType, x.Acceptor); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *PeerKey) VDLRead(dec vdl.Decoder) error {
	*x = PeerKey{}
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
		case "Dialer":
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Dialer = value
			}
		case "Acceptor":
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Acceptor = value
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

// PeerBehavior specifies characteristics of a connection.
type PeerBehavior struct {
	// Reachable specifies whether the outgoing or incoming connection can be
	// dialed or accepted.
	// TODO(suharshs): Make this a user defined error which vine will return instead of a bool.
	Reachable bool
	// Discoverable specifies whether the Dialer can advertise a discovery packet
	// to the Acceptor. This is useful for emulating neighborhoods.
	// TODO(suharshs): Discoverable should always be bidirectional. It is unrealistic for
	// A to discover B, but not vice versa.
	Discoverable bool
}

func (PeerBehavior) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/runtime/protocols/vine.PeerBehavior"`
}) {
}

func (x PeerBehavior) VDLIsZero() bool {
	return x == PeerBehavior{}
}

func (x PeerBehavior) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_2); err != nil {
		return err
	}
	if x.Reachable {
		if err := enc.NextFieldValueBool("Reachable", vdl.BoolType, x.Reachable); err != nil {
			return err
		}
	}
	if x.Discoverable {
		if err := enc.NextFieldValueBool("Discoverable", vdl.BoolType, x.Discoverable); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *PeerBehavior) VDLRead(dec vdl.Decoder) error {
	*x = PeerBehavior{}
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
		case "Reachable":
			switch value, err := dec.ReadValueBool(); {
			case err != nil:
				return err
			default:
				x.Reachable = value
			}
		case "Discoverable":
			switch value, err := dec.ReadValueBool(); {
			case err != nil:
				return err
			default:
				x.Discoverable = value
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

//////////////////////////////////////////////////
// Error definitions

var (
	ErrInvalidAddress       = verror.Register("v.io/x/ref/runtime/protocols/vine.InvalidAddress", verror.NoRetry, "{1:}{2:} invalid vine address {3}, address must be of the form 'network/address/tag'")
	ErrAddressNotReachable  = verror.Register("v.io/x/ref/runtime/protocols/vine.AddressNotReachable", verror.NoRetry, "{1:}{2:} address {3} not reachable")
	ErrNoRegisteredProtocol = verror.Register("v.io/x/ref/runtime/protocols/vine.NoRegisteredProtocol", verror.NoRetry, "{1:}{2:} no registered protocol {3}")
	ErrCantAcceptFromTag    = verror.Register("v.io/x/ref/runtime/protocols/vine.CantAcceptFromTag", verror.NoRetry, "{1:}{2:} can't accept connection from tag {3}")
)

// NewErrInvalidAddress returns an error with the ErrInvalidAddress ID.
func NewErrInvalidAddress(ctx *context.T, address string) error {
	return verror.New(ErrInvalidAddress, ctx, address)
}

// NewErrAddressNotReachable returns an error with the ErrAddressNotReachable ID.
func NewErrAddressNotReachable(ctx *context.T, address string) error {
	return verror.New(ErrAddressNotReachable, ctx, address)
}

// NewErrNoRegisteredProtocol returns an error with the ErrNoRegisteredProtocol ID.
func NewErrNoRegisteredProtocol(ctx *context.T, protocol string) error {
	return verror.New(ErrNoRegisteredProtocol, ctx, protocol)
}

// NewErrCantAcceptFromTag returns an error with the ErrCantAcceptFromTag ID.
func NewErrCantAcceptFromTag(ctx *context.T, tag string) error {
	return verror.New(ErrCantAcceptFromTag, ctx, tag)
}

//////////////////////////////////////////////////
// Interface definitions

// VineClientMethods is the client interface
// containing Vine methods.
//
// Vine is the interface to a vine service that can dynamically change the network
// behavior of connection's on the vine service's process.
type VineClientMethods interface {
	// SetBehaviors sets the policy that the accepting vine service's process
	// will use on connections.
	// behaviors is a map from server tag to the desired connection behavior.
	// For example,
	//   client.SetBehaviors(map[PeerKey]PeerBehavior{PeerKey{"foo", "bar"}, PeerBehavior{Reachable: false}})
	// will cause all vine protocol dial calls from "foo" to "bar" to fail.
	SetBehaviors(_ *context.T, behaviors map[PeerKey]PeerBehavior, _ ...rpc.CallOpt) error
}

// VineClientStub adds universal methods to VineClientMethods.
type VineClientStub interface {
	VineClientMethods
	rpc.UniversalServiceMethods
}

// VineClient returns a client stub for Vine.
func VineClient(name string) VineClientStub {
	return implVineClientStub{name}
}

type implVineClientStub struct {
	name string
}

func (c implVineClientStub) SetBehaviors(ctx *context.T, i0 map[PeerKey]PeerBehavior, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "SetBehaviors", []interface{}{i0}, nil, opts...)
	return
}

// VineServerMethods is the interface a server writer
// implements for Vine.
//
// Vine is the interface to a vine service that can dynamically change the network
// behavior of connection's on the vine service's process.
type VineServerMethods interface {
	// SetBehaviors sets the policy that the accepting vine service's process
	// will use on connections.
	// behaviors is a map from server tag to the desired connection behavior.
	// For example,
	//   client.SetBehaviors(map[PeerKey]PeerBehavior{PeerKey{"foo", "bar"}, PeerBehavior{Reachable: false}})
	// will cause all vine protocol dial calls from "foo" to "bar" to fail.
	SetBehaviors(_ *context.T, _ rpc.ServerCall, behaviors map[PeerKey]PeerBehavior) error
}

// VineServerStubMethods is the server interface containing
// Vine methods, as expected by rpc.Server.
// There is no difference between this interface and VineServerMethods
// since there are no streaming methods.
type VineServerStubMethods VineServerMethods

// VineServerStub adds universal methods to VineServerStubMethods.
type VineServerStub interface {
	VineServerStubMethods
	// Describe the Vine interfaces.
	Describe__() []rpc.InterfaceDesc
}

// VineServer returns a server stub for Vine.
// It converts an implementation of VineServerMethods into
// an object that may be used by rpc.Server.
func VineServer(impl VineServerMethods) VineServerStub {
	stub := implVineServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implVineServerStub struct {
	impl VineServerMethods
	gs   *rpc.GlobState
}

func (s implVineServerStub) SetBehaviors(ctx *context.T, call rpc.ServerCall, i0 map[PeerKey]PeerBehavior) error {
	return s.impl.SetBehaviors(ctx, call, i0)
}

func (s implVineServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implVineServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{VineDesc}
}

// VineDesc describes the Vine interface.
var VineDesc rpc.InterfaceDesc = descVine

// descVine hides the desc to keep godoc clean.
var descVine = rpc.InterfaceDesc{
	Name:    "Vine",
	PkgPath: "v.io/x/ref/runtime/protocols/vine",
	Doc:     "// Vine is the interface to a vine service that can dynamically change the network\n// behavior of connection's on the vine service's process.",
	Methods: []rpc.MethodDesc{
		{
			Name: "SetBehaviors",
			Doc:  "// SetBehaviors sets the policy that the accepting vine service's process\n// will use on connections.\n// behaviors is a map from server tag to the desired connection behavior.\n// For example,\n//   client.SetBehaviors(map[PeerKey]PeerBehavior{PeerKey{\"foo\", \"bar\"}, PeerBehavior{Reachable: false}})\n// will cause all vine protocol dial calls from \"foo\" to \"bar\" to fail.",
			InArgs: []rpc.ArgDesc{
				{"behaviors", ``}, // map[PeerKey]PeerBehavior
			},
		},
	},
}

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1 *vdl.Type
	__VDLType_struct_2 *vdl.Type
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
	vdl.Register((*PeerKey)(nil))
	vdl.Register((*PeerBehavior)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*PeerKey)(nil)).Elem()
	__VDLType_struct_2 = vdl.TypeOf((*PeerBehavior)(nil)).Elem()

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidAddress.ID), "{1:}{2:} invalid vine address {3}, address must be of the form 'network/address/tag'")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrAddressNotReachable.ID), "{1:}{2:} address {3} not reachable")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoRegisteredProtocol.ID), "{1:}{2:} no registered protocol {3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCantAcceptFromTag.ID), "{1:}{2:} can't accept connection from tag {3}")

	return struct{}{}
}
