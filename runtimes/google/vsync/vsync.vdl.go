// This file was auto-generated by the veyron vdl tool.
// Source: vsync.vdl

package vsync

import (
	"veyron/services/store/raw"

	"veyron2/storage"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// DeviceID is the globally unique ID of a device.
type DeviceID string

// GenID is the unique ID per generation per device.
type GenID uint64

// LSN is the log sequence number.
type LSN uint64

// GenVector is the generation vector.
type GenVector map[DeviceID]GenID

// LogRec represents a single log record that is exchanged between two
// peers.
//
// It contains log related metadata: DevID is the id of the
// device that created the log record, GNum is the ID of the
// generation that the log record is part of, LSN is the log
// sequence number of the log record in the generation GNum,
// and RecType is the type of log record.
//
// It also contains information relevant to the updates to an object
// in the store: ObjID is the id of the object that was
// updated. CurVers is the current version number of the
// object. Parents can contain 0, 1 or 2 parent versions that the
// current version is derived from, and Value is the actual value of
// the object mutation.
type LogRec struct {
	// Log related information.
	DevID   DeviceID
	GNum    GenID
	LSN     LSN
	RecType byte
	// Object related information.
	ObjID   storage.ID
	CurVers storage.Version
	Parents []storage.Version
	Value   LogValue
}

// LogValue represents an object mutation within a transaction.
type LogValue struct {
	// Mutation is the store mutation representing the change in the object.
	Mutation raw.Mutation
	// SyncTime is the timestamp of the mutation when it arrives at the Sync server.
	SyncTime int64
	// Delete indicates whether the mutation resulted in the object being
	// deleted from the store.
	Delete bool
	// Continued tracks the transaction boundaries in a range of mutations.
	// It is set to true in all transaction mutations except the last one
	// in which it is set to false to mark the end of the transaction.
	Continued bool
}

const (
	// NodeRec type log record adds a new node in the dag.
	NodeRec = byte(0)

	// LinkRec type log record adds a new link in the dag.
	LinkRec = byte(1)
)

// Sync allows a device to GetDeltas from another device.
// Sync is the interface the client binds and uses.
// Sync_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Sync_ExcludingUniversal interface {
	// GetDeltas returns a device's current generation vector and all the missing log records
	// when compared to the incoming generation vector.
	GetDeltas(ctx _gen_context.T, In GenVector, ClientID DeviceID, opts ..._gen_ipc.CallOpt) (reply SyncGetDeltasStream, err error)
}
type Sync interface {
	_gen_ipc.UniversalServiceMethods
	Sync_ExcludingUniversal
}

// SyncService is the interface the server implements.
type SyncService interface {

	// GetDeltas returns a device's current generation vector and all the missing log records
	// when compared to the incoming generation vector.
	GetDeltas(context _gen_ipc.ServerContext, In GenVector, ClientID DeviceID, stream SyncServiceGetDeltasStream) (reply GenVector, err error)
}

// SyncGetDeltasStream is the interface for streaming responses of the method
// GetDeltas in the service interface Sync.
type SyncGetDeltasStream interface {

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of
	// input.
	Recv() (item LogRec, err error)

	// Finish blocks until the server is done and returns the positional
	// return values for call.
	//
	// If Cancel has been called, Finish will return immediately; the output of
	// Finish could either be an error signalling cancelation, or the correct
	// positional return values from the server depending on the timing of the
	// call.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return a non-EOF error.
	// Finish should be called at most once.
	Finish() (reply GenVector, err error)

	// Cancel cancels the RPC, notifying the server to stop processing.  It
	// is safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

// Implementation of the SyncGetDeltasStream interface that is not exported.
type implSyncGetDeltasStream struct {
	clientCall _gen_ipc.Call
}

func (c *implSyncGetDeltasStream) Recv() (item LogRec, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implSyncGetDeltasStream) Finish() (reply GenVector, err error) {
	if ierr := c.clientCall.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implSyncGetDeltasStream) Cancel() {
	c.clientCall.Cancel()
}

// SyncServiceGetDeltasStream is the interface for streaming responses of the method
// GetDeltas in the service interface Sync.
type SyncServiceGetDeltasStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.  If the client has canceled, an error is returned.
	Send(item LogRec) error
}

