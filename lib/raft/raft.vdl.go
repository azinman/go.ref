// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: raft.vdl

package raft

import (
	"fmt"
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"
	"v.io/v23/vdl/vdlconv"
)

// Term is a counter incremented each time a member starts an election.  The log will
// show gaps in Term numbers because all elections need not be successful.
type Term uint64

func (Term) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/raft.Term"`
}) {
}

func (m *Term) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromUint(uint64((*m)), __VDLType_raft_v_io_x_ref_lib_raft_Term); err != nil {
		return err
	}
	return nil
}

func (m *Term) MakeVDLTarget() vdl.Target {
	return &TermTarget{Value: m}
}

type TermTarget struct {
	Value *Term
	vdl.TargetBase
}

func (t *TermTarget) FromUint(src uint64, tt *vdl.Type) error {
	*t.Value = Term(src)
	return nil
}
func (t *TermTarget) FromInt(src int64, tt *vdl.Type) error {
	val, err := vdlconv.Int64ToUint64(src)
	if err != nil {
		return err
	}
	*t.Value = Term(val)
	return nil
}
func (t *TermTarget) FromFloat(src float64, tt *vdl.Type) error {
	val, err := vdlconv.Float64ToUint64(src)
	if err != nil {
		return err
	}
	*t.Value = Term(val)
	return nil
}
func (t *TermTarget) FromComplex(src complex128, tt *vdl.Type) error {
	val, err := vdlconv.Complex128ToUint64(src)
	if err != nil {
		return err
	}
	*t.Value = Term(val)
	return nil
}

// Index is an index into the log.  The log entries are numbered sequentially.  At the moment
// the entries RaftClient.Apply()ed should be sequential but that will change if we introduce
// system entries. For example, we could have an entry type that is used to add members to the
// set of replicas.
type Index uint64

func (Index) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/raft.Index"`
}) {
}

func (m *Index) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromUint(uint64((*m)), __VDLType_raft_v_io_x_ref_lib_raft_Index); err != nil {
		return err
	}
	return nil
}

func (m *Index) MakeVDLTarget() vdl.Target {
	return &IndexTarget{Value: m}
}

type IndexTarget struct {
	Value *Index
	vdl.TargetBase
}

func (t *IndexTarget) FromUint(src uint64, tt *vdl.Type) error {
	*t.Value = Index(src)
	return nil
}
func (t *IndexTarget) FromInt(src int64, tt *vdl.Type) error {
	val, err := vdlconv.Int64ToUint64(src)
	if err != nil {
		return err
	}
	*t.Value = Index(val)
	return nil
}
func (t *IndexTarget) FromFloat(src float64, tt *vdl.Type) error {
	val, err := vdlconv.Float64ToUint64(src)
	if err != nil {
		return err
	}
	*t.Value = Index(val)
	return nil
}
func (t *IndexTarget) FromComplex(src complex128, tt *vdl.Type) error {
	val, err := vdlconv.Complex128ToUint64(src)
	if err != nil {
		return err
	}
	*t.Value = Index(val)
	return nil
}

// The LogEntry is what the log consists of.  'error' starts nil and is never written to stable
// storage.  It represents the result of RaftClient.Apply(Cmd, Index).  This is a hack but I
// haven't figured out a better way.
type LogEntry struct {
	Term  Term
	Index Index
	Cmd   []byte
	Type  byte
}

func (LogEntry) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/raft.LogEntry"`
}) {
}

func (m *LogEntry) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_raft_v_io_x_ref_lib_raft_LogEntry == nil || __VDLTyperaft0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Term")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Term.FillVDLTarget(fieldTarget3, __VDLType_raft_v_io_x_ref_lib_raft_Term); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Index")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Index.FillVDLTarget(fieldTarget5, __VDLType_raft_v_io_x_ref_lib_raft_Index); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Cmd")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget7.FromBytes([]byte(m.Cmd), __VDLTyperaft1); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Type")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget9.FromUint(uint64(m.Type), vdl.ByteType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *LogEntry) MakeVDLTarget() vdl.Target {
	return &LogEntryTarget{Value: m}
}

