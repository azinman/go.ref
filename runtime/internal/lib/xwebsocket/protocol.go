// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !nacl

package xwebsocket

import (
	"net"
	"net/http"
	"net/url"
	"time"

	"v.io/x/ref/runtime/internal/lib/tcputil"

	"github.com/gorilla/websocket"

	"v.io/v23/context"
	"v.io/v23/flow"
)

// TODO(jhahn): Figure out a way for this mapping to be shared.
var mapWebSocketToTCP = map[string]string{"ws": "tcp", "ws4": "tcp4", "ws6": "tcp6", "wsh": "tcp", "wsh4": "tcp4", "wsh6": "tcp6", "tcp": "tcp", "tcp4": "tcp4", "tcp6": "tcp6"}

const bufferSize = 4096

type WS struct{}

func (WS) Dial(ctx *context.T, protocol, address string, timeout time.Duration) (flow.MsgReadWriteCloser, error) {
	var deadline time.Time
	if timeout > 0 {
		deadline = time.Now().Add(timeout)
	}
	tcp := mapWebSocketToTCP[protocol]
	conn, err := net.DialTimeout(tcp, address, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetReadDeadline(deadline)
	if err := tcputil.EnableTCPKeepAlive(conn); err != nil {
		return nil, err
	}
	u, err := url.Parse("ws://" + address)
	if err != nil {
		return nil, err
	}
	ws, _, err := websocket.NewClient(conn, u, http.Header{}, bufferSize, bufferSize)
	if err != nil {
		return nil, err
	}
	var zero time.Time
	conn.SetDeadline(zero)
	return WebsocketConn(ws), nil
}

func (WS) Resolve(ctx *context.T, protocol, address string) (string, string, error) {
	tcp := mapWebSocketToTCP[protocol]
	tcpAddr, err := net.ResolveTCPAddr(tcp, address)
	if err != nil {
		return "", "", err
	}
	return "ws", tcpAddr.String(), nil
}

func (WS) Listen(ctx *context.T, protocol, address string) (flow.MsgListener, error) {
	return listener(protocol, address, false)
}
