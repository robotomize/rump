package gamestate

import (
	"context"
	"fmt"
	"github.com/valyala/fastrand"
	"rump/internal/codec"
	"rump/internal/logging"
	"sync"
	"time"
)

var (
	ErrPlayerNotFound = fmt.Errorf("игрок не найден в контексте сервиса")
)

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

type StateHandlerFn func(ctx context.Context, ID uint32) (*Player, error)

func (r *State) RcvPosition(ID uint32) (*Player, error) {
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

func (r *State) SyncPosition(bytes []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var p *Player
	if err := r.codec.Decode(bytes, &p); err != nil {
		return err
	}
	r.players[p.ID] = p
	fmt.Printf(
		"синхронизован игрок %d с координатами %f, %f, %f, latency %q",
		p.ID, p.Pos.X, p.Pos.Y, p.Pos.Z, time.Since(p.TimeStamp))
	return nil
}
