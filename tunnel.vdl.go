// This file was auto-generated by the veyron vdl tool.
// Source: tunnel.vdl

package tunnel

import (
	"veyron2/security"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdl "veyron2/vdl"
	_gen_wiretype "veyron2/wiretype"
)

type ShellOpts struct {
	UsePty      bool     // Whether to open a pseudo-terminal
	Environment []string // Environment variables to pass to the remote shell.
	Rows        uint32   // Window size.
	Cols        uint32
}

type ClientShellPacket struct {
	// Bytes going to the shell's stdin.
	Stdin []byte
	// A dynamic update of the window size. The default value of 0 means no-change.
	Rows uint32
	Cols uint32
}

type ServerShellPacket struct {
	// Bytes coming from the shell's stdout.
	Stdout []byte
	// Bytes coming from the shell's stderr.
	Stderr []byte
}

// Tunnel is the interface the client binds and uses.
// Tunnel_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Tunnel_ExcludingUniversal interface {
	// The Forward method is used for network forwarding. All the data sent over
	// the byte stream is forwarded to the requested network address and all the
	// data received from that network connection is sent back in the reply
	// stream.
	Forward(ctx _gen_ipc.Context, network string, address string, opts ..._gen_ipc.CallOpt) (reply TunnelForwardStream, err error)
	// The Shell method is used to either run shell commands remotely, or to open
	// an interactive shell. The data received over the byte stream is sent to the
	// shell's stdin, and the data received from the shell's stdout and stderr is
	// sent back in the reply stream. It returns the exit status of the shell
	// command.
	Shell(ctx _gen_ipc.Context, command string, shellOpts ShellOpts, opts ..._gen_ipc.CallOpt) (reply TunnelShellStream, err error)
}
type Tunnel interface {
	_gen_ipc.UniversalServiceMethods
	Tunnel_ExcludingUniversal
}

// TunnelService is the interface the server implements.
type TunnelService interface {

	// The Forward method is used for network forwarding. All the data sent over
	// the byte stream is forwarded to the requested network address and all the
	// data received from that network connection is sent back in the reply
	// stream.
	Forward(context _gen_ipc.ServerContext, network string, address string, stream TunnelServiceForwardStream) (err error)
	// The Shell method is used to either run shell commands remotely, or to open
	// an interactive shell. The data received over the byte stream is sent to the
	// shell's stdin, and the data received from the shell's stdout and stderr is
	// sent back in the reply stream. It returns the exit status of the shell
	// command.
	Shell(context _gen_ipc.ServerContext, command string, shellOpts ShellOpts, stream TunnelServiceShellStream) (reply int32, err error)
}

// TunnelForwardStream is the interface for streaming responses of the method
// Forward in the service interface Tunnel.
type TunnelForwardStream interface {

	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item []byte) error

	// CloseSend indicates to the server that no more items will be sent; server
	// Recv calls will receive io.EOF after all sent items.  Subsequent calls to
	// Send on the client will fail.  This is an optional call - it's used by
	// streaming clients that need the server to receive the io.EOF terminator.
	CloseSend() error

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item []byte, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the TunnelForwardStream interface that is not exported.
type implTunnelForwardStream struct {
	clientCall _gen_ipc.Call
}

func (c *implTunnelForwardStream) Send(item []byte) error {
	return c.clientCall.Send(item)
}

func (c *implTunnelForwardStream) CloseSend() error {
	return c.clientCall.CloseSend()
}

func (c *implTunnelForwardStream) Recv() (item []byte, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implTunnelForwardStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implTunnelForwardStream) Cancel() {
	c.clientCall.Cancel()
}

// TunnelServiceForwardStream is the interface for streaming responses of the method
// Forward in the service interface Tunnel.
type TunnelServiceForwardStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item []byte) error

	// Recv fills itemptr with the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item []byte, err error)
}

// Implementation of the TunnelServiceForwardStream interface that is not exported.
type implTunnelServiceForwardStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implTunnelServiceForwardStream) Send(item []byte) error {
	return s.serverCall.Send(item)
}

func (s *implTunnelServiceForwardStream) Recv() (item []byte, err error) {
	err = s.serverCall.Recv(&item)
	return
}

