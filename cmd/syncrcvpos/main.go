package main

import (
	"google.golang.org/grpc"
	_ "net/http/pprof"
	"rump/internal/app"
	"rump/internal/codec/encbinary"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"rump/internal/server"
	"rump/internal/shutdown"
	"sync"
)

func main() {
	var (
		config = &server.Config{}
		wg     sync.WaitGroup
	)
	ctx, cancel := shutdown.New()
	defer cancel()

	logger := logging.FromContext(ctx)

	errCh := make(chan error, 1)
	defer close(errCh)

	go func() {
		for err := range errCh {
			logger.Errorf("server: ", err)
			logger.Error(err)
		}
	}()

	state := gamestate.NewState(ctx, encbinary.New())
	realApp := app.New(config, grpc.NewServer(), state, app.EnableProfiling())
	realApp.Run(ctx, &wg, errCh, state, gamestate.NewUDPHandler(logger, state).SyncPlayerCodec)
	wg.Wait()
}
