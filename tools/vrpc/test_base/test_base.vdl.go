// This file was auto-generated by the veyron vdl tool.
// Source: test_base.vdl

package test_base

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdl "veyron2/vdl"
	_gen_wiretype "veyron2/wiretype"
)

type Struct struct {
	X int32
	Y int32
}

// TypeTester is the interface the client binds and uses.
// TypeTester_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type TypeTester_ExcludingUniversal interface {
	// Methods to test support for generic types.
	Bool(I1 bool, opts ..._gen_ipc.ClientCallOpt) (reply bool, err error)
	Float32(I1 float32, opts ..._gen_ipc.ClientCallOpt) (reply float32, err error)
	Float64(I1 float64, opts ..._gen_ipc.ClientCallOpt) (reply float64, err error)
	Int32(I1 int32, opts ..._gen_ipc.ClientCallOpt) (reply int32, err error)
	Int64(I1 int64, opts ..._gen_ipc.ClientCallOpt) (reply int64, err error)
	String(I1 string, opts ..._gen_ipc.ClientCallOpt) (reply string, err error)
	Byte(I1 byte, opts ..._gen_ipc.ClientCallOpt) (reply byte, err error)
	UInt32(I1 uint32, opts ..._gen_ipc.ClientCallOpt) (reply uint32, err error)
	UInt64(I1 uint64, opts ..._gen_ipc.ClientCallOpt) (reply uint64, err error)
	// Methods to test support for composite types.
	InputArray(I1 [2]byte, opts ..._gen_ipc.ClientCallOpt) (err error)
	InputMap(I1 map[byte]byte, opts ..._gen_ipc.ClientCallOpt) (err error)
	InputSlice(I1 []byte, opts ..._gen_ipc.ClientCallOpt) (err error)
	InputStruct(I1 Struct, opts ..._gen_ipc.ClientCallOpt) (err error)
	OutputArray(opts ..._gen_ipc.ClientCallOpt) (reply [2]byte, err error)
	OutputMap(opts ..._gen_ipc.ClientCallOpt) (reply map[byte]byte, err error)
	OutputSlice(opts ..._gen_ipc.ClientCallOpt) (reply []byte, err error)
	OutputStruct(opts ..._gen_ipc.ClientCallOpt) (reply Struct, err error)
	// Methods to test support for different number of arguments.
	NoArguments(opts ..._gen_ipc.ClientCallOpt) (err error)
	MultipleArguments(I1 int32, I2 int32, opts ..._gen_ipc.ClientCallOpt) (O1 int32, O2 int32, err error)
	// Methods to test support for streaming.
	StreamingOutput(NumStreamItems int32, StreamItem bool, opts ..._gen_ipc.ClientCallOpt) (reply TypeTesterStreamingOutputStream, err error)
}
type TypeTester interface {
	_gen_ipc.UniversalServiceMethods
	TypeTester_ExcludingUniversal
}

// TypeTesterService is the interface the server implements.
type TypeTesterService interface {

	// Methods to test support for generic types.
	Bool(context _gen_ipc.Context, I1 bool) (reply bool, err error)
	Float32(context _gen_ipc.Context, I1 float32) (reply float32, err error)
	Float64(context _gen_ipc.Context, I1 float64) (reply float64, err error)
	Int32(context _gen_ipc.Context, I1 int32) (reply int32, err error)
	Int64(context _gen_ipc.Context, I1 int64) (reply int64, err error)
	String(context _gen_ipc.Context, I1 string) (reply string, err error)
	Byte(context _gen_ipc.Context, I1 byte) (reply byte, err error)
	UInt32(context _gen_ipc.Context, I1 uint32) (reply uint32, err error)
	UInt64(context _gen_ipc.Context, I1 uint64) (reply uint64, err error)
	// Methods to test support for composite types.
	InputArray(context _gen_ipc.Context, I1 [2]byte) (err error)
	InputMap(context _gen_ipc.Context, I1 map[byte]byte) (err error)
	InputSlice(context _gen_ipc.Context, I1 []byte) (err error)
	InputStruct(context _gen_ipc.Context, I1 Struct) (err error)
	OutputArray(context _gen_ipc.Context) (reply [2]byte, err error)
	OutputMap(context _gen_ipc.Context) (reply map[byte]byte, err error)
	OutputSlice(context _gen_ipc.Context) (reply []byte, err error)
	OutputStruct(context _gen_ipc.Context) (reply Struct, err error)
	// Methods to test support for different number of arguments.
	NoArguments(context _gen_ipc.Context) (err error)
	MultipleArguments(context _gen_ipc.Context, I1 int32, I2 int32) (O1 int32, O2 int32, err error)
	// Methods to test support for streaming.
	StreamingOutput(context _gen_ipc.Context, NumStreamItems int32, StreamItem bool, stream TypeTesterServiceStreamingOutputStream) (err error)
}

