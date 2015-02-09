// This file was auto-generated by the veyron vdl tool.
// Source: protocol.vdl

package proxy

import (
	// VDL system imports
	"v.io/core/veyron2/vdl"
)

// Request is the message sent by a server to request that the proxy route
// traffic intended for the server's RoutingID to the network connection
// between the server and the proxy.
type Request struct {
}

func (Request) __VDLReflect(struct {
	Name string "v.io/core/veyron/runtimes/google/ipc/stream/proxy.Request"
}) {
}

// Response is sent by the proxy to the server after processing Request.
type Response struct {
	// Error is a description of why the proxy refused to proxy the server.
	// A nil error indicates that the proxy will route traffic to the server.
	Error error
	// Endpoint is the string representation of an endpoint that can be
	// used to communicate with the server through the proxy.
	Endpoint string
}

func (Response) __VDLReflect(struct {
	Name string "v.io/core/veyron/runtimes/google/ipc/stream/proxy.Response"
}) {
}

func init() {
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))
}
