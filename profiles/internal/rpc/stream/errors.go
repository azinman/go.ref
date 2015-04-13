// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stream

import (
	"net"

	"v.io/v23/verror"
)

const pkgPath = "v.io/x/ref/profiles/internal/rpc/stream"

// The stream family of packages guarantee to return one of the verror codes defined
// here, their messages are constructed so as to avoid embedding a component/method name
// and are thus more suitable for inclusion in other verrors.
// This practiced of omitting {1}{2} is used throughout the stream packages since all
// of their errors are intended to be used as arguments to higher level errors.
var (
	ErrSecurity = verror.Register(pkgPath+".errSecurity", verror.NoRetry, "{:3}")
	ErrNetwork  = verror.Register(pkgPath+".errNetwork", verror.NoRetry, "{:3}")
	ErrProxy    = verror.Register(pkgPath+".errProxy", verror.NoRetry, "{:3}")
	ErrBadArg   = verror.Register(pkgPath+".errBadArg", verror.NoRetry, "{:3}")
	ErrBadState = verror.Register(pkgPath+".errBadState", verror.NoRetry, "{:3}")
	ErrAborted  = verror.Register(pkgPath+".errAborted", verror.NoRetry, "{:3}")
	// TODO(cnicolaou): remove this when the rest of the stream sub packages are converted.
	ErrSecOrNet = verror.Register(pkgPath+".errSecOrNet", verror.NoRetry, "{:3}")
)

// NetError implements net.Error
type NetError struct {
	err           error
	timeout, temp bool
}

// NewNetError returns a new net.Error which will return the
// supplied error, timeout and temporary parameters when the corresponding
// methods are invoked.
func NewNetError(err error, timeout, temporary bool) net.Error {
	return &NetError{err, timeout, temporary}
}

func (t NetError) Error() string   { return t.err.Error() }
func (t NetError) Timeout() bool   { return t.timeout }
func (t NetError) Temporary() bool { return t.temp }
