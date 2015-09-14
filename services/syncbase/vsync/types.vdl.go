// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package vsync

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/x/ref/services/syncbase/server/interfaces"
)

// syncData represents the persistent state of the sync module.
type syncData struct {
	Id uint64
}

func (syncData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.syncData"`
}) {
}

// localGenInfo represents the persistent state corresponding to local generations.
type localGenInfo struct {
	Gen        uint64 // local generation number incremented on every local update.
	CheckptGen uint64 // local generation number advertised to remote peers (used by the responder).
}

func (localGenInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.localGenInfo"`
}) {
}

// dbSyncState represents the persistent sync state of a Database.
type dbSyncState struct {
	Data     localGenInfo
	Sgs      map[interfaces.GroupId]localGenInfo
	GenVec   interfaces.GenVector // generation vector capturing the locally-known generations of remote peers for data in Database.
	SgGenVec interfaces.GenVector // generation vector capturing the locally-known generations of remote peers for SyncGroups in Database.
}

func (dbSyncState) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.dbSyncState"`
}) {
}

// localLogRec represents the persistent local state of a log record. Metadata
// is synced across peers, while pos is local-only.
type localLogRec struct {
	Metadata interfaces.LogRecMetadata
	Pos      uint64 // position in the Database log.
}

func (localLogRec) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.localLogRec"`
}) {
}

// sgLocalState holds the SyncGroup local state, only relevant to this member
// (i.e. the local Syncbase).  This is needed for crash recovery of the internal
// state transitions of the SyncGroup.
type sgLocalState struct {
	// The count of local joiners to the same SyncGroup.
	NumLocalJoiners uint32
	// The SyncGroup is watched when the sync Watcher starts processing the
	// SyncGroup data.  When a SyncGroup is created or joined, an entry is
	// added to the Watcher queue (log) to inform it from which point to
	// start accepting store mutations, an asynchronous notification similar
	// to regular store mutations.  When the Watcher processes that queue
	// entry, it sets this bit to true.  When Syncbase restarts, the value
	// of this bit allows the new sync Watcher to recreate its in-memory
	// state by resuming to watch only the prefixes of SyncGroups that were
	// previously being watched.
	Watched bool
	// The SyncGroup was published here by this remote peer (if non-empty
	// string), typically the SyncGroup creator.  In this case the SyncGroup
	// cannot be GCed locally even if it has no local joiners.
	RemotePublisher string
	// The SyncGroup is in pending state on a device that learns the current
	// state of the SyncGroup from another device but has not yet received
	// through peer-to-peer sync the history of the changes (DAG and logs).
	// This happens in two cases:
	// 1- A joiner was accepted into a SyncGroup by a SyncGroup admin and
	//    only given the current SyncGroup info synchronously and will
	//    receive the full history later via p2p sync.
	// 2- A remote server where the SyncGroup is published was told by the
	//    SyncGroup publisher the current SyncGroup info synchronously and
	//    will receive the full history later via p2p sync.
	// The pending state is over when the device reaches or exceeds the
	// knowledge level indicated in the pending genvec.  While SyncPending
	// is true, no local SyncGroup mutations are allowed (i.e. no join or
	// set-spec requests).
	SyncPending   bool
	PendingGenVec interfaces.PrefixGenVector
}

func (sgLocalState) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.sgLocalState"`
}) {
}

func init() {
	vdl.Register((*syncData)(nil))
	vdl.Register((*localGenInfo)(nil))
	vdl.Register((*dbSyncState)(nil))
	vdl.Register((*localLogRec)(nil))
	vdl.Register((*sgLocalState)(nil))
}

const logPrefix = "log" // log state.

const logDataPrefix = "data" // data log state.

const dbssPrefix = "dbss" // database sync state.

const dagPrefix = "dag" // dag state.

const sgPrefix = "sg" // syncgroup state.
