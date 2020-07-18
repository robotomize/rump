package gamestate

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/valyala/fastrand"
	"rump/internal/codec"
	"rump/internal/logging"
	"rump/internal/pb"
	"rump/internal/utils"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrPlayerNotFound = fmt.Errorf("игрок не найден в контексте сервиса")
)

var playerPool = sync.Pool{
	New: func() interface{} { return &Player{} },
}

func GetPlayerPool() (p *Player) {
	ifc := playerPool.Get()
	if ifc != nil {
		p = ifc.(*Player)
	}
	return
}

var syncPbPlayerPool = sync.Pool{
	New: func() interface{} { return &pb.SyncPos{} },
}

func GetSyncPbPool() (p *pb.SyncPos) {
	ifc := syncPbPlayerPool.Get()
	if ifc != nil {
		p = ifc.(*pb.SyncPos)
	}
	return
}

func PutSyncPbPool(p *pb.SyncPos) {
	syncPbPlayerPool.Put(p)
}

func PutPlayerPool(p *Player) {
	playerPool.Put(p)
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
	TimeStamp int64
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

type Players map[uint32]*Player

func NewState(ctx context.Context, cd codec.Codec) *State {
	r := &State{
		mu:      sync.RWMutex{},
		players: Players{},
		logger:  logging.FromContext(ctx),
		codec:   cd,
	}
	return r
}

type State struct {
	mu      sync.RWMutex
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
	r.state.SyncPlayer(player)
	if atomic.LoadUint32(&cnt) == 25*1000 {
		r.logger.Debugf("%s", time.Now().UnixNano()-player.TimeStamp)
		atomic.StoreUint32(&cnt, 0)
	}
	atomic.AddUint32(&cnt, 1)
	return &pb.SyncPlayerResponse{ID: player.ID}, nil
}

type UDPHandler struct {
	logger logging.Logger
	state  *State
}

func NewUDPHandler(logger logging.Logger, state *State) *UDPHandler {
	return &UDPHandler{
		logger: logger,
		state:  state,
	}
}

func (r *UDPHandler) SyncPlayerBinary(cnt uint32, payload []byte) error {
	if r.state.codec == nil {
		return fmt.Errorf("кодек для декодирования данных не передан")
	}
	player := GetPlayerPool()
	defer PutPlayerPool(player)

	b := utils.GetBuffer()
	b.Write(payload)
	defer utils.PutBuffer(b)

	if err := binary.Read(b, binary.LittleEndian, player); err != nil {
		r.logger.Debugf("%q", err)
		return err
	}
	if atomic.LoadUint32(&cnt) == 25*1000 {
		r.logger.Debugf("%s", time.Duration(time.Now().UnixNano()-player.TimeStamp))
	}
	return nil
}

func (r *UDPHandler) SyncPlayerCodec(cnt uint32, payload []byte) error {
	if r.state.codec == nil {
		return fmt.Errorf("кодек для декодирования данных не передан")
	}
	player := GetPlayerPool()
	defer PutPlayerPool(player)
	if err := r.state.codec.Decode(payload, player); err != nil {
		return err
	}
	r.state.SyncPlayer(player)
	if atomic.LoadUint32(&cnt) == 25*1000 {
		r.logger.Debugf("%s", time.Duration(time.Now().UnixNano()-player.TimeStamp))
	}
	return nil
}

func (r *UDPHandler) SyncPlayerProtobuf(cnt uint32, payload []byte) error {
	in := GetSyncPbPool()
	if err := proto.Unmarshal(payload, in); err != nil {
		return err
	}
	PutSyncPbPool(in)
	player := GetPlayerPool()
	defer PutPlayerPool(player)
	player = NewPlayer(in.ID, Vector{X: in.Pos.X, Y: in.Pos.Y, Z: in.Pos.Z})
	player.TimeStamp = in.Timestamp
	r.state.SyncPlayer(player)
	if atomic.LoadUint32(&cnt) == 25*1000 {
		r.logger.Debugf("%s", time.Duration(time.Now().UnixNano()-player.TimeStamp))
	}
	return nil
}
