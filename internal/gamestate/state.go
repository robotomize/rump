package gamestate

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/valyala/fastrand"
	"rump/internal/codec"
	"rump/internal/logging"
	"rump/internal/pb"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrPlayerNotFound = fmt.Errorf("игрок не найден в контексте сервиса")
)

var syncPbPlayerPool = sync.Pool{
	New: func() interface{} { return &pb.SyncPos{} },
}

func GetSyncPoll() (p *pb.SyncPos) {
	ifc := syncPbPlayerPool.Get()
	if ifc != nil {
		p = ifc.(*pb.SyncPos)
	}
	return
}

func PutSyncPoll(p *pb.SyncPos) {
	syncPbPlayerPool.Put(p)
}

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

func GenerateVector() Vector {
	division := float64(fastrand.Uint32n(1<<4) + 1)
	return Vector{
		X: float64(fastrand.Uint32()+1) / division,
		Y: float64(fastrand.Uint32()+1) / division,
		Z: float64(fastrand.Uint32()+1) / division,
	}
}

type Player struct {
	ID        uint32
	Pos       Vector
	TimeStamp time.Time
}

func GeneratePlayer() *Player {
	return &Player{
		ID:  fastrand.Uint32n(1000) + 1,
		Pos: GenerateVector(),
	}
}

func NewPlayer(ID uint32, pos Vector) *Player {
	return &Player{
		ID:  ID,
		Pos: pos,
	}
}

type Option func(*State)

type Options struct {
	debug bool
}

func NewDebug() Option {
	return func(state *State) {
		state.opts.debug = true
	}
}

type Players map[uint32]*Player

func NewState(ctx context.Context, codec codec.Codec, opts ...Option) *State {
	r := &State{
		mu:      sync.RWMutex{},
		codec:   codec,
		players: Players{},
		logger:  logging.FromContext(ctx),
	}
	for _, f := range opts {
		f(r)
	}
	return r
}

type State struct {
	mu      sync.RWMutex
	opts    Options
	ctx     context.Context
	codec   codec.Codec
	players Players
	logger  logging.Logger
}

func (r *State) RcvPlayer(ID uint32) (*Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if player, ok := r.players[ID]; ok {
		return &Player{
			ID: ID,
			Pos: Vector{
				X: player.Pos.X,
				Y: player.Pos.Y,
				Z: player.Pos.Z,
			},
			TimeStamp: player.TimeStamp,
		}, nil
	}
	return nil, ErrPlayerNotFound
}

func (r *State) SyncPlayer(player *Player) {
	r.mu.Lock()
	r.players[player.ID] = player
	r.mu.Unlock()
}

type GRPCHandler struct {
	pb.UnimplementedSyncStateServer

	logger logging.Logger
	state  *State
}

func NewGRPCHandler(logger logging.Logger, state *State) *GRPCHandler {
	return &GRPCHandler{
		UnimplementedSyncStateServer: pb.UnimplementedSyncStateServer{},
		logger:                       logger,
		state:                        state,
	}
}

func (r *GRPCHandler) RcvPlayer(ctx context.Context, req *pb.RcvPlayerRequest) (*pb.RcvPlayerResponse, error) {
	player, err := r.state.RcvPlayer(req.ID)
	if err != nil {
		return nil, err
	}
	return &pb.RcvPlayerResponse{
		ID: player.ID,
		Pos: &pb.Vector3{
			X: player.Pos.X,
			Y: player.Pos.Y,
			Z: player.Pos.Z,
		},
	}, nil
}

var cnt uint32

func (r *GRPCHandler) SyncPlayer(ctx context.Context, req *pb.SyncPlayerRequest) (*pb.SyncPlayerResponse, error) {
	player := NewPlayer(req.ID, Vector{X: req.Pos.X, Y: req.Pos.Y, Z: req.Pos.Z})
	player.TimeStamp = time.Unix(0, req.Timestamp)
	r.state.SyncPlayer(player)
	if atomic.LoadUint32(&cnt) == 25*1000 {
		fmt.Println(time.Since(player.TimeStamp))
		atomic.StoreUint32(&cnt, 0)
	}
	atomic.AddUint32(&cnt, 1)
	return &pb.SyncPlayerResponse{ID: player.ID}, nil
}

type UDPHandler struct {
	logger logging.Logger
	state  *State
	codec.Codec
}

func NewUDPHandler(logger logging.Logger, state *State) *UDPHandler {
	return &UDPHandler{
		logger: logger,
		state:  state,
	}
}

func (r *UDPHandler) SyncPlayerCodec(cnt uint32, payload []byte) error {
	player := Player{}
	if err := r.Codec.Decode(payload, &player); err != nil {
		return err
	}
	r.state.SyncPlayer(&player)
	if atomic.LoadUint32(&cnt) == 25*1000 {
		fmt.Println(time.Since(player.TimeStamp))
	}
	return nil
}

func (r *UDPHandler) SyncPlayerProtobuf(cnt uint32, payload []byte) error {
	in := GetSyncPoll()
	if err := proto.Unmarshal(payload, in); err != nil {
		return err
	}
	PutSyncPoll(in)
	player := NewPlayer(in.ID, Vector{X: in.Pos.X, Y: in.Pos.Y, Z: in.Pos.Z})
	player.TimeStamp = time.Unix(0, in.Timestamp)
	r.state.SyncPlayer(player)
	if atomic.LoadUint32(&cnt) == 25*1000 {
		fmt.Println(time.Since(player.TimeStamp))
	}
	return nil
}
