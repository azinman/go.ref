// This file was auto-generated by the veyron vdl tool.
// Source: boxes.vdl

// Boxes is an android app that uses veyron to share views
// between peer devices.
package boxes

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// Box describes the name and co-ordinates of a given box that
// is displayed in the View of a peer device.
type Box struct {
	// DeviceID that generated the box
	DeviceId string
	// BoxId is a unique name for a box
	BoxId string
	// Points are the co-ordinates of a given box
	Points [4]float32
}

// BoxSignalling allows peers to rendezvous with each other
// BoxSignalling is the interface the client binds and uses.
// BoxSignalling_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type BoxSignalling_ExcludingUniversal interface {
	// Add endpoint information to the signalling server.
	Add(ctx _gen_context.T, Endpoint string, opts ..._gen_ipc.CallOpt) (err error)
	// Get endpoint information about a peer.
	Get(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply string, err error)
}
type BoxSignalling interface {
	_gen_ipc.UniversalServiceMethods
	BoxSignalling_ExcludingUniversal
}

// BoxSignallingService is the interface the server implements.
type BoxSignallingService interface {

	// Add endpoint information to the signalling server.
	Add(context _gen_ipc.ServerContext, Endpoint string) (err error)
	// Get endpoint information about a peer.
	Get(context _gen_ipc.ServerContext) (reply string, err error)
}

// BindBoxSignalling returns the client stub implementing the BoxSignalling
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindBoxSignalling(name string, opts ..._gen_ipc.BindOpt) (BoxSignalling, error) {
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
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubBoxSignalling{client: client, name: name}

	return stub, nil
}

// NewServerBoxSignalling creates a new server stub.
//
// It takes a regular server implementing the BoxSignallingService
// interface, and returns a new server stub.
func NewServerBoxSignalling(server BoxSignallingService) interface{} {
	return &ServerStubBoxSignalling{
		service: server,
	}
}

// clientStubBoxSignalling implements BoxSignalling.
type clientStubBoxSignalling struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubBoxSignalling) Add(ctx _gen_context.T, Endpoint string, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Add", []interface{}{Endpoint}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBoxSignalling) Get(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Get", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBoxSignalling) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBoxSignalling) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBoxSignalling) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubBoxSignalling wraps a server that implements
// BoxSignallingService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubBoxSignalling struct {
	service BoxSignallingService
}

func (__gen_s *ServerStubBoxSignalling) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Add":
		return []interface{}{}, nil
	case "Get":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubBoxSignalling) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Add"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Endpoint", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "Err", Type: 65},
		},
	}
	result.Methods["Get"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "Endpoint", Type: 3},
			{Name: "Err", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubBoxSignalling) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubBoxSignalling) Add(call _gen_ipc.ServerCall, Endpoint string) (err error) {
	err = __gen_s.service.Add(call, Endpoint)
	return
}

func (__gen_s *ServerStubBoxSignalling) Get(call _gen_ipc.ServerCall) (reply string, err error) {
	reply, err = __gen_s.service.Get(call)
	return
}

// DrawInterface enables adding a box on another peer
// DrawInterface is the interface the client binds and uses.
// DrawInterface_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type DrawInterface_ExcludingUniversal interface {
	// Draw is used to send/receive a stream of boxes to another peer
	Draw(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply DrawInterfaceDrawStream, err error)
	// SyncBoxes is used to setup a sync service over store to send/receive
	// boxes to another peer
	SyncBoxes(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
}
type DrawInterface interface {
	_gen_ipc.UniversalServiceMethods
	DrawInterface_ExcludingUniversal
}

// DrawInterfaceService is the interface the server implements.
type DrawInterfaceService interface {

	// Draw is used to send/receive a stream of boxes to another peer
	Draw(context _gen_ipc.ServerContext, stream DrawInterfaceServiceDrawStream) (err error)
	// SyncBoxes is used to setup a sync service over store to send/receive
	// boxes to another peer
	SyncBoxes(context _gen_ipc.ServerContext) (err error)
}

