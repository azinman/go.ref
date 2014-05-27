// This file was auto-generated by the veyron vdl tool.
// Source: service.vdl

package rockpaperscissors

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdl "veyron2/vdl"
	_gen_wiretype "veyron2/wiretype"
)

// A GameID is used to uniquely identify a game within one Judge.
type GameID struct {
	ID string
}

// GameOptions specifies the parameters of a game.
type GameOptions struct {
	NumRounds int32       // The number of rounds that a player must win to win the game.
	GameType  GameTypeTag // The type of game to play: Classic or LizardSpock.
}

type GameTypeTag byte

type PlayerAction struct {
	Move string // The move that the player wants to make.
	Quit bool   // Whether the player wants to quit the game.
}

type JudgeAction struct {
	PlayerNum    int32     // The player's number.
	OpponentName string    // The name of the opponent.
	MoveOptions  []string  // A list of allowed moves that the player must choose from. Not always present.
	RoundResult  Round     // The result of the previous round. Not always present.
	Score        ScoreCard // The result of the game. Not always present.
}

// Round represents the state of a round.
type Round struct {
	Moves       [2]string // Each player's move.
	Winner      WinnerTag // Who won the round.
	StartTimeNS int64     // The time at which the round started.
	EndTimeNS   int64     // The time at which the round ended.
}

// WinnerTag is a type used to indicate whether a round or a game was a draw,
// was won by player 1 or was won by player 2.
type WinnerTag byte

// PlayResult is the value returned by the Play method. It indicates the outcome of the game.
type PlayResult struct {
	YouWon bool // True if the player receiving the result won the game.
}

type ScoreCard struct {
	Opts        GameOptions // The game options.
	Judge       string      // The name of the judge.
	Players     []string    // The name of the players.
	Rounds      []Round     // The outcome of each round.
	StartTimeNS int64       // The time at which the game started.
	EndTimeNS   int64       // The time at which the game ended.
	Winner      WinnerTag   // Who won the game.
}

const (
	Classic = GameTypeTag(0) // Rock-Paper-Scissors

	LizardSpock = GameTypeTag(1) // Rock-Paper-Scissors-Lizard-Spock

	Draw = WinnerTag(0)

	Player1 = WinnerTag(1)

	Player2 = WinnerTag(2)
)

// Judge is the interface the client binds and uses.
// Judge_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Judge_ExcludingUniversal interface {
	// CreateGame creates a new game with the given game options and returns a game
	// identifier that can be used by the players to join the game.
	CreateGame(Opts GameOptions, opts ..._gen_ipc.ClientCallOpt) (reply GameID, err error)
	// Play lets a player join an existing game and play.
	Play(ID GameID, opts ..._gen_ipc.ClientCallOpt) (reply JudgePlayStream, err error)
}
type Judge interface {
	_gen_ipc.UniversalServiceMethods
	Judge_ExcludingUniversal
}

// JudgeService is the interface the server implements.
type JudgeService interface {

	// CreateGame creates a new game with the given game options and returns a game
	// identifier that can be used by the players to join the game.
	CreateGame(context _gen_ipc.Context, Opts GameOptions) (reply GameID, err error)
	// Play lets a player join an existing game and play.
	Play(context _gen_ipc.Context, ID GameID, stream JudgeServicePlayStream) (reply PlayResult, err error)
}

