// An implementation of a server for WSPR

package server

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"v.io/x/ref/services/wsprd/lib"
	"v.io/x/ref/services/wsprd/principal"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/ipc"
	"v.io/v23/naming"
	"v.io/v23/security"
	"v.io/v23/vdl"
	"v.io/v23/vdlroot/signature"
	vdltime "v.io/v23/vdlroot/time"
	"v.io/v23/verror"
	"v.io/x/lib/vlog"
)

type Flow struct {
	ID     int32
	Writer lib.ClientWriter
}

// A request from the proxy to javascript to handle an RPC
type ServerRPCRequest struct {
	ServerId uint32
	Handle   int32
	Method   string
	Args     []interface{}
	Call     ServerRPCRequestCall
}

type ServerRPCRequestCall struct {
	SecurityCall SecurityCall
	Deadline     vdltime.Deadline
}

type FlowHandler interface {
	CreateNewFlow(server *Server, sender ipc.Stream) *Flow

	CleanupFlow(id int32)
}

type HandleStore interface {
	// Adds blessings to the store and returns handle to the blessings
	AddBlessings(blessings security.Blessings) int32
}

type ServerHelper interface {
	FlowHandler
	HandleStore

	Context() *context.T
}

type authReply struct {
	Err *verror.E
}

// AuthRequest is a request for a javascript authorizer to run
// This is exported to make the app test easier.
type AuthRequest struct {
	ServerID uint32       `json:"serverID"`
	Handle   int32        `json:"handle"`
	Call     SecurityCall `json:"call"`
}

type Server struct {
	// serverStateLock should be aquired when starting or stopping the server.
	// This should be locked before outstandingRequestLock.
	serverStateLock sync.Mutex

	// The ipc.ListenSpec to use with server.Listen
	listenSpec *ipc.ListenSpec

	// The server that handles the ipc layer.  Listen on this server is
	// lazily started.
	server ipc.Server

	// The saved dispatcher to reuse when serve is called multiple times.
	dispatcher *dispatcher

	// Whether the server is listening.
	isListening bool

	// The server id.
	id     uint32
	helper ServerHelper

	// outstandingRequestLock should be acquired only to update the
	// outstanding request maps below.
	outstandingRequestLock sync.Mutex

	// The set of outstanding server requests.
	outstandingServerRequests map[int32]chan *lib.ServerRPCReply

	outstandingAuthRequests map[int32]chan error

	outstandingValidationRequests map[int32]chan []error
}

func NewServer(id uint32, listenSpec *ipc.ListenSpec, helper ServerHelper) (*Server, error) {
	server := &Server{
		id:                            id,
		helper:                        helper,
		listenSpec:                    listenSpec,
		outstandingServerRequests:     make(map[int32]chan *lib.ServerRPCReply),
		outstandingAuthRequests:       make(map[int32]chan error),
		outstandingValidationRequests: make(map[int32]chan []error),
	}
	var err error
	ctx := helper.Context()
	ctx = context.WithValue(ctx, "customChainValidator", server.wsprCaveatValidator)
	if server.server, err = v23.NewServer(ctx); err != nil {
		return nil, err
	}
	return server, nil
}

// remoteInvokeFunc is a type of function that can invoke a remote method and
// communicate the result back via a channel to the caller
type remoteInvokeFunc func(methodName string, args []interface{}, call ipc.StreamServerCall) <-chan *lib.ServerRPCReply