type LogEntryTarget struct {
	Value *LogEntry
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *LogEntryTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_raft_v_io_x_ref_lib_raft_LogEntry) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_raft_v_io_x_ref_lib_raft_LogEntry)
	}
	return t, nil
}
func (t *LogEntryTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Term":
		val, err := &TermTarget{Value: &t.Value.Term}, error(nil)
		return nil, val, err
	case "Index":
		val, err := &IndexTarget{Value: &t.Value.Index}, error(nil)
		return nil, val, err
	case "Cmd":
		val, err := &vdl.BytesTarget{Value: &t.Value.Cmd}, error(nil)
		return nil, val, err
	case "Type":
		val, err := &vdl.ByteTarget{Value: &t.Value.Type}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_raft_v_io_x_ref_lib_raft_LogEntry)
	}
}
func (t *LogEntryTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *LogEntryTarget) FinishFields(_ vdl.FieldsTarget) error {
	return nil
}

func init() {
	vdl.Register((*Term)(nil))
	vdl.Register((*Index)(nil))
	vdl.Register((*LogEntry)(nil))
}

var __VDLTyperaft0 *vdl.Type = vdl.TypeOf((*LogEntry)(nil))
var __VDLTyperaft1 *vdl.Type = vdl.TypeOf([]byte(nil))
var __VDLType_raft_v_io_x_ref_lib_raft_Index *vdl.Type = vdl.TypeOf(Index(0))
var __VDLType_raft_v_io_x_ref_lib_raft_LogEntry *vdl.Type = vdl.TypeOf(LogEntry{})
var __VDLType_raft_v_io_x_ref_lib_raft_Term *vdl.Type = vdl.TypeOf(Term(0))

func __VDLEnsureNativeBuilt_raft() {
}

const ClientEntry = byte(0)

const RaftEntry = byte(1)

// raftProtoClientMethods is the client interface
// containing raftProto methods.
//
// raftProto is used by the members of a raft set to communicate with each other.
type raftProtoClientMethods interface {
	// Members returns the current set of ids of raft members.
	Members(*context.T, ...rpc.CallOpt) ([]string, error)
	// Leader returns the id of the current leader.
	Leader(*context.T, ...rpc.CallOpt) (string, error)
	// RequestVote starts a new round of voting.  It returns the server's current Term and true if
	// the server voted for the client.
	RequestVote(_ *context.T, term Term, candidateId string, lastLogTerm Term, lastLogIndex Index, _ ...rpc.CallOpt) (Term Term, Granted bool, _ error)
	// AppendToLog is sent by the leader to tell followers to append an entry.  If cmds
	// is empty, this is a keep alive message (at a random interval after a keep alive, followers
	// will initiate a new round of voting).
	//   term -- the current term of the sender
	//   leaderId -- the id of the sender
	//   prevIndex -- the index of the log entry immediately preceding cmds
	//   prevTerm -- the term of the log entry immediately preceding cmds.  The receiver must have
	//               received the previous index'd entry and it must have had the same term.  Otherwise
	//               an error is returned.
	//   leaderCommit -- the index of the last committed entry, i.e., the one a quorum has gauranteed
	//                   to have logged.
	//   cmds -- sequential log entries starting at prevIndex+1
	AppendToLog(_ *context.T, term Term, leaderId string, prevIndex Index, prevTerm Term, leaderCommit Index, cmds []LogEntry, _ ...rpc.CallOpt) error
	// Append is sent to the leader by followers.  Only the leader is allowed to send AppendToLog.
	// If a follower receives an Append() call it performs an Append() to the leader to run the actual
	// Raft algorithm.  The leader will respond after it has RaftClient.Apply()ed the command.
	//
	// Returns the term and index of the append entry or an error.
	Append(_ *context.T, cmd []byte, _ ...rpc.CallOpt) (term Term, index Index, _ error)
	// Committed returns the commit index of the leader.
	Committed(*context.T, ...rpc.CallOpt) (index Index, _ error)
	// InstallSnapshot is sent from the leader to follower to install the given snapshot.  It is
	// sent when it becomes apparent that the leader does not have log entries needed by the follower
	// to progress.  'term' and 'index' represent the last LogEntry RaftClient.Apply()ed to the
	// snapshot.
	InstallSnapshot(_ *context.T, term Term, leaderId string, appliedTerm Term, appliedIndex Index, _ ...rpc.CallOpt) (raftProtoInstallSnapshotClientCall, error)
}