// JudgePlayStream is the interface for streaming responses of the method
// Play in the service interface Judge.
type JudgePlayStream interface {

	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item PlayerAction) error

	// CloseSend indicates to the server that no more items will be sent; server
	// Recv calls will receive io.EOF after all sent items.  Subsequent calls to
	// Send on the client will fail.  This is an optional call - it's used by
	// streaming clients that need the server to receive the io.EOF terminator.
	CloseSend() error

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item JudgeAction, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (reply PlayResult, err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the JudgePlayStream interface that is not exported.
type implJudgePlayStream struct {
	clientCall _gen_ipc.ClientCall
}

func (c *implJudgePlayStream) Send(item PlayerAction) error {
	return c.clientCall.Send(item)
}

func (c *implJudgePlayStream) CloseSend() error {
	return c.clientCall.CloseSend()
}

func (c *implJudgePlayStream) Recv() (item JudgeAction, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implJudgePlayStream) Finish() (reply PlayResult, err error) {
	if ierr := c.clientCall.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implJudgePlayStream) Cancel() {
	c.clientCall.Cancel()
}

// JudgeServicePlayStream is the interface for streaming responses of the method
// Play in the service interface Judge.
type JudgeServicePlayStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item JudgeAction) error

	// Recv fills itemptr with the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item PlayerAction, err error)
}

// Implementation of the JudgeServicePlayStream interface that is not exported.
type implJudgeServicePlayStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implJudgeServicePlayStream) Send(item JudgeAction) error {
	return s.serverCall.Send(item)
}

func (s *implJudgeServicePlayStream) Recv() (item PlayerAction, err error) {
	err = s.serverCall.Recv(&item)
	return
}

// BindJudge returns the client stub implementing the Judge
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindJudge(name string, opts ..._gen_ipc.BindOpt) (Judge, error) {
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
	stub := &clientStubJudge{client: client, name: name}

	return stub, nil
}

// NewServerJudge creates a new server stub.
//
// It takes a regular server implementing the JudgeService
// interface, and returns a new server stub.
func NewServerJudge(server JudgeService) interface{} {
	return &ServerStubJudge{
		service: server,
	}
}

// clientStubJudge implements Judge.
type clientStubJudge struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubJudge) CreateGame(Opts GameOptions, opts ..._gen_ipc.ClientCallOpt) (reply GameID, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "CreateGame", []interface{}{Opts}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubJudge) Play(ID GameID, opts ..._gen_ipc.ClientCallOpt) (reply JudgePlayStream, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Play", []interface{}{ID}, opts...); err != nil {
		return
	}
	reply = &implJudgePlayStream{clientCall: call}
	return
}

func (__gen_c *clientStubJudge) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubJudge) Signature(opts ..._gen_ipc.ClientCallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubJudge) GetMethodTags(method string, opts ..._gen_ipc.ClientCallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubJudge wraps a server that implements
// JudgeService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubJudge struct {
	service JudgeService
}

func (__gen_s *ServerStubJudge) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "CreateGame":
		return []interface{}{}, nil
	case "Play":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubJudge) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["CreateGame"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Opts", Type: 66},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 67},
			{Name: "", Type: 68},
		},
	}
	result.Methods["Play"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "ID", Type: 67},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 69},
			{Name: "", Type: 68},
		},
		InStream:  70,
		OutStream: 76,
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "veyron/examples/rockpaperscissors.GameTypeTag", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x24, Name: "NumRounds"},
				_gen_wiretype.FieldType{Type: 0x41, Name: "GameType"},
			},
			"veyron/examples/rockpaperscissors.GameOptions", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "ID"},
			},
			"veyron/examples/rockpaperscissors.GameID", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x2, Name: "YouWon"},
			},
			"veyron/examples/rockpaperscissors.PlayResult", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Move"},
				_gen_wiretype.FieldType{Type: 0x2, Name: "Quit"},
			},
			"veyron/examples/rockpaperscissors.PlayerAction", []string(nil)},
		_gen_wiretype.ArrayType{Elem: 0x3, Len: 0x2, Name: "", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "veyron/examples/rockpaperscissors.WinnerTag", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x47, Name: "Moves"},
				_gen_wiretype.FieldType{Type: 0x48, Name: "Winner"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "StartTimeNS"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "EndTimeNS"},
			},
			"veyron/examples/rockpaperscissors.Round", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x49, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x42, Name: "Opts"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "Judge"},
				_gen_wiretype.FieldType{Type: 0x3d, Name: "Players"},
				_gen_wiretype.FieldType{Type: 0x4a, Name: "Rounds"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "StartTimeNS"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "EndTimeNS"},
				_gen_wiretype.FieldType{Type: 0x48, Name: "Winner"},
			},
			"veyron/examples/rockpaperscissors.ScoreCard", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x24, Name: "PlayerNum"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "OpponentName"},
				_gen_wiretype.FieldType{Type: 0x3d, Name: "MoveOptions"},
				_gen_wiretype.FieldType{Type: 0x49, Name: "RoundResult"},
				_gen_wiretype.FieldType{Type: 0x4b, Name: "Score"},
			},
			"veyron/examples/rockpaperscissors.JudgeAction", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubJudge) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubJudge) CreateGame(call _gen_ipc.ServerCall, Opts GameOptions) (reply GameID, err error) {
	reply, err = __gen_s.service.CreateGame(call, Opts)
	return
}