// TypeTesterStreamingOutputStream is the interface for streaming responses of the method
// StreamingOutput in the service interface TypeTester.
type TypeTesterStreamingOutputStream interface {

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item bool, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the TypeTesterStreamingOutputStream interface that is not exported.
type implTypeTesterStreamingOutputStream struct {
	clientCall _gen_ipc.ClientCall
}

func (c *implTypeTesterStreamingOutputStream) Recv() (item bool, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implTypeTesterStreamingOutputStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implTypeTesterStreamingOutputStream) Cancel() {
	c.clientCall.Cancel()
}

// TypeTesterServiceStreamingOutputStream is the interface for streaming responses of the method
// StreamingOutput in the service interface TypeTester.
type TypeTesterServiceStreamingOutputStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item bool) error
}

// Implementation of the TypeTesterServiceStreamingOutputStream interface that is not exported.
type implTypeTesterServiceStreamingOutputStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implTypeTesterServiceStreamingOutputStream) Send(item bool) error {
	return s.serverCall.Send(item)
}

// BindTypeTester returns the client stub implementing the TypeTester
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindTypeTester(name string, opts ..._gen_ipc.BindOpt) (TypeTester, error) {
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
	stub := &clientStubTypeTester{client: client, name: name}

	return stub, nil
}

// NewServerTypeTester creates a new server stub.
//
// It takes a regular server implementing the TypeTesterService
// interface, and returns a new server stub.
func NewServerTypeTester(server TypeTesterService) interface{} {
	return &ServerStubTypeTester{
		service: server,
	}
}

// clientStubTypeTester implements TypeTester.
type clientStubTypeTester struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubTypeTester) Bool(I1 bool, opts ..._gen_ipc.ClientCallOpt) (reply bool, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Bool", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) Float32(I1 float32, opts ..._gen_ipc.ClientCallOpt) (reply float32, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Float32", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) Float64(I1 float64, opts ..._gen_ipc.ClientCallOpt) (reply float64, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Float64", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) Int32(I1 int32, opts ..._gen_ipc.ClientCallOpt) (reply int32, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Int32", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) Int64(I1 int64, opts ..._gen_ipc.ClientCallOpt) (reply int64, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Int64", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) String(I1 string, opts ..._gen_ipc.ClientCallOpt) (reply string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "String", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) Byte(I1 byte, opts ..._gen_ipc.ClientCallOpt) (reply byte, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Byte", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) UInt32(I1 uint32, opts ..._gen_ipc.ClientCallOpt) (reply uint32, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "UInt32", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) UInt64(I1 uint64, opts ..._gen_ipc.ClientCallOpt) (reply uint64, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "UInt64", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) InputArray(I1 [2]byte, opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "InputArray", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) InputMap(I1 map[byte]byte, opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "InputMap", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) InputSlice(I1 []byte, opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "InputSlice", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) InputStruct(I1 Struct, opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "InputStruct", []interface{}{I1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) OutputArray(opts ..._gen_ipc.ClientCallOpt) (reply [2]byte, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "OutputArray", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) OutputMap(opts ..._gen_ipc.ClientCallOpt) (reply map[byte]byte, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "OutputMap", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) OutputSlice(opts ..._gen_ipc.ClientCallOpt) (reply []byte, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "OutputSlice", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) OutputStruct(opts ..._gen_ipc.ClientCallOpt) (reply Struct, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "OutputStruct", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) NoArguments(opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "NoArguments", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) MultipleArguments(I1 int32, I2 int32, opts ..._gen_ipc.ClientCallOpt) (O1 int32, O2 int32, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "MultipleArguments", []interface{}{I1, I2}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&O1, &O2, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) StreamingOutput(NumStreamItems int32, StreamItem bool, opts ..._gen_ipc.ClientCallOpt) (reply TypeTesterStreamingOutputStream, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "StreamingOutput", []interface{}{NumStreamItems, StreamItem}, opts...); err != nil {
		return
	}
	reply = &implTypeTesterStreamingOutputStream{clientCall: call}
	return
}

func (__gen_c *clientStubTypeTester) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) Signature(opts ..._gen_ipc.ClientCallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTypeTester) GetMethodTags(method string, opts ..._gen_ipc.ClientCallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubTypeTester wraps a server that implements
// TypeTesterService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubTypeTester struct {
	service TypeTesterService
}

