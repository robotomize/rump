package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"rump/internal/codec/msgpack"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"rump/internal/server"
	"rump/internal/server/std"
	"rump/internal/setup"
	"rump/internal/shutdown"
	"rump/pb"
)

func main() {
	ctx, cancel := shutdown.New()
	defer cancel()

	if err := run(ctx); err != nil {
		logger := logging.FromContext(ctx)
		logger.Fatal(err)
	}
}

func run(ctx context.Context) error {
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
	config.Kind = server.KindTCP
	config.Addr = ":5577"
	srv, err := std.New(env)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("server: %w", err)
	}

	state := gamestate.NewState(ctx, msgpack.New())
	grpcServer := grpc.NewServer()
	pb.RegisterSyncStateServer(grpcServer, state)

	logger.Infof("запущен на порту :%s", env.SrvConfig.Addr)
	return srv.ServeGRPC(ctx, grpcServer)
}