// raftProtoClientStub adds universal methods to raftProtoClientMethods.
type raftProtoClientStub interface {
	raftProtoClientMethods
	rpc.UniversalServiceMethods
}

// raftProtoClient returns a client stub for raftProto.
func raftProtoClient(name string) raftProtoClientStub {
	return implraftProtoClientStub{name}
}

type implraftProtoClientStub struct {
	name string
}

func (c implraftProtoClientStub) Members(ctx *context.T, opts ...rpc.CallOpt) (o0 []string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Members", nil, []interface{}{&o0}, opts...)
	return
}

func (c implraftProtoClientStub) Leader(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Leader", nil, []interface{}{&o0}, opts...)
	return
}

func (c implraftProtoClientStub) RequestVote(ctx *context.T, i0 Term, i1 string, i2 Term, i3 Index, opts ...rpc.CallOpt) (o0 Term, o1 bool, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "RequestVote", []interface{}{i0, i1, i2, i3}, []interface{}{&o0, &o1}, opts...)
	return
}

func (c implraftProtoClientStub) AppendToLog(ctx *context.T, i0 Term, i1 string, i2 Index, i3 Term, i4 Index, i5 []LogEntry, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "AppendToLog", []interface{}{i0, i1, i2, i3, i4, i5}, nil, opts...)
	return
}

func (c implraftProtoClientStub) Append(ctx *context.T, i0 []byte, opts ...rpc.CallOpt) (o0 Term, o1 Index, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Append", []interface{}{i0}, []interface{}{&o0, &o1}, opts...)
	return
}

func (c implraftProtoClientStub) Committed(ctx *context.T, opts ...rpc.CallOpt) (o0 Index, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Committed", nil, []interface{}{&o0}, opts...)
	return
}

func (c implraftProtoClientStub) InstallSnapshot(ctx *context.T, i0 Term, i1 string, i2 Term, i3 Index, opts ...rpc.CallOpt) (ocall raftProtoInstallSnapshotClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "InstallSnapshot", []interface{}{i0, i1, i2, i3}, opts...); err != nil {
		return
	}
	ocall = &implraftProtoInstallSnapshotClientCall{ClientCall: call}
	return
}

// raftProtoInstallSnapshotClientStream is the client stream for raftProto.InstallSnapshot.
type raftProtoInstallSnapshotClientStream interface {
	// SendStream returns the send side of the raftProto.InstallSnapshot client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item []byte) error
		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.
		// This is an optional call - e.g. a client might call Close if it
		// needs to continue receiving items from the server after it's
		// done sending.  Returns errors encountered while closing, or if
		// Close is called after the stream has been canceled.  Like Send,
		// blocks if there is no buffer space available.
		Close() error
	}
}