func (s *Server) createRemoteInvokerFunc(handle int32) remoteInvokeFunc {
	return func(methodName string, args []interface{}, call ipc.StreamServerCall) <-chan *lib.ServerRPCReply {
		securityCall := s.convertSecurityCall(call, true)

		flow := s.helper.CreateNewFlow(s, call)
		replyChan := make(chan *lib.ServerRPCReply, 1)
		s.outstandingRequestLock.Lock()
		s.outstandingServerRequests[flow.ID] = replyChan
		s.outstandingRequestLock.Unlock()

		var timeout vdltime.Deadline
		if deadline, ok := call.Context().Deadline(); ok {
			timeout.Time = deadline
		}

		errHandler := func(err error) <-chan *lib.ServerRPCReply {
			if ch := s.popServerRequest(flow.ID); ch != nil {
				stdErr := verror.Convert(verror.ErrInternal, call.Context(), err).(verror.E)
				ch <- &lib.ServerRPCReply{nil, &stdErr}
				s.helper.CleanupFlow(flow.ID)
			}
			return replyChan

		}

		rpcCall := ServerRPCRequestCall{
			SecurityCall: securityCall,
			Deadline:     timeout,
		}

		// Send a invocation request to JavaScript
		message := ServerRPCRequest{
			ServerId: s.id,
			Handle:   handle,
			Method:   lib.LowercaseFirstCharacter(methodName),
			Args:     args,
			Call:     rpcCall,
		}
		vomMessage, err := lib.VomEncode(message)
		if err != nil {
			return errHandler(err)
		}
		if err := flow.Writer.Send(lib.ResponseServerRequest, vomMessage); err != nil {
			return errHandler(err)
		}

		vlog.VI(3).Infof("calling method %q with args %v, MessageID %d assigned\n", methodName, args, flow.ID)

		// Watch for cancellation.
		go func() {
			<-call.Context().Done()
			ch := s.popServerRequest(flow.ID)
			if ch == nil {
				return
			}

			// Send a cancel message to the JS server.
			flow.Writer.Send(lib.ResponseCancel, nil)
			s.helper.CleanupFlow(flow.ID)

			err := verror.Convert(verror.ErrAborted, call.Context(), call.Context().Err()).(verror.E)
			ch <- &lib.ServerRPCReply{nil, &err}
		}()

		go proxyStream(call, flow.Writer)

		return replyChan
	}
}

type globStream struct {
	ch  chan naming.VDLGlobReply
	ctx *context.T
}

func (g *globStream) Send(item interface{}) error {
	if v, ok := item.(naming.VDLGlobReply); ok {
		g.ch <- v
		return nil
	}
	return verror.New(verror.ErrBadArg, g.ctx, item)
}

func (g *globStream) Recv(itemptr interface{}) error {
	return verror.New(verror.ErrNoExist, g.ctx, "Can't call recieve on glob stream")
}

func (g *globStream) CloseSend() error {
	close(g.ch)
	return nil
}

// remoteGlobFunc is a type of function that can invoke a remote glob and
// communicate the result back via the channel returned
type remoteGlobFunc func(pattern string, call ipc.ServerCall) (<-chan naming.VDLGlobReply, error)

func (s *Server) createRemoteGlobFunc(handle int32) remoteGlobFunc {
	return func(pattern string, call ipc.ServerCall) (<-chan naming.VDLGlobReply, error) {
		// Until the tests get fixed, we need to create a security context before creating the flow
		// because creating the security context creates a flow and flow ids will be off.
		// See https://github.com/veyron/release-issues/issues/1181
		securityCall := s.convertSecurityCall(call, true)

		globChan := make(chan naming.VDLGlobReply, 1)
		flow := s.helper.CreateNewFlow(s, &globStream{
			ch:  globChan,
			ctx: call.Context(),
		})
		replyChan := make(chan *lib.ServerRPCReply, 1)
		s.outstandingRequestLock.Lock()
		s.outstandingServerRequests[flow.ID] = replyChan
		s.outstandingRequestLock.Unlock()

		var timeout vdltime.Deadline
		if deadline, ok := call.Context().Deadline(); ok {
			timeout.Time = deadline
		}

		errHandler := func(err error) (<-chan naming.VDLGlobReply, error) {
			if ch := s.popServerRequest(flow.ID); ch != nil {
				s.helper.CleanupFlow(flow.ID)
			}
			return nil, verror.Convert(verror.ErrInternal, call.Context(), err).(verror.E)
		}

		rpcCall := ServerRPCRequestCall{
			SecurityCall: securityCall,
			Deadline:     timeout,
		}

		// Send a invocation request to JavaScript
		message := ServerRPCRequest{
			ServerId: s.id,
			Handle:   handle,
			Method:   "Glob__",
			Args:     []interface{}{pattern},
			Call:     rpcCall,
		}
		vomMessage, err := lib.VomEncode(message)
		if err != nil {
			return errHandler(err)
		}
		if err := flow.Writer.Send(lib.ResponseServerRequest, vomMessage); err != nil {
			return errHandler(err)
		}

		vlog.VI(3).Infof("calling method 'Glob__' with args %v, MessageID %d assigned\n", []interface{}{pattern}, flow.ID)

		// Watch for cancellation.
		go func() {
			<-call.Context().Done()
			ch := s.popServerRequest(flow.ID)
			if ch == nil {
				return
			}

			// Send a cancel message to the JS server.
			flow.Writer.Send(lib.ResponseCancel, nil)
			s.helper.CleanupFlow(flow.ID)

			err := verror.Convert(verror.ErrAborted, call.Context(), call.Context().Err()).(verror.E)
			ch <- &lib.ServerRPCReply{nil, &err}
		}()

		return globChan, nil
	}
}

