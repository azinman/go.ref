// This file was auto-generated by the veyron vdl tool.
// Source: identity.vdl

// Package identity defines services for identity providers in the veyron ecosystem.
package identity

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron.io/veyron/veyron2"
	_gen_context "veyron.io/veyron/veyron2/context"
	_gen_ipc "veyron.io/veyron/veyron2/ipc"
	_gen_naming "veyron.io/veyron/veyron2/naming"
	_gen_vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	_gen_wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// OAuthBlesser exchanges OAuth access tokens for
// an email address from an OAuth-based identity provider and uses the email
// address obtained to bless the client.
//
// OAuth is described in RFC 6749 (http://tools.ietf.org/html/rfc6749),
// though the Google implementation also has informative documentation at
// https://developers.google.com/accounts/docs/OAuth2
//
// WARNING: There is no binding between the channel over which the
// authorization code or access token was obtained (typically https)
// and the channel used to make the RPC (a veyron virtual circuit).
// Thus, if Mallory possesses the authorization code or access token
// associated with Alice's account, she may be able to obtain a blessing
// with Alice's name on it.
//
// TODO(ashankar,toddw): Once the "OneOf" type becomes available in VDL,
// then the "any" should be replaced by:
// OneOf<wire.ChainPublicID, []wire.ChainPublicID>
// where wire is from:
// import "veyron.io/veyron/veyron2/security/wire"
// OAuthBlesser is the interface the client binds and uses.
// OAuthBlesser_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type OAuthBlesser_ExcludingUniversal interface {
	// BlessUsingAccessToken uses the provided access token to obtain the email
	// address and returns a blessing.
	BlessUsingAccessToken(ctx _gen_context.T, token string, opts ..._gen_ipc.CallOpt) (reply _gen_vdlutil.Any, err error)
}
type OAuthBlesser interface {
	_gen_ipc.UniversalServiceMethods
	OAuthBlesser_ExcludingUniversal
}

// OAuthBlesserService is the interface the server implements.
type OAuthBlesserService interface {

	// BlessUsingAccessToken uses the provided access token to obtain the email
	// address and returns a blessing.
	BlessUsingAccessToken(context _gen_ipc.ServerContext, token string) (reply _gen_vdlutil.Any, err error)
}

// BindOAuthBlesser returns the client stub implementing the OAuthBlesser
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindOAuthBlesser(name string, opts ..._gen_ipc.BindOpt) (OAuthBlesser, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		// Do nothing.
	case 1:
		if clientOpt, ok := opts[0].(_gen_ipc.Client); opts[0] == nil || ok {
			client = clientOpt
		} else {
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubOAuthBlesser{defaultClient: client, name: name}

	return stub, nil
}

// NewServerOAuthBlesser creates a new server stub.
//
// It takes a regular server implementing the OAuthBlesserService
// interface, and returns a new server stub.
func NewServerOAuthBlesser(server OAuthBlesserService) interface{} {
	return &ServerStubOAuthBlesser{
		service: server,
	}
}

// clientStubOAuthBlesser implements OAuthBlesser.
type clientStubOAuthBlesser struct {
	defaultClient _gen_ipc.Client
	name          string
}

func (__gen_c *clientStubOAuthBlesser) client(ctx _gen_context.T) _gen_ipc.Client {
	if __gen_c.defaultClient != nil {
		return __gen_c.defaultClient
	}
	return _gen_veyron2.RuntimeFromContext(ctx).Client()
}

func (__gen_c *clientStubOAuthBlesser) BlessUsingAccessToken(ctx _gen_context.T, token string, opts ..._gen_ipc.CallOpt) (reply _gen_vdlutil.Any, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "BlessUsingAccessToken", []interface{}{token}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubOAuthBlesser) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubOAuthBlesser) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubOAuthBlesser) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubOAuthBlesser wraps a server that implements
// OAuthBlesserService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubOAuthBlesser struct {
	service OAuthBlesserService
}

func (__gen_s *ServerStubOAuthBlesser) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "BlessUsingAccessToken":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubOAuthBlesser) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["BlessUsingAccessToken"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "token", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "blessing", Type: 65},
			{Name: "err", Type: 66},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubOAuthBlesser) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubOAuthBlesser) BlessUsingAccessToken(call _gen_ipc.ServerCall, token string) (reply _gen_vdlutil.Any, err error) {
	reply, err = __gen_s.service.BlessUsingAccessToken(call, token)
	return
}

// MacaroonBlesser returns a blessing given the provided macaroon string.
// MacaroonBlesser is the interface the client binds and uses.
// MacaroonBlesser_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type MacaroonBlesser_ExcludingUniversal interface {
	// Bless uses the provided macaroon (which contains email and caveats)
	// to return a blessing for the client.
	Bless(ctx _gen_context.T, macaroon string, opts ..._gen_ipc.CallOpt) (reply _gen_vdlutil.Any, err error)
}
type MacaroonBlesser interface {
	_gen_ipc.UniversalServiceMethods
	MacaroonBlesser_ExcludingUniversal
}

// MacaroonBlesserService is the interface the server implements.
type MacaroonBlesserService interface {

	// Bless uses the provided macaroon (which contains email and caveats)
	// to return a blessing for the client.
	Bless(context _gen_ipc.ServerContext, macaroon string) (reply _gen_vdlutil.Any, err error)
}

// BindMacaroonBlesser returns the client stub implementing the MacaroonBlesser
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindMacaroonBlesser(name string, opts ..._gen_ipc.BindOpt) (MacaroonBlesser, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		// Do nothing.
	case 1:
		if clientOpt, ok := opts[0].(_gen_ipc.Client); opts[0] == nil || ok {
			client = clientOpt
		} else {
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubMacaroonBlesser{defaultClient: client, name: name}

	return stub, nil
}

// NewServerMacaroonBlesser creates a new server stub.
//
// It takes a regular server implementing the MacaroonBlesserService
// interface, and returns a new server stub.
func NewServerMacaroonBlesser(server MacaroonBlesserService) interface{} {
	return &ServerStubMacaroonBlesser{
		service: server,
	}
}

// clientStubMacaroonBlesser implements MacaroonBlesser.
type clientStubMacaroonBlesser struct {
	defaultClient _gen_ipc.Client
	name          string
}

func (__gen_c *clientStubMacaroonBlesser) client(ctx _gen_context.T) _gen_ipc.Client {
	if __gen_c.defaultClient != nil {
		return __gen_c.defaultClient
	}
	return _gen_veyron2.RuntimeFromContext(ctx).Client()
}

func (__gen_c *clientStubMacaroonBlesser) Bless(ctx _gen_context.T, macaroon string, opts ..._gen_ipc.CallOpt) (reply _gen_vdlutil.Any, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Bless", []interface{}{macaroon}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMacaroonBlesser) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMacaroonBlesser) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMacaroonBlesser) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubMacaroonBlesser wraps a server that implements
// MacaroonBlesserService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubMacaroonBlesser struct {
	service MacaroonBlesserService
}

func (__gen_s *ServerStubMacaroonBlesser) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Bless":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubMacaroonBlesser) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Bless"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "macaroon", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "blessing", Type: 65},
			{Name: "err", Type: 66},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubMacaroonBlesser) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubMacaroonBlesser) Bless(call _gen_ipc.ServerCall, macaroon string) (reply _gen_vdlutil.Any, err error) {
	reply, err = __gen_s.service.Bless(call, macaroon)
	return
}