func (__gen_s *ServerStubTypeTester) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Bool":
		return []interface{}{}, nil
	case "Float32":
		return []interface{}{}, nil
	case "Float64":
		return []interface{}{}, nil
	case "Int32":
		return []interface{}{}, nil
	case "Int64":
		return []interface{}{}, nil
	case "String":
		return []interface{}{}, nil
	case "Byte":
		return []interface{}{}, nil
	case "UInt32":
		return []interface{}{}, nil
	case "UInt64":
		return []interface{}{}, nil
	case "InputArray":
		return []interface{}{}, nil
	case "InputMap":
		return []interface{}{}, nil
	case "InputSlice":
		return []interface{}{}, nil
	case "InputStruct":
		return []interface{}{}, nil
	case "OutputArray":
		return []interface{}{}, nil
	case "OutputMap":
		return []interface{}{}, nil
	case "OutputSlice":
		return []interface{}{}, nil
	case "OutputStruct":
		return []interface{}{}, nil
	case "NoArguments":
		return []interface{}{}, nil
	case "MultipleArguments":
		return []interface{}{}, nil
	case "StreamingOutput":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubTypeTester) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Bool"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 2},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 2},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["Byte"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 66},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 66},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["Float32"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 25},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 25},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["Float64"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 26},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 26},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["InputArray"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 67},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "E", Type: 65},
		},
	}
	result.Methods["InputMap"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 68},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "E", Type: 65},
		},
	}
	result.Methods["InputSlice"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 69},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "E", Type: 65},
		},
	}
	result.Methods["InputStruct"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 70},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "E", Type: 65},
		},
	}
	result.Methods["Int32"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 36},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 36},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["Int64"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 37},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 37},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["MultipleArguments"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 36},
			{Name: "I2", Type: 36},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 36},
			{Name: "O2", Type: 36},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["NoArguments"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["OutputArray"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 67},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["OutputMap"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 68},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["OutputSlice"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 69},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["OutputStruct"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 70},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["StreamingOutput"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "NumStreamItems", Type: 36},
			{Name: "StreamItem", Type: 2},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 2,
	}
	result.Methods["String"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 3},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["UInt32"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 52},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 52},
			{Name: "E", Type: 65},
		},
	}
	result.Methods["UInt64"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "I1", Type: 53},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "O1", Type: 53},
			{Name: "E", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.ArrayType{Elem: 0x42, Len: 0x2, Name: "", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x42, Elem: 0x42, Name: "", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x24, Name: "X"},
				_gen_wiretype.FieldType{Type: 0x24, Name: "Y"},
			},
			"veyron/tools/vrpc/test_base.Struct", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubTypeTester) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubTypeTester) Bool(call _gen_ipc.ServerCall, I1 bool) (reply bool, err error) {
	reply, err = __gen_s.service.Bool(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) Float32(call _gen_ipc.ServerCall, I1 float32) (reply float32, err error) {
	reply, err = __gen_s.service.Float32(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) Float64(call _gen_ipc.ServerCall, I1 float64) (reply float64, err error) {
	reply, err = __gen_s.service.Float64(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) Int32(call _gen_ipc.ServerCall, I1 int32) (reply int32, err error) {
	reply, err = __gen_s.service.Int32(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) Int64(call _gen_ipc.ServerCall, I1 int64) (reply int64, err error) {
	reply, err = __gen_s.service.Int64(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) String(call _gen_ipc.ServerCall, I1 string) (reply string, err error) {
	reply, err = __gen_s.service.String(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) Byte(call _gen_ipc.ServerCall, I1 byte) (reply byte, err error) {
	reply, err = __gen_s.service.Byte(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) UInt32(call _gen_ipc.ServerCall, I1 uint32) (reply uint32, err error) {
	reply, err = __gen_s.service.UInt32(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) UInt64(call _gen_ipc.ServerCall, I1 uint64) (reply uint64, err error) {
	reply, err = __gen_s.service.UInt64(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) InputArray(call _gen_ipc.ServerCall, I1 [2]byte) (err error) {
	err = __gen_s.service.InputArray(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) InputMap(call _gen_ipc.ServerCall, I1 map[byte]byte) (err error) {
	err = __gen_s.service.InputMap(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) InputSlice(call _gen_ipc.ServerCall, I1 []byte) (err error) {
	err = __gen_s.service.InputSlice(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) InputStruct(call _gen_ipc.ServerCall, I1 Struct) (err error) {
	err = __gen_s.service.InputStruct(call, I1)
	return
}

func (__gen_s *ServerStubTypeTester) OutputArray(call _gen_ipc.ServerCall) (reply [2]byte, err error) {
	reply, err = __gen_s.service.OutputArray(call)
	return
}

func (__gen_s *ServerStubTypeTester) OutputMap(call _gen_ipc.ServerCall) (reply map[byte]byte, err error) {
	reply, err = __gen_s.service.OutputMap(call)
	return
}

func (__gen_s *ServerStubTypeTester) OutputSlice(call _gen_ipc.ServerCall) (reply []byte, err error) {
	reply, err = __gen_s.service.OutputSlice(call)
	return
}

func (__gen_s *ServerStubTypeTester) OutputStruct(call _gen_ipc.ServerCall) (reply Struct, err error) {
	reply, err = __gen_s.service.OutputStruct(call)
	return
}

func (__gen_s *ServerStubTypeTester) NoArguments(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.NoArguments(call)
	return
}

func (__gen_s *ServerStubTypeTester) MultipleArguments(call _gen_ipc.ServerCall, I1 int32, I2 int32) (O1 int32, O2 int32, err error) {
	O1, O2, err = __gen_s.service.MultipleArguments(call, I1, I2)
	return
}

func (__gen_s *ServerStubTypeTester) StreamingOutput(call _gen_ipc.ServerCall, NumStreamItems int32, StreamItem bool) (err error) {
	stream := &implTypeTesterServiceStreamingOutputStream{serverCall: call}
	err = __gen_s.service.StreamingOutput(call, NumStreamItems, StreamItem, stream)
	return
}