func proxyStream(stream ipc.Stream, w lib.ClientWriter) {
	var item interface{}
	for err := stream.Recv(&item); err == nil; err = stream.Recv(&item) {
		vomItem, err := lib.VomEncode(item)
		if err != nil {
			w.Error(verror.Convert(verror.ErrInternal, nil, err))
			return
		}
		if err := w.Send(lib.ResponseStream, vomItem); err != nil {
			w.Error(verror.Convert(verror.ErrInternal, nil, err))
			return
		}
	}
	if err := w.Send(lib.ResponseStreamClose, nil); err != nil {
		w.Error(verror.Convert(verror.ErrInternal, nil, err))
		return
	}
}

func (s *Server) convertBlessingsToHandle(blessings security.Blessings) principal.BlessingsHandle {
	return *principal.ConvertBlessingsToHandle(blessings, s.helper.AddBlessings(blessings))
}

func makeListOfErrors(numErrors int, err error) []error {
	errs := make([]error, numErrors)
	for i := 0; i < numErrors; i++ {
		errs[i] = err
	}
	return errs
}

// wsprCaveatValidator validates caveats in javascript.
// It resolves each []security.Caveat in cavs to an error (or nil) and collects them in a slice.
// TODO(ataly, ashankar, bprosnitz): Update this method so tha it also conveys the CallSide to
// JavaScript.
func (s *Server) wsprCaveatValidator(call security.Call, _ security.CallSide, cavs [][]security.Caveat) []error {
	flow := s.helper.CreateNewFlow(s, nil)
	req := CaveatValidationRequest{
		Call: s.convertSecurityCall(call, false),
		Cavs: cavs,
	}

	replyChan := make(chan []error, 1)
	s.outstandingRequestLock.Lock()
	s.outstandingValidationRequests[flow.ID] = replyChan
	s.outstandingRequestLock.Unlock()

	defer func() {
		s.outstandingRequestLock.Lock()
		delete(s.outstandingValidationRequests, flow.ID)
		s.outstandingRequestLock.Unlock()
		s.cleanupFlow(flow.ID)
	}()

	if err := flow.Writer.Send(lib.ResponseValidate, req); err != nil {
		vlog.VI(2).Infof("Failed to send validate response: %v", err)
		replyChan <- makeListOfErrors(len(cavs), err)
	}

	// TODO(bprosnitz) Consider using a different timeout than the standard ipc timeout.
	var timeoutChan <-chan time.Time
	if deadline, ok := call.Context().Deadline(); ok {
		timeoutChan = time.After(deadline.Sub(time.Now()))
	}

	select {
	case <-timeoutChan:
		return makeListOfErrors(len(cavs), NewErrCaveatValidationTimeout(call.Context()))
	case reply := <-replyChan:
		if len(reply) != len(cavs) {
			vlog.VI(2).Infof("Wspr caveat validator received %d results from javascript but expected %d", len(reply), len(cavs))
			return makeListOfErrors(len(cavs), NewErrInvalidValidationResponseFromJavascript(call.Context()))
		}

		return reply
	}
}

