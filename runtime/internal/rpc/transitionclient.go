// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"v.io/v23/context"
	"v.io/v23/flow"
	"v.io/v23/flow/message"
	"v.io/v23/namespace"
	"v.io/v23/rpc"
	"v.io/v23/verror"
	"v.io/x/ref/runtime/internal/rpc/stream"
)

type transitionClient struct {
	c, xc rpc.Client
}

var _ = rpc.Client((*transitionClient)(nil))

func NewTransitionClient(ctx *context.T, streamMgr stream.Manager, flowMgr flow.Manager, ns namespace.T, opts ...rpc.ClientOpt) (rpc.Client, error) {
	var err error
	ret := &transitionClient{}
	if ret.xc, err = NewXClient(ctx, flowMgr, ns, opts...); err != nil {
		return nil, err
	}
	if ret.c, err = DeprecatedNewClient(streamMgr, ns, opts...); err != nil {
		ret.xc.Close()
		return nil, err
	}
	return ret, nil
}

func (t *transitionClient) StartCall(ctx *context.T, name, method string, args []interface{}, opts ...rpc.CallOpt) (rpc.ClientCall, error) {
	call, err := t.xc.StartCall(ctx, name, method, args, opts...)
	if verror.ErrorID(err) == message.ErrWrongProtocol.ID {
		call, err = t.c.StartCall(ctx, name, method, args, opts...)
	}
	return call, err
}

func (t *transitionClient) Call(ctx *context.T, name, method string, in, out []interface{}, opts ...rpc.CallOpt) error {
	err := t.xc.Call(ctx, name, method, in, out, opts...)
	if verror.ErrorID(err) == message.ErrWrongProtocol.ID {
		err = t.c.Call(ctx, name, method, in, out, opts...)
	}
	return err
}

func (t *transitionClient) Close() {
	t.xc.Close()
	t.c.Close()
}
