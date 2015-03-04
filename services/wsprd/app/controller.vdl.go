// This file was auto-generated by the veyron vdl tool.
// Source: controller.vdl

package app

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/ipc"

	// VDL user imports
	"time"
	"v.io/v23/security"
	"v.io/v23/vdlroot/signature"
	_ "v.io/v23/vdlroot/time"
)

// ControllerClientMethods is the client interface
// containing Controller methods.
type ControllerClientMethods interface {
	// Serve instructs WSPR to start listening for calls on behalf
	// of a javascript server.
	Serve(ctx *context.T, name string, serverId uint32, opts ...ipc.CallOpt) error
	// Stop instructs WSPR to stop listening for calls for the
	// given javascript server.
	Stop(ctx *context.T, serverId uint32, opts ...ipc.CallOpt) error
	// AddName adds a published name to an existing server.
	AddName(ctx *context.T, serverId uint32, name string, opts ...ipc.CallOpt) error
	// RemoveName removes a published name from an existing server.
	RemoveName(ctx *context.T, serverId uint32, name string, opts ...ipc.CallOpt) error
	// UnlinkJSBlessings removes the given blessings from the blessings store.
	UnlinkJSBlessings(ctx *context.T, handle int32, opts ...ipc.CallOpt) error
	// BlessPublicKey creates a new blessing.
	BlessPublicKey(ctx *context.T, fromHandle int32, caveats []security.Caveat, durationMs time.Duration, extension string, opts ...ipc.CallOpt) (handle int32, publicKey string, err error)
	// CreateBlessings creates a new principal self-blessed with the given extension.
	CreateBlessings(ctx *context.T, extension string, opts ...ipc.CallOpt) (handle int32, publicKey string, err error)
	// RemoteBlessings fetches the remote blessings for a given name and method.
	RemoteBlessings(ctx *context.T, name string, method string, opts ...ipc.CallOpt) ([]string, error)
	// Signature fetches the signature for a given name.
	Signature(ctx *context.T, name string, opts ...ipc.CallOpt) ([]signature.Interface, error)
}

// ControllerClientStub adds universal methods to ControllerClientMethods.
type ControllerClientStub interface {
	ControllerClientMethods
	ipc.UniversalServiceMethods
}

// ControllerClient returns a client stub for Controller.
func ControllerClient(name string, opts ...ipc.BindOpt) ControllerClientStub {
	var client ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(ipc.Client); ok {
			client = clientOpt
		}
	}
	return implControllerClientStub{name, client}
}

type implControllerClientStub struct {
	name   string
	client ipc.Client
}

func (c implControllerClientStub) c(ctx *context.T) ipc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implControllerClientStub) Serve(ctx *context.T, i0 string, i1 uint32, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Serve", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) Stop(ctx *context.T, i0 uint32, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Stop", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) AddName(ctx *context.T, i0 uint32, i1 string, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "AddName", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) RemoveName(ctx *context.T, i0 uint32, i1 string, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "RemoveName", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) UnlinkJSBlessings(ctx *context.T, i0 int32, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "UnlinkJSBlessings", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) BlessPublicKey(ctx *context.T, i0 int32, i1 []security.Caveat, i2 time.Duration, i3 string, opts ...ipc.CallOpt) (o0 int32, o1 string, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessPublicKey", []interface{}{i0, i1, i2, i3}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0, &o1)
	return
}

func (c implControllerClientStub) CreateBlessings(ctx *context.T, i0 string, opts ...ipc.CallOpt) (o0 int32, o1 string, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "CreateBlessings", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0, &o1)
	return
}