func (__gen_s *ServerStubJudge) Play(call _gen_ipc.ServerCall, ID GameID) (reply PlayResult, err error) {
	stream := &implJudgeServicePlayStream{serverCall: call}
	reply, err = __gen_s.service.Play(call, ID, stream)
	return
}

// Player can receive challenges from other players.
// Player is the interface the client binds and uses.
// Player_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Player_ExcludingUniversal interface {
	// Challenge is used by other players to challenge this player to a game. If
	// the challenge is accepted, the method returns true.
	Challenge(Address string, ID GameID, opts ..._gen_ipc.ClientCallOpt) (reply bool, err error)
}
type Player interface {
	_gen_ipc.UniversalServiceMethods
	Player_ExcludingUniversal
}

// PlayerService is the interface the server implements.
type PlayerService interface {

	// Challenge is used by other players to challenge this player to a game. If
	// the challenge is accepted, the method returns true.
	Challenge(context _gen_ipc.Context, Address string, ID GameID) (reply bool, err error)
}

// BindPlayer returns the client stub implementing the Player
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindPlayer(name string, opts ..._gen_ipc.BindOpt) (Player, error) {
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
	stub := &clientStubPlayer{client: client, name: name}

	return stub, nil
}

// NewServerPlayer creates a new server stub.
//
// It takes a regular server implementing the PlayerService
// interface, and returns a new server stub.
func NewServerPlayer(server PlayerService) interface{} {
	return &ServerStubPlayer{
		service: server,
	}
}

// clientStubPlayer implements Player.
type clientStubPlayer struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubPlayer) Challenge(Address string, ID GameID, opts ..._gen_ipc.ClientCallOpt) (reply bool, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Challenge", []interface{}{Address, ID}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubPlayer) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubPlayer) Signature(opts ..._gen_ipc.ClientCallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubPlayer) GetMethodTags(method string, opts ..._gen_ipc.ClientCallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubPlayer wraps a server that implements
// PlayerService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubPlayer struct {
	service PlayerService
}

func (__gen_s *ServerStubPlayer) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Challenge":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubPlayer) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Challenge"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Address", Type: 3},
			{Name: "ID", Type: 65},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 2},
			{Name: "", Type: 66},
		},
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "ID"},
			},
			"veyron/examples/rockpaperscissors.GameID", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubPlayer) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubPlayer) Challenge(call _gen_ipc.ServerCall, Address string, ID GameID) (reply bool, err error) {
	reply, err = __gen_s.service.Challenge(call, Address, ID)
	return
}

// ScoreKeeper receives the outcome of games from Judges.
// ScoreKeeper is the interface the client binds and uses.
// ScoreKeeper_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type ScoreKeeper_ExcludingUniversal interface {
	Record(Score ScoreCard, opts ..._gen_ipc.ClientCallOpt) (err error)
}
type ScoreKeeper interface {
	_gen_ipc.UniversalServiceMethods
	ScoreKeeper_ExcludingUniversal
}

