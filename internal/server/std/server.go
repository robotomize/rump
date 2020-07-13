package std

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"rump/internal/logging"
	"rump/internal/srvenv"
)

type Server struct {
	listener net.Listener
	conn     net.Conn
}

func New(config *srvenv.SrvEnv) (*Server, error) {
	s := &Server{}
	listener, err := net.Listen("tcp", config.SrvConfig.RcvAddr)
	if err != nil {
		return nil, fmt.Errorf("создание listener завершилось с ошибкой на %s: %w", config.SrvConfig, err)
	}
	s.listener = listener
	return s, nil
}

func (s *Server) ServeGRPC(ctx context.Context, srv *grpc.Server) error {
	logger := logging.FromContext(ctx)
	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()
		logger.Debugf("server: завершение по контексту GRPC")
		srv.GracefulStop()
	}()

	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return fmt.Errorf("ошибка запуска GRPC: %w", err)
	}

	logger.Debugf("server: сервис остановлен")

	select {
	case err := <-errCh:
		return fmt.Errorf("ошибка graceful shutdown: %w", err)
	}
}
