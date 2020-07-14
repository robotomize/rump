package main

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
	_ "net/http/pprof"
	"rump/internal/codec/msgpack"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"rump/internal/pb"
	"rump/internal/server"
	"rump/internal/server/gnet"
	"rump/internal/setup"
	"rump/internal/shutdown"
	"sync"
)

func main() {
	ctx, cancel := shutdown.New()
	defer cancel()
	logger := logging.FromContext(ctx)
	errCh := make(chan error, 1)

	go func() {
		logger := logging.FromContext(ctx)
		for err := range errCh {
			logger.Error(err)
		}
	}()

	var config server.Config

	env, err := setup.Setup(ctx, &config)
	if err != nil {
		errCh <- err
	}
	srv, err := server.New(env)
	if err != nil {
		logger.Fatal("server: %w", err)
	}
	state := gamestate.NewState(ctx, msgpack.New())

	defer close(errCh)
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		if err := serveGRPC(ctx, logger, srv, state); err != nil {
			errCh <- err
		}
	}()

	go func() {
		defer wg.Done()
		listener, err := gnet.New(env)
		if err != nil {
			errCh <- err
		}
		udp := server.NewUDPServer(listener, gamestate.NewUDPHandler(logger, state).SyncPlayerProtobuf)
		if err := serveUDP(ctx, srv, udp); err != nil {
			errCh <- err
		}
	}()

	go func() {
		if err := http.ListenAndServe("0.0.0.0:8181", nil); err != nil {
			errCh <- err
		}
	}()

	wg.Wait()
}

func serveGRPC(ctx context.Context, logger logging.Logger, srv *server.Server, state *gamestate.State) error {
	grpcServer := grpc.NewServer()
	pb.RegisterSyncStateServer(grpcServer, gamestate.NewGRPCHandler(logger, state))
	return srv.ServeGRPC(ctx, grpcServer)
}

func serveUDP(ctx context.Context, srv *server.Server, udp *server.UDPServer) error {
	return srv.ServeUDP(ctx, udp)
}