func (s *Server) convertSecurityCall(call security.Call, includeBlessingStrings bool) SecurityCall {
	// TODO(bprosnitz) Local/Remote Endpoint should always be non-nil, but isn't
	// due to a TODO in vc/auth.go
	var localEndpoint string
	if call.LocalEndpoint() != nil {
		localEndpoint = call.LocalEndpoint().String()
	}
	var remoteEndpoint string
	if call.RemoteEndpoint() != nil {
		remoteEndpoint = call.RemoteEndpoint().String()
	}
	var localBlessings principal.BlessingsHandle
	if !call.LocalBlessings().IsZero() {
		localBlessings = s.convertBlessingsToHandle(call.LocalBlessings())
	}
	anymtags := make([]*vdl.Value, len(call.MethodTags()))
	for i, mtag := range call.MethodTags() {
		anymtags[i] = mtag
	}
	secCtx := SecurityCall{
		Method:          lib.LowercaseFirstCharacter(call.Method()),
		Suffix:          call.Suffix(),
		MethodTags:      anymtags,
		LocalEndpoint:   localEndpoint,
		RemoteEndpoint:  remoteEndpoint,
		LocalBlessings:  localBlessings,
		RemoteBlessings: s.convertBlessingsToHandle(call.RemoteBlessings()),
	}
	if includeBlessingStrings {
		secCtx.LocalBlessingStrings, _ = call.LocalBlessings().ForCall(call)
		secCtx.RemoteBlessingStrings, _ = call.RemoteBlessings().ForCall(call)
	}
	return secCtx
}

type remoteAuthFunc func(call security.Call) error

func (s *Server) createRemoteAuthFunc(handle int32) remoteAuthFunc {
	return func(call security.Call) error {
		// Until the tests get fixed, we need to create a security context before creating the flow
		// because creating the security context creates a flow and flow ids will be off.
		securityCall := s.convertSecurityCall(call, true)

		flow := s.helper.CreateNewFlow(s, nil)
		replyChan := make(chan error, 1)
		s.outstandingRequestLock.Lock()
		s.outstandingAuthRequests[flow.ID] = replyChan
		s.outstandingRequestLock.Unlock()
		message := AuthRequest{
			ServerID: s.id,
			Handle:   handle,
			Call:     securityCall,
		}
		vlog.VI(0).Infof("Sending out auth request for %v, %v", flow.ID, message)

		vomMessage, err := lib.VomEncode(message)
		if err != nil {
			replyChan <- verror.Convert(verror.ErrInternal, nil, err)
		} else if err := flow.Writer.Send(lib.ResponseAuthRequest, vomMessage); err != nil {
			replyChan <- verror.Convert(verror.ErrInternal, nil, err)
		}

		err = <-replyChan
		vlog.VI(0).Infof("going to respond with %v", err)
		s.outstandingRequestLock.Lock()
		delete(s.outstandingAuthRequests, flow.ID)
		s.outstandingRequestLock.Unlock()
		s.helper.CleanupFlow(flow.ID)
		return err
	}
}

func (s *Server) Serve(name string) error {
	s.serverStateLock.Lock()
	defer s.serverStateLock.Unlock()

	if s.dispatcher == nil {
		s.dispatcher = newDispatcher(s.id, s, s, s)
	}

	if !s.isListening {
		_, err := s.server.Listen(*s.listenSpec)
		if err != nil {
			return err
		}
		s.isListening = true
	}
	if err := s.server.ServeDispatcher(name, s.dispatcher); err != nil {
		return err
	}
	return nil
}

func (s *Server) popServerRequest(id int32) chan *lib.ServerRPCReply {
	s.outstandingRequestLock.Lock()
	defer s.outstandingRequestLock.Unlock()
	ch := s.outstandingServerRequests[id]
	delete(s.outstandingServerRequests, id)

	return ch
}

func (s *Server) HandleServerResponse(id int32, data string) {
	ch := s.popServerRequest(id)
	if ch == nil {
		vlog.Errorf("unexpected result from JavaScript. No channel "+
			"for MessageId: %d exists. Ignoring the results.", id)
		// Ignore unknown responses that don't belong to any channel
		return
	}

	// Decode the result and send it through the channel
	var reply lib.ServerRPCReply
	if err := lib.VomDecode(data, &reply); err != nil {
		reply.Err = err
	}

	vlog.VI(0).Infof("response received from JavaScript server for "+
		"MessageId %d with result %v", id, reply)
	s.helper.CleanupFlow(id)
	ch <- &reply
}

