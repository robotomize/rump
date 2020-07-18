package app

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"rump/internal/pb"
	"rump/internal/server"
	"rump/internal/server/gnet"
	"rump/internal/setup"
	"sync"
)

type Options struct {
	profiling bool
}

type Option func(*App)

func EnableProfiling() Option {
	return func(a *App) {
		a.opts.profiling = true
	}
}

func New(config *server.Config, grpc *grpc.Server, g *gamestate.State, options ...Option) *App {
	a := &App{
		state:  g,
		grpc:   grpc,
		config: config,
	}
	for _, f := range options {
		f(a)
	}
	return a
}

type App struct {
	opts   Options
	state  *gamestate.State
	grpc   *grpc.Server
	config *server.Config
}

func (a *App) Run(ctx context.Context, wg *sync.WaitGroup, errCh chan error, state *gamestate.State, handlers ...server.UDPHandler) {
	logger := logging.FromContext(ctx)

	env, err := setup.Setup(ctx, a.config)
	if err != nil {
		errCh <- err
	}

	srv, err := server.New(env)
	if err != nil {
		errCh <- err
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := a.serveGRPC(ctx, logger, a.grpc, srv, state); err != nil {
			errCh <- err
		}
	}()

	go func() {
		defer wg.Done()
		listener, err := gnet.New(env)
		if err != nil {
			errCh <- err
		}
		udp := server.NewUDPServer(listener, handlers...)
		if err := a.serveUDP(ctx, srv, udp); err != nil {
			errCh <- err
		}
	}()
	if a.opts.profiling {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := http.ListenAndServe("0.0.0.0:8181", nil); err != nil {
				errCh <- err
			}
		}()
	}
}

func (a *App) serveGRPC(ctx context.Context, logger logging.Logger, grpc *grpc.Server, srv *server.Server, state *gamestate.State) error {
	pb.RegisterSyncStateServer(grpc, gamestate.NewGRPCHandler(logger, state))
	return srv.ServeGRPC(ctx, grpc)
}

func (a *App) serveUDP(ctx context.Context, srv *server.Server, udp *server.UDPServer) error {
	return srv.ServeUDP(ctx, udp)
}