// ScoreKeeperService is the interface the server implements.
type ScoreKeeperService interface {
	Record(context _gen_ipc.Context, Score ScoreCard) (err error)
}

// BindScoreKeeper returns the client stub implementing the ScoreKeeper
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindScoreKeeper(name string, opts ..._gen_ipc.BindOpt) (ScoreKeeper, error) {
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
	stub := &clientStubScoreKeeper{client: client, name: name}

	return stub, nil
}

// NewServerScoreKeeper creates a new server stub.
//
// It takes a regular server implementing the ScoreKeeperService
// interface, and returns a new server stub.
func NewServerScoreKeeper(server ScoreKeeperService) interface{} {
	return &ServerStubScoreKeeper{
		service: server,
	}
}

// clientStubScoreKeeper implements ScoreKeeper.
type clientStubScoreKeeper struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubScoreKeeper) Record(Score ScoreCard, opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Record", []interface{}{Score}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubScoreKeeper) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubScoreKeeper) Signature(opts ..._gen_ipc.ClientCallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubScoreKeeper) GetMethodTags(method string, opts ..._gen_ipc.ClientCallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubScoreKeeper wraps a server that implements
// ScoreKeeperService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubScoreKeeper struct {
	service ScoreKeeperService
}

func (__gen_s *ServerStubScoreKeeper) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Record":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubScoreKeeper) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Record"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Score", Type: 71},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 72},
		},
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "veyron/examples/rockpaperscissors.GameTypeTag", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x24, Name: "NumRounds"},
				_gen_wiretype.FieldType{Type: 0x41, Name: "GameType"},
			},
			"veyron/examples/rockpaperscissors.GameOptions", []string(nil)},
		_gen_wiretype.ArrayType{Elem: 0x3, Len: 0x2, Name: "", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "veyron/examples/rockpaperscissors.WinnerTag", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x43, Name: "Moves"},
				_gen_wiretype.FieldType{Type: 0x44, Name: "Winner"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "StartTimeNS"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "EndTimeNS"},
			},
			"veyron/examples/rockpaperscissors.Round", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x45, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x42, Name: "Opts"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "Judge"},
				_gen_wiretype.FieldType{Type: 0x3d, Name: "Players"},
				_gen_wiretype.FieldType{Type: 0x46, Name: "Rounds"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "StartTimeNS"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "EndTimeNS"},
				_gen_wiretype.FieldType{Type: 0x44, Name: "Winner"},
			},
			"veyron/examples/rockpaperscissors.ScoreCard", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubScoreKeeper) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubScoreKeeper) Record(call _gen_ipc.ServerCall, Score ScoreCard) (err error) {
	err = __gen_s.service.Record(call, Score)
	return
}

// RockPaperScissors is the interface the client binds and uses.
// RockPaperScissors_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type RockPaperScissors_ExcludingUniversal interface {
	Judge_ExcludingUniversal
	// Player can receive challenges from other players.
	Player_ExcludingUniversal
	// ScoreKeeper receives the outcome of games from Judges.
	ScoreKeeper_ExcludingUniversal
}
type RockPaperScissors interface {
	_gen_ipc.UniversalServiceMethods
	RockPaperScissors_ExcludingUniversal
}

// RockPaperScissorsService is the interface the server implements.
type RockPaperScissorsService interface {
	JudgeService
	// Player can receive challenges from other players.
	PlayerService
	// ScoreKeeper receives the outcome of games from Judges.
	ScoreKeeperService
}

// BindRockPaperScissors returns the client stub implementing the RockPaperScissors
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindRockPaperScissors(name string, opts ..._gen_ipc.BindOpt) (RockPaperScissors, error) {
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
	stub := &clientStubRockPaperScissors{client: client, name: name}
	stub.Judge_ExcludingUniversal, _ = BindJudge(name, client)
	stub.Player_ExcludingUniversal, _ = BindPlayer(name, client)
	stub.ScoreKeeper_ExcludingUniversal, _ = BindScoreKeeper(name, client)

	return stub, nil
}