// raftProtoInstallSnapshotClientCall represents the call returned from raftProto.InstallSnapshot.
type raftProtoInstallSnapshotClientCall interface {
	raftProtoInstallSnapshotClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implraftProtoInstallSnapshotClientCall struct {
	rpc.ClientCall
}

func (c *implraftProtoInstallSnapshotClientCall) SendStream() interface {
	Send(item []byte) error
	Close() error
} {
	return implraftProtoInstallSnapshotClientCallSend{c}
}

type implraftProtoInstallSnapshotClientCallSend struct {
	c *implraftProtoInstallSnapshotClientCall
}

func (c implraftProtoInstallSnapshotClientCallSend) Send(item []byte) error {
	return c.c.Send(item)
}
func (c implraftProtoInstallSnapshotClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implraftProtoInstallSnapshotClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// raftProtoServerMethods is the interface a server writer
// implements for raftProto.
//
// raftProto is used by the members of a raft set to communicate with each other.
type raftProtoServerMethods interface {
	// Members returns the current set of ids of raft members.
	Members(*context.T, rpc.ServerCall) ([]string, error)
	// Leader returns the id of the current leader.
	Leader(*context.T, rpc.ServerCall) (string, error)
	// RequestVote starts a new round of voting.  It returns the server's current Term and true if
	// the server voted for the client.
	RequestVote(_ *context.T, _ rpc.ServerCall, term Term, candidateId string, lastLogTerm Term, lastLogIndex Index) (Term Term, Granted bool, _ error)
	// AppendToLog is sent by the leader to tell followers to append an entry.  If cmds
	// is empty, this is a keep alive message (at a random interval after a keep alive, followers
	// will initiate a new round of voting).
	//   term -- the current term of the sender
	//   leaderId -- the id of the sender
	//   prevIndex -- the index of the log entry immediately preceding cmds
	//   prevTerm -- the term of the log entry immediately preceding cmds.  The receiver must have
	//               received the previous index'd entry and it must have had the same term.  Otherwise
	//               an error is returned.
	//   leaderCommit -- the index of the last committed entry, i.e., the one a quorum has gauranteed
	//                   to have logged.
	//   cmds -- sequential log entries starting at prevIndex+1
	AppendToLog(_ *context.T, _ rpc.ServerCall, term Term, leaderId string, prevIndex Index, prevTerm Term, leaderCommit Index, cmds []LogEntry) error
	// Append is sent to the leader by followers.  Only the leader is allowed to send AppendToLog.
	// If a follower receives an Append() call it performs an Append() to the leader to run the actual
	// Raft algorithm.  The leader will respond after it has RaftClient.Apply()ed the command.
	//
	// Returns the term and index of the append entry or an error.
	Append(_ *context.T, _ rpc.ServerCall, cmd []byte) (term Term, index Index, _ error)
	// Committed returns the commit index of the leader.
	Committed(*context.T, rpc.ServerCall) (index Index, _ error)
	// InstallSnapshot is sent from the leader to follower to install the given snapshot.  It is
	// sent when it becomes apparent that the leader does not have log entries needed by the follower
	// to progress.  'term' and 'index' represent the last LogEntry RaftClient.Apply()ed to the
	// snapshot.
	InstallSnapshot(_ *context.T, _ raftProtoInstallSnapshotServerCall, term Term, leaderId string, appliedTerm Term, appliedIndex Index) error
}

// raftProtoServerStubMethods is the server interface containing
// raftProto methods, as expected by rpc.Server.
// The only difference between this interface and raftProtoServerMethods
// is the streaming methods.
type raftProtoServerStubMethods interface {
	// Members returns the current set of ids of raft members.
	Members(*context.T, rpc.ServerCall) ([]string, error)
	// Leader returns the id of the current leader.
	Leader(*context.T, rpc.ServerCall) (string, error)
	// RequestVote starts a new round of voting.  It returns the server's current Term and true if
	// the server voted for the client.
	RequestVote(_ *context.T, _ rpc.ServerCall, term Term, candidateId string, lastLogTerm Term, lastLogIndex Index) (Term Term, Granted bool, _ error)
	// AppendToLog is sent by the leader to tell followers to append an entry.  If cmds
	// is empty, this is a keep alive message (at a random interval after a keep alive, followers
	// will initiate a new round of voting).
	//   term -- the current term of the sender
	//   leaderId -- the id of the sender
	//   prevIndex -- the index of the log entry immediately preceding cmds
	//   prevTerm -- the term of the log entry immediately preceding cmds.  The receiver must have
	//               received the previous index'd entry and it must have had the same term.  Otherwise
	//               an error is returned.
	//   leaderCommit -- the index of the last committed entry, i.e., the one a quorum has gauranteed
	//                   to have logged.
	//   cmds -- sequential log entries starting at prevIndex+1
	AppendToLog(_ *context.T, _ rpc.ServerCall, term Term, leaderId string, prevIndex Index, prevTerm Term, leaderCommit Index, cmds []LogEntry) error
	// Append is sent to the leader by followers.  Only the leader is allowed to send AppendToLog.
	// If a follower receives an Append() call it performs an Append() to the leader to run the actual
	// Raft algorithm.  The leader will respond after it has RaftClient.Apply()ed the command.
	//
	// Returns the term and index of the append entry or an error.
	Append(_ *context.T, _ rpc.ServerCall, cmd []byte) (term Term, index Index, _ error)
	// Committed returns the commit index of the leader.
	Committed(*context.T, rpc.ServerCall) (index Index, _ error)
	// InstallSnapshot is sent from the leader to follower to install the given snapshot.  It is
	// sent when it becomes apparent that the leader does not have log entries needed by the follower
	// to progress.  'term' and 'index' represent the last LogEntry RaftClient.Apply()ed to the
	// snapshot.
	InstallSnapshot(_ *context.T, _ *raftProtoInstallSnapshotServerCallStub, term Term, leaderId string, appliedTerm Term, appliedIndex Index) error
}

// raftProtoServerStub adds universal methods to raftProtoServerStubMethods.
type raftProtoServerStub interface {
	raftProtoServerStubMethods
	// Describe the raftProto interfaces.
	Describe__() []rpc.InterfaceDesc
}

// raftProtoServer returns a server stub for raftProto.
// It converts an implementation of raftProtoServerMethods into
// an object that may be used by rpc.Server.
func raftProtoServer(impl raftProtoServerMethods) raftProtoServerStub {
	stub := implraftProtoServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implraftProtoServerStub struct {
	impl raftProtoServerMethods
	gs   *rpc.GlobState
}

func (s implraftProtoServerStub) Members(ctx *context.T, call rpc.ServerCall) ([]string, error) {
	return s.impl.Members(ctx, call)
}

func (s implraftProtoServerStub) Leader(ctx *context.T, call rpc.ServerCall) (string, error) {
	return s.impl.Leader(ctx, call)
}

func (s implraftProtoServerStub) RequestVote(ctx *context.T, call rpc.ServerCall, i0 Term, i1 string, i2 Term, i3 Index) (Term, bool, error) {
	return s.impl.RequestVote(ctx, call, i0, i1, i2, i3)
}

func (s implraftProtoServerStub) AppendToLog(ctx *context.T, call rpc.ServerCall, i0 Term, i1 string, i2 Index, i3 Term, i4 Index, i5 []LogEntry) error {
	return s.impl.AppendToLog(ctx, call, i0, i1, i2, i3, i4, i5)
}

func (s implraftProtoServerStub) Append(ctx *context.T, call rpc.ServerCall, i0 []byte) (Term, Index, error) {
	return s.impl.Append(ctx, call, i0)
}

func (s implraftProtoServerStub) Committed(ctx *context.T, call rpc.ServerCall) (Index, error) {
	return s.impl.Committed(ctx, call)
}

func (s implraftProtoServerStub) InstallSnapshot(ctx *context.T, call *raftProtoInstallSnapshotServerCallStub, i0 Term, i1 string, i2 Term, i3 Index) error {
	return s.impl.InstallSnapshot(ctx, call, i0, i1, i2, i3)
}

func (s implraftProtoServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implraftProtoServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{raftProtoDesc}
}

// raftProtoDesc describes the raftProto interface.
var raftProtoDesc rpc.InterfaceDesc = descraftProto

// descraftProto hides the desc to keep godoc clean.
var descraftProto = rpc.InterfaceDesc{
	Name:    "raftProto",
	PkgPath: "v.io/x/ref/lib/raft",
	Doc:     "// raftProto is used by the members of a raft set to communicate with each other.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Members",
			Doc:  "// Members returns the current set of ids of raft members.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []string
			},
		},
		{
			Name: "Leader",
			Doc:  "// Leader returns the id of the current leader.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
		},
		{
			Name: "RequestVote",
			Doc:  "// RequestVote starts a new round of voting.  It returns the server's current Term and true if\n// the server voted for the client.",
			InArgs: []rpc.ArgDesc{
				{"term", ``},         // Term
				{"candidateId", ``},  // string
				{"lastLogTerm", ``},  // Term
				{"lastLogIndex", ``}, // Index
			},
			OutArgs: []rpc.ArgDesc{
				{"Term", ``},    // Term
				{"Granted", ``}, // bool
			},
		},
		{
			Name: "AppendToLog",
			Doc:  "// AppendToLog is sent by the leader to tell followers to append an entry.  If cmds\n// is empty, this is a keep alive message (at a random interval after a keep alive, followers\n// will initiate a new round of voting).\n//   term -- the current term of the sender\n//   leaderId -- the id of the sender\n//   prevIndex -- the index of the log entry immediately preceding cmds\n//   prevTerm -- the term of the log entry immediately preceding cmds.  The receiver must have\n//               received the previous index'd entry and it must have had the same term.  Otherwise\n//               an error is returned.\n//   leaderCommit -- the index of the last committed entry, i.e., the one a quorum has gauranteed\n//                   to have logged.\n//   cmds -- sequential log entries starting at prevIndex+1",
			InArgs: []rpc.ArgDesc{
				{"term", ``},         // Term
				{"leaderId", ``},     // string
				{"prevIndex", ``},    // Index
				{"prevTerm", ``},     // Term
				{"leaderCommit", ``}, // Index
				{"cmds", ``},         // []LogEntry
			},
		},
		{
			Name: "Append",
			Doc:  "// Append is sent to the leader by followers.  Only the leader is allowed to send AppendToLog.\n// If a follower receives an Append() call it performs an Append() to the leader to run the actual\n// Raft algorithm.  The leader will respond after it has RaftClient.Apply()ed the command.\n//\n// Returns the term and index of the append entry or an error.",
			InArgs: []rpc.ArgDesc{
				{"cmd", ``}, // []byte
			},
			OutArgs: []rpc.ArgDesc{
				{"term", ``},  // Term
				{"index", ``}, // Index
			},
		},
		{
			Name: "Committed",
			Doc:  "// Committed returns the commit index of the leader.",
			OutArgs: []rpc.ArgDesc{
				{"index", ``}, // Index
			},
		},
		{
			Name: "InstallSnapshot",
			Doc:  "// InstallSnapshot is sent from the leader to follower to install the given snapshot.  It is\n// sent when it becomes apparent that the leader does not have log entries needed by the follower\n// to progress.  'term' and 'index' represent the last LogEntry RaftClient.Apply()ed to the\n// snapshot.",
			InArgs: []rpc.ArgDesc{
				{"term", ``},         // Term
				{"leaderId", ``},     // string
				{"appliedTerm", ``},  // Term
				{"appliedIndex", ``}, // Index
			},
		},
	},
}

