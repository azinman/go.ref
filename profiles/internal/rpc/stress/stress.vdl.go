// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: stress.vdl

package stress

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security/access"
)

type Arg struct {
	ABool        bool
	AInt64       int64
	AListOfBytes []byte
}

func (Arg) __VDLReflect(struct {
	Name string "v.io/x/ref/profiles/internal/rpc/stress.Arg"
}) {
}

type Stats struct {
	SumCount       uint64
	SumStreamCount uint64
}

func (Stats) __VDLReflect(struct {
	Name string "v.io/x/ref/profiles/internal/rpc/stress.Stats"
}) {
}

func init() {
	vdl.Register((*Arg)(nil))
	vdl.Register((*Stats)(nil))
}

// StressClientMethods is the client interface
// containing Stress methods.
type StressClientMethods interface {
	// Do returns the checksum of the payload that it receives.
	Sum(ctx *context.T, arg Arg, opts ...rpc.CallOpt) ([]byte, error)
	// DoStream returns the checksum of the payload that it receives via the stream.
	SumStream(*context.T, ...rpc.CallOpt) (StressSumStreamClientCall, error)
	// GetStats returns the stats on the calls that the server received.
	GetStats(*context.T, ...rpc.CallOpt) (Stats, error)
	// Stop stops the server.
	Stop(*context.T, ...rpc.CallOpt) error
}

// StressClientStub adds universal methods to StressClientMethods.
type StressClientStub interface {
	StressClientMethods
	rpc.UniversalServiceMethods
}

// StressClient returns a client stub for Stress.
func StressClient(name string, opts ...rpc.BindOpt) StressClientStub {
	var client rpc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(rpc.Client); ok {
			client = clientOpt
		}
	}
	return implStressClientStub{name, client}
}

type implStressClientStub struct {
	name   string
	client rpc.Client
}

func (c implStressClientStub) c(ctx *context.T) rpc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implStressClientStub) Sum(ctx *context.T, i0 Arg, opts ...rpc.CallOpt) (o0 []byte, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Sum", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implStressClientStub) SumStream(ctx *context.T, opts ...rpc.CallOpt) (ocall StressSumStreamClientCall, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "SumStream", nil, opts...); err != nil {
		return
	}
	ocall = &implStressSumStreamClientCall{ClientCall: call}
	return
}

func (c implStressClientStub) GetStats(ctx *context.T, opts ...rpc.CallOpt) (o0 Stats, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "GetStats", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implStressClientStub) Stop(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Stop", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

// StressSumStreamClientStream is the client stream for Stress.SumStream.
type StressSumStreamClientStream interface {
	// RecvStream returns the receiver side of the Stress.SumStream client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Stress.SumStream client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item Arg) error
		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.
		// This is an optional call - e.g. a client might call Close if it
		// needs to continue receiving items from the server after it's
		// done sending.  Returns errors encountered while closing, or if
		// Close is called after the stream has been canceled.  Like Send,
		// blocks if there is no buffer space available.
		Close() error
	}
}

