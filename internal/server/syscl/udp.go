package syscl

import (
	"context"
	"fmt"
	"rump/internal/logging"
	"rump/internal/server"
	"rump/internal/srvenv"
	"rump/internal/utils"
	"syscall"
)

type Server struct {
	env      *srvenv.SrvEnv
	sockAddr syscall.Sockaddr
}

func New(config *srvenv.SrvEnv) (*Server, error) {
	addr, err := utils.ParseAddr(config.SrvConfig.Addr)
	if err != nil {
		return nil, err
	}
	return &Server{sockAddr: addr}, nil
}

func (s *Server) ServeUDP(ctx context.Context, fn ...server.HandleFn) error {
	var done func()
	logger := logging.FromContext(ctx)
	ctx, done = context.WithCancel(ctx)
	errCh := make(chan error, 1)
	serverFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		logger.Fatal(err)
	}

	if err := syscall.Bind(serverFd, s.sockAddr); err != nil {
		logger.Error(err)
	}

	fmt.Println("server: сервер запущен ", s.sockAddr)
	go func() {
		<-ctx.Done()
		logger.Debugf("server: завершение по контексту UDP")
		_ = syscall.Close(serverFd)
	}()
	buf := make([]byte, 1500)
	go func() {
		defer func() {
			fmt.Println("exit")
		}()
		for {
			n, _, err := syscall.Recvfrom(serverFd, buf[:], 0)
			if err != nil {
				errCh <- err
				if err.Error() == "bad file descriptor" {
					fmt.Println("hello", err)
					return
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
		return fmt.Errorf("ошибка graceful shutdown UDP: %w", err)
	}
}
