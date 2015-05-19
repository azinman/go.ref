// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package message

import (
	"fmt"

	"v.io/x/ref/runtime/internal/lib/iobuf"
	"v.io/x/ref/runtime/internal/rpc/stream/id"
)

// Data encapsulates an application data message.
type Data struct {
	VCI     id.VC // Must be non-zero.
	Flow    id.Flow
	flags   uint8
	Payload *iobuf.Slice
}

// Close returns true if the sender of the data message requested that the flow be closed.
func (d *Data) Close() bool { return d.flags&0x1 == 1 }

// SetClose sets the Close flag of the message.
func (d *Data) SetClose() { d.flags |= 0x1 }

// Release releases the Payload
func (d *Data) Release() {
	if d.Payload != nil {
		d.Payload.Release()
		d.Payload = nil
	}
}

func (d *Data) PayloadSize() int {
	if d.Payload == nil {
		return 0
	}
	return d.Payload.Size()
}

func (d *Data) String() string {
	return fmt.Sprintf("VCI:%d Flow:%d Flags:%02x Payload:(%d bytes)", d.VCI, d.Flow, d.flags, d.PayloadSize())
}