func (s *Server) HandleLookupResponse(id int32, data string) {
	s.dispatcher.handleLookupResponse(id, data)
}

func (s *Server) HandleAuthResponse(id int32, data string) {
	s.outstandingRequestLock.Lock()
	ch := s.outstandingAuthRequests[id]
	s.outstandingRequestLock.Unlock()
	if ch == nil {
		vlog.Errorf("unexpected result from JavaScript. No channel "+
			"for MessageId: %d exists. Ignoring the results(%s)", id, data)
		//Ignore unknown responses that don't belong to any channel
		return
	}
	// Decode the result and send it through the channel
	var reply authReply
	if decoderErr := json.Unmarshal([]byte(data), &reply); decoderErr != nil {
		err := verror.Convert(verror.ErrInternal, nil, decoderErr).(verror.E)
		reply = authReply{Err: &err}
	}

	vlog.VI(0).Infof("response received from JavaScript server for "+
		"MessageId %d with result %v", id, reply)
	s.helper.CleanupFlow(id)
	// A nil verror.E does not result in an nil error.  Instead, we have create
	// a variable for the error interface and only set it's value if the struct is non-
	// nil.
	var err error
	if reply.Err != nil {
		err = reply.Err
	}
	ch <- err
}

func (s *Server) HandleCaveatValidationResponse(id int32, data string) {
	s.outstandingRequestLock.Lock()
	ch := s.outstandingValidationRequests[id]
	s.outstandingRequestLock.Unlock()
	if ch == nil {
		vlog.Errorf("unexpected result from JavaScript. No channel "+
			"for validation response with MessageId: %d exists. Ignoring the results(%s)", id, data)
		//Ignore unknown responses that don't belong to any channel
		return
	}

	var reply CaveatValidationResponse
	if err := lib.VomDecode(data, &reply); err != nil {
		vlog.Errorf("failed to decode validation response %q: error %v", data, err)
		ch <- []error{}
		return
	}

	ch <- reply.Results
}

func (s *Server) createFlow() *Flow {
	return s.helper.CreateNewFlow(s, nil)
}

func (s *Server) cleanupFlow(id int32) {
	s.helper.CleanupFlow(id)
}

func (s *Server) createInvoker(handle int32, sig []signature.Interface, hasGlobber bool) (ipc.Invoker, error) {
	remoteInvokeFunc := s.createRemoteInvokerFunc(handle)
	var globFunc remoteGlobFunc
	if hasGlobber {
		globFunc = s.createRemoteGlobFunc(handle)
	}
	return newInvoker(sig, remoteInvokeFunc, globFunc), nil
}

func (s *Server) createAuthorizer(handle int32, hasAuthorizer bool) (security.Authorizer, error) {
	if hasAuthorizer {
		return &authorizer{authFunc: s.createRemoteAuthFunc(handle)}, nil
	}
	return nil, nil
}

func (s *Server) Stop() {
	stdErr := verror.New(verror.ErrTimeout, nil).(verror.E)
	result := lib.ServerRPCReply{
		Results: nil,
		Err:     &stdErr,
	}
	s.serverStateLock.Lock()

	if s.dispatcher != nil {
		s.dispatcher.Cleanup()
	}

	for _, ch := range s.outstandingAuthRequests {
		ch <- fmt.Errorf("Cleaning up server")
	}
	s.outstandingAuthRequests = make(map[int32]chan error)

	for _, ch := range s.outstandingServerRequests {
		select {
		case ch <- &result:
		default:
		}
	}
	s.outstandingRequestLock.Lock()
	s.outstandingAuthRequests = make(map[int32]chan error)
	s.outstandingServerRequests = make(map[int32]chan *lib.ServerRPCReply)
	s.outstandingValidationRequests = make(map[int32]chan []error)
	s.outstandingRequestLock.Unlock()
	s.serverStateLock.Unlock()
	s.server.Stop()
}

func (s *Server) AddName(name string) error {
	return s.server.AddName(name)
}

func (s *Server) RemoveName(name string) {
	s.server.RemoveName(name)
}
