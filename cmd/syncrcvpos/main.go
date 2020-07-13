package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"rump/internal/codec/msgpack"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"rump/internal/pb"
	"rump/internal/pb/gamestatepb"
	"rump/internal/server"
	"rump/internal/server/std"
	"rump/internal/server/syscl"
	"rump/internal/setup"
	"rump/internal/shutdown"
	"sync"
)

func main() {
	ctx, cancel := shutdown.New()
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		logger := logging.FromContext(ctx)
		for err := range errCh {
			logger.Error(err)
		}
	}()
	state := gamestate.NewState(ctx, msgpack.New())
	pbGameState := gamestatepb.New(state)
	defer close(errCh)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := runGRPCun(ctx, pbGameState); err != nil {
			errCh <- err
		}
	}()
	go func() {
		defer wg.Done()
		if err := runUDP(ctx, state); err != nil {
			errCh <- err
		}
	}()
	wg.Wait()
}

func runGRPCun(ctx context.Context, pbSrv *gamestatepb.Srv) error {
	ctx, _ = shutdown.New()
	logger := logging.FromContext(ctx)
	var (
		config server.Config
	)
	env, err := setup.Setup(ctx, &config)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("server: %w", err)
	}
	srv, err := std.New(env)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("server: %w", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSyncStateServer(grpcServer, pbSrv)

	logger.Infof("запущен на порту :%s", env.SrvConfig.RcvAddr)
	return srv.ServeGRPC(ctx, grpcServer)
}

func runUDP(ctx context.Context, state *gamestate.State) error {
	ctx, _ = shutdown.New()
	logger := logging.FromContext(ctx)
	var config server.Config
	env, err := setup.Setup(ctx, &config)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("server: %w", err)
	}
	srv, err := syscl.New(env)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("server: %w", err)
	}
	return srv.ServeUDP(ctx, state.SyncPosition)
}