func (c implControllerClientStub) RemoteBlessings(ctx *context.T, i0 string, i1 string, opts ...ipc.CallOpt) (o0 []string, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "RemoteBlessings", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implControllerClientStub) Signature(ctx *context.T, i0 string, opts ...ipc.CallOpt) (o0 []signature.Interface, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

// ControllerServerMethods is the interface a server writer
// implements for Controller.
type ControllerServerMethods interface {
	// Serve instructs WSPR to start listening for calls on behalf
	// of a javascript server.
	Serve(call ipc.ServerCall, name string, serverId uint32) error
	// Stop instructs WSPR to stop listening for calls for the
	// given javascript server.
	Stop(call ipc.ServerCall, serverId uint32) error
	// AddName adds a published name to an existing server.
	AddName(call ipc.ServerCall, serverId uint32, name string) error
	// RemoveName removes a published name from an existing server.
	RemoveName(call ipc.ServerCall, serverId uint32, name string) error
	// UnlinkJSBlessings removes the given blessings from the blessings store.
	UnlinkJSBlessings(call ipc.ServerCall, handle int32) error
	// BlessPublicKey creates a new blessing.
	BlessPublicKey(call ipc.ServerCall, fromHandle int32, caveats []security.Caveat, durationMs time.Duration, extension string) (handle int32, publicKey string, err error)
	// CreateBlessings creates a new principal self-blessed with the given extension.
	CreateBlessings(call ipc.ServerCall, extension string) (handle int32, publicKey string, err error)
	// RemoteBlessings fetches the remote blessings for a given name and method.
	RemoteBlessings(call ipc.ServerCall, name string, method string) ([]string, error)
	// Signature fetches the signature for a given name.
	Signature(call ipc.ServerCall, name string) ([]signature.Interface, error)
}

// ControllerServerStubMethods is the server interface containing
// Controller methods, as expected by ipc.Server.
// There is no difference between this interface and ControllerServerMethods
// since there are no streaming methods.
type ControllerServerStubMethods ControllerServerMethods

// ControllerServerStub adds universal methods to ControllerServerStubMethods.
type ControllerServerStub interface {
	ControllerServerStubMethods
	// Describe the Controller interfaces.
	Describe__() []ipc.InterfaceDesc
}

// ControllerServer returns a server stub for Controller.
// It converts an implementation of ControllerServerMethods into
// an object that may be used by ipc.Server.
func ControllerServer(impl ControllerServerMethods) ControllerServerStub {
	stub := implControllerServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implControllerServerStub struct {
	impl ControllerServerMethods
	gs   *ipc.GlobState
}

func (s implControllerServerStub) Serve(call ipc.ServerCall, i0 string, i1 uint32) error {
	return s.impl.Serve(call, i0, i1)
}

func (s implControllerServerStub) Stop(call ipc.ServerCall, i0 uint32) error {
	return s.impl.Stop(call, i0)
}

func (s implControllerServerStub) AddName(call ipc.ServerCall, i0 uint32, i1 string) error {
	return s.impl.AddName(call, i0, i1)
}

func (s implControllerServerStub) RemoveName(call ipc.ServerCall, i0 uint32, i1 string) error {
	return s.impl.RemoveName(call, i0, i1)
}

func (s implControllerServerStub) UnlinkJSBlessings(call ipc.ServerCall, i0 int32) error {
	return s.impl.UnlinkJSBlessings(call, i0)
}

func (s implControllerServerStub) BlessPublicKey(call ipc.ServerCall, i0 int32, i1 []security.Caveat, i2 time.Duration, i3 string) (int32, string, error) {
	return s.impl.BlessPublicKey(call, i0, i1, i2, i3)
}

func (s implControllerServerStub) CreateBlessings(call ipc.ServerCall, i0 string) (int32, string, error) {
	return s.impl.CreateBlessings(call, i0)
}

func (s implControllerServerStub) RemoteBlessings(call ipc.ServerCall, i0 string, i1 string) ([]string, error) {
	return s.impl.RemoteBlessings(call, i0, i1)
}

func (s implControllerServerStub) Signature(call ipc.ServerCall, i0 string) ([]signature.Interface, error) {
	return s.impl.Signature(call, i0)
}

func (s implControllerServerStub) Globber() *ipc.GlobState {
	return s.gs
}

func (s implControllerServerStub) Describe__() []ipc.InterfaceDesc {
	return []ipc.InterfaceDesc{ControllerDesc}
}

// ControllerDesc describes the Controller interface.
var ControllerDesc ipc.InterfaceDesc = descController

// descController hides the desc to keep godoc clean.
var descController = ipc.InterfaceDesc{
	Name:    "Controller",
	PkgPath: "v.io/x/ref/services/wsprd/app",
	Methods: []ipc.MethodDesc{
		{
			Name: "Serve",
			Doc:  "// Serve instructs WSPR to start listening for calls on behalf\n// of a javascript server.",
			InArgs: []ipc.ArgDesc{
				{"name", ``},     // string
				{"serverId", ``}, // uint32
			},
		},
		{
			Name: "Stop",
			Doc:  "// Stop instructs WSPR to stop listening for calls for the\n// given javascript server.",
			InArgs: []ipc.ArgDesc{
				{"serverId", ``}, // uint32
			},
		},
		{
			Name: "AddName",
			Doc:  "// AddName adds a published name to an existing server.",
			InArgs: []ipc.ArgDesc{
				{"serverId", ``}, // uint32
				{"name", ``},     // string
			},
		},
		{
			Name: "RemoveName",
			Doc:  "// RemoveName removes a published name from an existing server.",
			InArgs: []ipc.ArgDesc{
				{"serverId", ``}, // uint32
				{"name", ``},     // string
			},
		},
		{
			Name: "UnlinkJSBlessings",
			Doc:  "// UnlinkJSBlessings removes the given blessings from the blessings store.",
			InArgs: []ipc.ArgDesc{
				{"handle", ``}, // int32
			},
		},
		{
			Name: "BlessPublicKey",
			Doc:  "// BlessPublicKey creates a new blessing.",
			InArgs: []ipc.ArgDesc{
				{"fromHandle", ``}, // int32
				{"caveats", ``},    // []security.Caveat
				{"durationMs", ``}, // time.Duration
				{"extension", ``},  // string
			},
			OutArgs: []ipc.ArgDesc{
				{"handle", ``},    // int32
				{"publicKey", ``}, // string
			},
		},
		{
			Name: "CreateBlessings",
			Doc:  "// CreateBlessings creates a new principal self-blessed with the given extension.",
			InArgs: []ipc.ArgDesc{
				{"extension", ``}, // string
			},
			OutArgs: []ipc.ArgDesc{
				{"handle", ``},    // int32
				{"publicKey", ``}, // string
			},
		},
		{
			Name: "RemoteBlessings",
			Doc:  "// RemoteBlessings fetches the remote blessings for a given name and method.",
			InArgs: []ipc.ArgDesc{
				{"name", ``},   // string
				{"method", ``}, // string
			},
			OutArgs: []ipc.ArgDesc{
				{"", ``}, // []string
			},
		},
		{
			Name: "Signature",
			Doc:  "// Signature fetches the signature for a given name.",
			InArgs: []ipc.ArgDesc{
				{"name", ``}, // string
			},
			OutArgs: []ipc.ArgDesc{
				{"", ``}, // []signature.Interface
			},
		},
	},
}