// Implementation of the SyncServiceGetDeltasStream interface that is not exported.
type implSyncServiceGetDeltasStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implSyncServiceGetDeltasStream) Send(item LogRec) error {
	return s.serverCall.Send(item)
}

// BindSync returns the client stub implementing the Sync
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindSync(name string, opts ..._gen_ipc.BindOpt) (Sync, error) {
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
	stub := &clientStubSync{client: client, name: name}

	return stub, nil
}

// NewServerSync creates a new server stub.
//
// It takes a regular server implementing the SyncService
// interface, and returns a new server stub.
func NewServerSync(server SyncService) interface{} {
	return &ServerStubSync{
		service: server,
	}
}

// clientStubSync implements Sync.
type clientStubSync struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubSync) GetDeltas(ctx _gen_context.T, In GenVector, ClientID DeviceID, opts ..._gen_ipc.CallOpt) (reply SyncGetDeltasStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetDeltas", []interface{}{In, ClientID}, opts...); err != nil {
		return
	}
	reply = &implSyncGetDeltasStream{clientCall: call}
	return
}

func (__gen_c *clientStubSync) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSync) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSync) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubSync wraps a server that implements
// SyncService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubSync struct {
	service SyncService
}

func (__gen_s *ServerStubSync) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "GetDeltas":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubSync) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["GetDeltas"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "In", Type: 67},
			{Name: "ClientID", Type: 65},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "Out", Type: 67},
			{Name: "Err", Type: 68},
		},

		OutStream: 82,
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x3, Name: "veyron/runtimes/google/vsync.DeviceID", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x35, Name: "veyron/runtimes/google/vsync.GenID", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x41, Elem: 0x42, Name: "veyron/runtimes/google/vsync.GenVector", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x35, Name: "veyron/runtimes/google/vsync.LSN", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.ArrayType{Elem: 0x46, Len: 0x10, Name: "veyron2/storage.ID", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x35, Name: "veyron2/storage.Version", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x48, Name: "", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "veyron2/storage.TagOp", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x4b, Name: "Op"},
				_gen_wiretype.FieldType{Type: 0x47, Name: "ACL"},
			},
			"veyron2/storage.Tag", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x4c, Name: "veyron2/storage.TagList", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x47, Name: "ID"},
			},
			"veyron2/storage.DEntry", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x4e, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x47, Name: "ID"},
				_gen_wiretype.FieldType{Type: 0x48, Name: "PriorVersion"},
				_gen_wiretype.FieldType{Type: 0x48, Name: "Version"},
				_gen_wiretype.FieldType{Type: 0x2, Name: "IsRoot"},
				_gen_wiretype.FieldType{Type: 0x4a, Name: "Value"},
				_gen_wiretype.FieldType{Type: 0x4d, Name: "Tags"},
				_gen_wiretype.FieldType{Type: 0x4f, Name: "Dir"},
			},
			"veyron/services/store/raw.Mutation", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x50, Name: "Mutation"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "SyncTime"},
				_gen_wiretype.FieldType{Type: 0x2, Name: "Delete"},
				_gen_wiretype.FieldType{Type: 0x2, Name: "Continued"},
			},
			"veyron/runtimes/google/vsync.LogValue", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x41, Name: "DevID"},
				_gen_wiretype.FieldType{Type: 0x42, Name: "GNum"},
				_gen_wiretype.FieldType{Type: 0x45, Name: "LSN"},
				_gen_wiretype.FieldType{Type: 0x46, Name: "RecType"},
				_gen_wiretype.FieldType{Type: 0x47, Name: "ObjID"},
				_gen_wiretype.FieldType{Type: 0x48, Name: "CurVers"},
				_gen_wiretype.FieldType{Type: 0x49, Name: "Parents"},
				_gen_wiretype.FieldType{Type: 0x51, Name: "Value"},
			},
			"veyron/runtimes/google/vsync.LogRec", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubSync) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubSync) GetDeltas(call _gen_ipc.ServerCall, In GenVector, ClientID DeviceID) (reply GenVector, err error) {
	stream := &implSyncServiceGetDeltasStream{serverCall: call}
	reply, err = __gen_s.service.GetDeltas(call, In, ClientID, stream)
	return
}