// DrawInterfaceDrawStream is the interface for streaming responses of the method
// Draw in the service interface DrawInterface.
type DrawInterfaceDrawStream interface {

	// Send places the item onto the output stream, blocking if there is no
	// buffer space available.  Calls to Send after having called CloseSend
	// or Cancel will fail.  Any blocked Send calls will be unblocked upon
	// calling Cancel.
	Send(item Box) error

	// CloseSend indicates to the server that no more items will be sent;
	// server Recv calls will receive io.EOF after all sent items.  This is
	// an optional call - it's used by streaming clients that need the
	// server to receive the io.EOF terminator before the client calls
	// Finish (for example, if the client needs to continue receiving items
	// from the server after having finished sending).
	// Calls to CloseSend after having called Cancel will fail.
	// Like Send, CloseSend blocks when there's no buffer space available.
	CloseSend() error

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of
	// input.
	Recv() (item Box, err error)

	// Finish performs the equivalent of CloseSend, then blocks until the server
	// is done, and returns the positional return values for call.
	//
	// If Cancel has been called, Finish will return immediately; the output of
	// Finish could either be an error signalling cancelation, or the correct
	// positional return values from the server depending on the timing of the
	// call.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return a non-EOF error.
	// Finish should be called at most once.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.  It
	// is safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

// Implementation of the DrawInterfaceDrawStream interface that is not exported.
type implDrawInterfaceDrawStream struct {
	clientCall _gen_ipc.Call
}

func (c *implDrawInterfaceDrawStream) Send(item Box) error {
	return c.clientCall.Send(item)
}

func (c *implDrawInterfaceDrawStream) CloseSend() error {
	return c.clientCall.CloseSend()
}

func (c *implDrawInterfaceDrawStream) Recv() (item Box, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implDrawInterfaceDrawStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implDrawInterfaceDrawStream) Cancel() {
	c.clientCall.Cancel()
}

// DrawInterfaceServiceDrawStream is the interface for streaming responses of the method
// Draw in the service interface DrawInterface.
type DrawInterfaceServiceDrawStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.  If the client has canceled, an error is returned.
	Send(item Box) error

	// Recv fills itemptr with the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item Box, err error)
}

// Implementation of the DrawInterfaceServiceDrawStream interface that is not exported.
type implDrawInterfaceServiceDrawStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implDrawInterfaceServiceDrawStream) Send(item Box) error {
	return s.serverCall.Send(item)
}

func (s *implDrawInterfaceServiceDrawStream) Recv() (item Box, err error) {
	err = s.serverCall.Recv(&item)
	return
}

// BindDrawInterface returns the client stub implementing the DrawInterface
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindDrawInterface(name string, opts ..._gen_ipc.BindOpt) (DrawInterface, error) {
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
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubDrawInterface{client: client, name: name}

	return stub, nil
}

// NewServerDrawInterface creates a new server stub.
//
// It takes a regular server implementing the DrawInterfaceService
// interface, and returns a new server stub.
func NewServerDrawInterface(server DrawInterfaceService) interface{} {
	return &ServerStubDrawInterface{
		service: server,
	}
}

// clientStubDrawInterface implements DrawInterface.
type clientStubDrawInterface struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubDrawInterface) Draw(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply DrawInterfaceDrawStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Draw", nil, opts...); err != nil {
		return
	}
	reply = &implDrawInterfaceDrawStream{clientCall: call}
	return
}

func (__gen_c *clientStubDrawInterface) SyncBoxes(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "SyncBoxes", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubDrawInterface) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubDrawInterface) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubDrawInterface) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubDrawInterface wraps a server that implements
// DrawInterfaceService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubDrawInterface struct {
	service DrawInterfaceService
}

func (__gen_s *ServerStubDrawInterface) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Draw":
		return []interface{}{}, nil
	case "SyncBoxes":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubDrawInterface) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Draw"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "Err", Type: 65},
		},
		InStream:  67,
		OutStream: 67,
	}
	result.Methods["SyncBoxes"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "Err", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.ArrayType{Elem: 0x19, Len: 0x4, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "DeviceId"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "BoxId"},
				_gen_wiretype.FieldType{Type: 0x42, Name: "Points"},
			},
			"veyron/examples/boxes.Box", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubDrawInterface) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubDrawInterface) Draw(call _gen_ipc.ServerCall) (err error) {
	stream := &implDrawInterfaceServiceDrawStream{serverCall: call}
	err = __gen_s.service.Draw(call, stream)
	return
}

func (__gen_s *ServerStubDrawInterface) SyncBoxes(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.SyncBoxes(call)
	return
}