// raftProtoInstallSnapshotServerStream is the server stream for raftProto.InstallSnapshot.
type raftProtoInstallSnapshotServerStream interface {
	// RecvStream returns the receiver side of the raftProto.InstallSnapshot server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// raftProtoInstallSnapshotServerCall represents the context passed to raftProto.InstallSnapshot.
type raftProtoInstallSnapshotServerCall interface {
	rpc.ServerCall
	raftProtoInstallSnapshotServerStream
}

// raftProtoInstallSnapshotServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements raftProtoInstallSnapshotServerCall.
type raftProtoInstallSnapshotServerCallStub struct {
	rpc.StreamServerCall
	valRecv []byte
	errRecv error
}

// Init initializes raftProtoInstallSnapshotServerCallStub from rpc.StreamServerCall.
func (s *raftProtoInstallSnapshotServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the raftProto.InstallSnapshot server stream.
func (s *raftProtoInstallSnapshotServerCallStub) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implraftProtoInstallSnapshotServerCallRecv{s}
}

type implraftProtoInstallSnapshotServerCallRecv struct {
	s *raftProtoInstallSnapshotServerCallStub
}

func (s implraftProtoInstallSnapshotServerCallRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implraftProtoInstallSnapshotServerCallRecv) Value() []byte {
	return s.s.valRecv
}
func (s implraftProtoInstallSnapshotServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}
