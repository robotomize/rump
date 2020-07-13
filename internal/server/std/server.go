package std

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"rump/internal/logging"
	"rump/internal/server"
	"rump/internal/srvenv"
)

type Server struct {
	listener net.Listener
	conn     net.Conn
}

func New(config *srvenv.SrvEnv) (*Server, error) {
	s := &Server{}
	if config.SrvConfig.IsTCP() {
		listener, err := net.Listen("tcp", config.SrvConfig.Addr)
		if err != nil {
			return nil, fmt.Errorf("создание listener завершилось с ошибкой на %s: %w", config.SrvConfig, err)
		}
		s.listener = listener
	}
	if config.SrvConfig.IsUDP() {
		udpAddr, err := net.ResolveUDPAddr("udp", config.SrvConfig.Addr)
		if err != nil {
			log.Fatalln(err)
		}
		conn, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			log.Fatalln(err)
		}
		s.conn = conn
	}
	return s, nil
}

func (s *Server) ServeUDP(ctx context.Context, fn ...server.HandleFn) error {
	var done func()
	logger := logging.FromContext(ctx)
	ctx, done = context.WithCancel(ctx)
	errCh := make(chan error, 1)

	go func() {
		<-ctx.Done()
		var err error
		logger.Debugf("server: завершение по контексту UDP")
		if s.listener != nil {
			err = s.listener.Close()
		} else {
			err = s.conn.Close()
		}
		select {
		case errCh <- err:
		default:
		}
	}()

	buf := make([]byte, 1500)
	go func() {
		for {
			n, err := s.conn.Read(buf[:])
			if err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
			if n <= 0 {
				done()
				return
			}
			for _, handleFn := range fn {
				if err := handleFn(buf[:n]); err != nil {
					select {
					case errCh <- err:
					default:
					}
				}
			}
		}
	}()
	select {
	case err := <-errCh:
		return fmt.Errorf("ошибка graceful shutdown GRPC: %w", err)
	}
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
