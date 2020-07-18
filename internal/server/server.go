package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"rump/internal/logging"
	"rump/internal/srvenv"
)

type UDPHandler func(uint32, []byte) error

type NonStreamListener interface {
	Listen(ctx context.Context, fn ...UDPHandler) error
}

type UDPServer struct {
	listener NonStreamListener
	handlers []UDPHandler
}

func NewUDPServer(listener NonStreamListener, handlers ...UDPHandler) *UDPServer {
	return &UDPServer{
		listener: listener,
		handlers: handlers,
	}
}

type Server struct {
	env *srvenv.SrvEnv
}

func New(env *srvenv.SrvEnv) (*Server, error) {
	return &Server{env: env}, nil
}

func (s *Server) ServeUDP(ctx context.Context, srv *UDPServer) error {
	logger := logging.FromContext(ctx)
	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()
		logger.Debugf("server: завершение по контексту UDP")
	}()
	logger.Debugf("server: сервис UDP запущен")
	if err := srv.listener.Listen(ctx, srv.handlers...); err != nil {
		return fmt.Errorf("server: ошибка запуска GRPC: %w", err)
	}
	logger.Debugf("server: сервис UDP остановлен")
	select {
	case err := <-errCh:
		return fmt.Errorf("server: ошибка graceful shutdown UDP: %w", err)
	}
}

func (s *Server) ServeGRPC(ctx context.Context, srv *grpc.Server) error {
	logger := logging.FromContext(ctx)
	errCh := make(chan error, 1)
	listener, err := net.Listen("tcp", s.env.SrvConfig.GetGRPCAddr())
	if err != nil {
		return fmt.Errorf("server: создание GRPC завершилось с ошибкой на %s: %w", s.env.SrvConfig.GetGRPCAddr(), err)
	}
	logger.Debugf("server: сервер GRPC запущен на порту %s", s.env.SrvConfig.GetGRPCAddr())
	go func() {
		<-ctx.Done()
		logger.Debugf("server: завершение по контексту GRPC")
		srv.GracefulStop()
	}()

	if err := srv.Serve(listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return fmt.Errorf("server: ошибка запуска GRPC: %w", err)
	}

	logger.Debugf("server: сервис GRPC остановлен")

	select {
	case err := <-errCh:
		return fmt.Errorf("server: ошибка graceful shutdown: %w", err)
	}
}