// NewServerRockPaperScissors creates a new server stub.
//
// It takes a regular server implementing the RockPaperScissorsService
// interface, and returns a new server stub.
func NewServerRockPaperScissors(server RockPaperScissorsService) interface{} {
	return &ServerStubRockPaperScissors{
		ServerStubJudge:       *NewServerJudge(server).(*ServerStubJudge),
		ServerStubPlayer:      *NewServerPlayer(server).(*ServerStubPlayer),
		ServerStubScoreKeeper: *NewServerScoreKeeper(server).(*ServerStubScoreKeeper),
		service:               server,
	}
}

// clientStubRockPaperScissors implements RockPaperScissors.
type clientStubRockPaperScissors struct {
	Judge_ExcludingUniversal
	Player_ExcludingUniversal
	ScoreKeeper_ExcludingUniversal

	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubRockPaperScissors) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubRockPaperScissors) Signature(opts ..._gen_ipc.ClientCallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubRockPaperScissors) GetMethodTags(method string, opts ..._gen_ipc.ClientCallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubRockPaperScissors wraps a server that implements
// RockPaperScissorsService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubRockPaperScissors struct {
	ServerStubJudge
	ServerStubPlayer
	ServerStubScoreKeeper

	service RockPaperScissorsService
}

func (__gen_s *ServerStubRockPaperScissors) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	if resp, err := __gen_s.ServerStubJudge.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	if resp, err := __gen_s.ServerStubPlayer.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	if resp, err := __gen_s.ServerStubScoreKeeper.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	return nil, nil
}

func (__gen_s *ServerStubRockPaperScissors) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}

	result.TypeDefs = []_gen_vdl.Any{}
	var ss _gen_ipc.ServiceSignature
	var firstAdded int
	ss, _ = __gen_s.ServerStubJudge.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.InArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.OutArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= _gen_wiretype.TypeIDFirst {
			v.InStream += _gen_wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= _gen_wiretype.TypeIDFirst {
			v.OutStream += _gen_wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case _gen_wiretype.SliceType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.ArrayType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.MapType:
			if wt.Key >= _gen_wiretype.TypeIDFirst {
				wt.Key += _gen_wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.StructType:
			for _, fld := range wt.Fields {
				if fld.Type >= _gen_wiretype.TypeIDFirst {
					fld.Type += _gen_wiretype.TypeID(firstAdded)
				}
			}
			d = wt
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}
	ss, _ = __gen_s.ServerStubPlayer.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.InArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.OutArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= _gen_wiretype.TypeIDFirst {
			v.InStream += _gen_wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= _gen_wiretype.TypeIDFirst {
			v.OutStream += _gen_wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case _gen_wiretype.SliceType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.ArrayType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.MapType:
			if wt.Key >= _gen_wiretype.TypeIDFirst {
				wt.Key += _gen_wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.StructType:
			for _, fld := range wt.Fields {
				if fld.Type >= _gen_wiretype.TypeIDFirst {
					fld.Type += _gen_wiretype.TypeID(firstAdded)
				}
			}
			d = wt
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}
	ss, _ = __gen_s.ServerStubScoreKeeper.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.InArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.OutArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= _gen_wiretype.TypeIDFirst {
			v.InStream += _gen_wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= _gen_wiretype.TypeIDFirst {
			v.OutStream += _gen_wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case _gen_wiretype.SliceType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.ArrayType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.MapType:
			if wt.Key >= _gen_wiretype.TypeIDFirst {
				wt.Key += _gen_wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.StructType:
			for _, fld := range wt.Fields {
				if fld.Type >= _gen_wiretype.TypeIDFirst {
					fld.Type += _gen_wiretype.TypeID(firstAdded)
				}
			}
			d = wt
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}

	return result, nil
}

func (__gen_s *ServerStubRockPaperScissors) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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