// TunnelShellStream is the interface for streaming responses of the method
// Shell in the service interface Tunnel.
type TunnelShellStream interface {

	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item ClientShellPacket) error

	// CloseSend indicates to the server that no more items will be sent; server
	// Recv calls will receive io.EOF after all sent items.  Subsequent calls to
	// Send on the client will fail.  This is an optional call - it's used by
	// streaming clients that need the server to receive the io.EOF terminator.
	CloseSend() error

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item ServerShellPacket, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (reply int32, err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the TunnelShellStream interface that is not exported.
type implTunnelShellStream struct {
	clientCall _gen_ipc.Call
}

func (c *implTunnelShellStream) Send(item ClientShellPacket) error {
	return c.clientCall.Send(item)
}

func (c *implTunnelShellStream) CloseSend() error {
	return c.clientCall.CloseSend()
}

func (c *implTunnelShellStream) Recv() (item ServerShellPacket, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implTunnelShellStream) Finish() (reply int32, err error) {
	if ierr := c.clientCall.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implTunnelShellStream) Cancel() {
	c.clientCall.Cancel()
}

// TunnelServiceShellStream is the interface for streaming responses of the method
// Shell in the service interface Tunnel.
type TunnelServiceShellStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item ServerShellPacket) error

	// Recv fills itemptr with the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item ClientShellPacket, err error)
}

// Implementation of the TunnelServiceShellStream interface that is not exported.
type implTunnelServiceShellStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implTunnelServiceShellStream) Send(item ServerShellPacket) error {
	return s.serverCall.Send(item)
}

func (s *implTunnelServiceShellStream) Recv() (item ClientShellPacket, err error) {
	err = s.serverCall.Recv(&item)
	return
}

// BindTunnel returns the client stub implementing the Tunnel
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindTunnel(name string, opts ..._gen_ipc.BindOpt) (Tunnel, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_veyron2.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdl.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdl.ErrTooManyOptionsToBind
	}
	stub := &clientStubTunnel{client: client, name: name}

	return stub, nil
}

// NewServerTunnel creates a new server stub.
//
// It takes a regular server implementing the TunnelService
// interface, and returns a new server stub.
func NewServerTunnel(server TunnelService) interface{} {
	return &ServerStubTunnel{
		service: server,
	}
}

// clientStubTunnel implements Tunnel.
type clientStubTunnel struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubTunnel) Forward(ctx _gen_ipc.Context, network string, address string, opts ..._gen_ipc.CallOpt) (reply TunnelForwardStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Forward", []interface{}{network, address}, opts...); err != nil {
		return
	}
	reply = &implTunnelForwardStream{clientCall: call}
	return
}

func (__gen_c *clientStubTunnel) Shell(ctx _gen_ipc.Context, command string, shellOpts ShellOpts, opts ..._gen_ipc.CallOpt) (reply TunnelShellStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Shell", []interface{}{command, shellOpts}, opts...); err != nil {
		return
	}
	reply = &implTunnelShellStream{clientCall: call}
	return
}

func (__gen_c *clientStubTunnel) UnresolveStep(ctx _gen_ipc.Context, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTunnel) Signature(ctx _gen_ipc.Context, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTunnel) GetMethodTags(ctx _gen_ipc.Context, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubTunnel wraps a server that implements
// TunnelService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubTunnel struct {
	service TunnelService
}

func (__gen_s *ServerStubTunnel) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Forward":
		return []interface{}{security.Label(4)}, nil
	case "Shell":
		return []interface{}{security.Label(4)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubTunnel) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Forward"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "network", Type: 3},
			{Name: "address", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
		InStream:  67,
		OutStream: 67,
	}
	result.Methods["Shell"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "command", Type: 3},
			{Name: "shellOpts", Type: 68},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 36},
			{Name: "", Type: 65},
		},
		InStream:  69,
		OutStream: 70,
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x2, Name: "UsePty"},
				_gen_wiretype.FieldType{Type: 0x3d, Name: "Environment"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "Rows"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "Cols"},
			},
			"veyron/examples/tunnel.ShellOpts", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x43, Name: "Stdin"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "Rows"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "Cols"},
			},
			"veyron/examples/tunnel.ClientShellPacket", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x43, Name: "Stdout"},
				_gen_wiretype.FieldType{Type: 0x43, Name: "Stderr"},
			},
			"veyron/examples/tunnel.ServerShellPacket", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubTunnel) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubTunnel) Forward(call _gen_ipc.ServerCall, network string, address string) (err error) {
	stream := &implTunnelServiceForwardStream{serverCall: call}
	err = __gen_s.service.Forward(call, network, address, stream)
	return
}

func (__gen_s *ServerStubTunnel) Shell(call _gen_ipc.ServerCall, command string, shellOpts ShellOpts) (reply int32, err error) {
	stream := &implTunnelServiceShellStream{serverCall: call}
	reply, err = __gen_s.service.Shell(call, command, shellOpts, stream)
	return
}