// StressSumStreamClientCall represents the call returned from Stress.SumStream.
type StressSumStreamClientCall interface {
	StressSumStreamClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implStressSumStreamClientCall struct {
	rpc.ClientCall
	valRecv []byte
	errRecv error
}

func (c *implStressSumStreamClientCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implStressSumStreamClientCallRecv{c}
}

type implStressSumStreamClientCallRecv struct {
	c *implStressSumStreamClientCall
}

func (c implStressSumStreamClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implStressSumStreamClientCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implStressSumStreamClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implStressSumStreamClientCall) SendStream() interface {
	Send(item Arg) error
	Close() error
} {
	return implStressSumStreamClientCallSend{c}
}

type implStressSumStreamClientCallSend struct {
	c *implStressSumStreamClientCall
}

func (c implStressSumStreamClientCallSend) Send(item Arg) error {
	return c.c.Send(item)
}
func (c implStressSumStreamClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implStressSumStreamClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// StressServerMethods is the interface a server writer
// implements for Stress.
type StressServerMethods interface {
	// Do returns the checksum of the payload that it receives.
	Sum(call rpc.ServerCall, arg Arg) ([]byte, error)
	// DoStream returns the checksum of the payload that it receives via the stream.
	SumStream(StressSumStreamServerCall) error
	// GetStats returns the stats on the calls that the server received.
	GetStats(rpc.ServerCall) (Stats, error)
	// Stop stops the server.
	Stop(rpc.ServerCall) error
}

// StressServerStubMethods is the server interface containing
// Stress methods, as expected by rpc.Server.
// The only difference between this interface and StressServerMethods
// is the streaming methods.
type StressServerStubMethods interface {
	// Do returns the checksum of the payload that it receives.
	Sum(call rpc.ServerCall, arg Arg) ([]byte, error)
	// DoStream returns the checksum of the payload that it receives via the stream.
	SumStream(*StressSumStreamServerCallStub) error
	// GetStats returns the stats on the calls that the server received.
	GetStats(rpc.ServerCall) (Stats, error)
	// Stop stops the server.
	Stop(rpc.ServerCall) error
}

// StressServerStub adds universal methods to StressServerStubMethods.
type StressServerStub interface {
	StressServerStubMethods
	// Describe the Stress interfaces.
	Describe__() []rpc.InterfaceDesc
}

// StressServer returns a server stub for Stress.
// It converts an implementation of StressServerMethods into
// an object that may be used by rpc.Server.
func StressServer(impl StressServerMethods) StressServerStub {
	stub := implStressServerStub{
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

type implStressServerStub struct {
	impl StressServerMethods
	gs   *rpc.GlobState
}

func (s implStressServerStub) Sum(call rpc.ServerCall, i0 Arg) ([]byte, error) {
	return s.impl.Sum(call, i0)
}

func (s implStressServerStub) SumStream(call *StressSumStreamServerCallStub) error {
	return s.impl.SumStream(call)
}

func (s implStressServerStub) GetStats(call rpc.ServerCall) (Stats, error) {
	return s.impl.GetStats(call)
}

func (s implStressServerStub) Stop(call rpc.ServerCall) error {
	return s.impl.Stop(call)
}

func (s implStressServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implStressServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{StressDesc}
}

// StressDesc describes the Stress interface.
var StressDesc rpc.InterfaceDesc = descStress

// descStress hides the desc to keep godoc clean.
var descStress = rpc.InterfaceDesc{
	Name:    "Stress",
	PkgPath: "v.io/x/ref/profiles/internal/rpc/stress",
	Methods: []rpc.MethodDesc{
		{
			Name: "Sum",
			Doc:  "// Do returns the checksum of the payload that it receives.",
			InArgs: []rpc.ArgDesc{
				{"arg", ``}, // Arg
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []byte
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "SumStream",
			Doc:  "// DoStream returns the checksum of the payload that it receives via the stream.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "GetStats",
			Doc:  "// GetStats returns the stats on the calls that the server received.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // Stats
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "Stop",
			Doc:  "// Stop stops the server.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Admin"))},
		},
	},
}

// StressSumStreamServerStream is the server stream for Stress.SumStream.
type StressSumStreamServerStream interface {
	// RecvStream returns the receiver side of the Stress.SumStream server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() Arg
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Stress.SumStream server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item []byte) error
	}
}

// StressSumStreamServerCall represents the context passed to Stress.SumStream.
type StressSumStreamServerCall interface {
	rpc.ServerCall
	StressSumStreamServerStream
}

// StressSumStreamServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements StressSumStreamServerCall.
type StressSumStreamServerCallStub struct {
	rpc.StreamServerCall
	valRecv Arg
	errRecv error
}

// Init initializes StressSumStreamServerCallStub from rpc.StreamServerCall.
func (s *StressSumStreamServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Stress.SumStream server stream.
func (s *StressSumStreamServerCallStub) RecvStream() interface {
	Advance() bool
	Value() Arg
	Err() error
} {
	return implStressSumStreamServerCallRecv{s}
}

type implStressSumStreamServerCallRecv struct {
	s *StressSumStreamServerCallStub
}

func (s implStressSumStreamServerCallRecv) Advance() bool {
	s.s.valRecv = Arg{}
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implStressSumStreamServerCallRecv) Value() Arg {
	return s.s.valRecv
}
func (s implStressSumStreamServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Stress.SumStream server stream.
func (s *StressSumStreamServerCallStub) SendStream() interface {
	Send(item []byte) error
} {
	return implStressSumStreamServerCallSend{s}
}

type implStressSumStreamServerCallSend struct {
	s *StressSumStreamServerCallStub
}

func (s implStressSumStreamServerCallSend) Send(item []byte) error {
	return s.s.Send(item)
}
