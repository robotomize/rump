package gamestatepb

import (
	"context"
	"rump/internal/gamestate"
	"rump/internal/pb"
)

type Srv struct {
	pb.UnimplementedSyncStateServer

	state *gamestate.State
}

func New(state *gamestate.State) *Srv {
	return &Srv{
		UnimplementedSyncStateServer: pb.UnimplementedSyncStateServer{},
		state:                        state,
	}
}

func (r *Srv) RcvPosition(ctx context.Context, req *pb.RcvPositionRequest) (*pb.RcvPositionResponse, error) {
	player, err := r.state.RcvPosition(req.ID)
	if err != nil {
		return nil, err
	}
	return &pb.RcvPositionResponse{
		ID: player.ID,
		Pos: &pb.Vector3{
			X: player.Pos.X,
			Y: player.Pos.Y,
			Z: player.Pos.Z,
		},
	}, nil
}

func (r *Srv) SyncPosition(ctx context.Context, req *pb.SyncPositionRequest) (*pb.SyncPositionResponse, error) {
	return (*pb.SyncPositionResponse)(nil), nil
}
