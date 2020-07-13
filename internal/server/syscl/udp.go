package syscl

import (
	"context"
	"fmt"
	"rump/internal/logging"
	"rump/internal/server"
	"rump/internal/srvenv"
	"rump/internal/utils"
	"runtime"
	"sync"
	"syscall"
)

type Server struct {
	env      *srvenv.SrvEnv
	sockAddr syscall.Sockaddr
}

func New(config *srvenv.SrvEnv) (*Server, error) {
	addr, err := utils.ParseAddr(config.SrvConfig.SyncAddr)
	if err != nil {
		return nil, err
	}
	return &Server{sockAddr: addr, env: config}, nil
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

	fmt.Println("server: сервер запущен на", s.env.SrvConfig.SyncAddr)
	go func() {
		<-ctx.Done()
		logger.Debugf("server: завершение по контексту UDP")
		_ = syscall.Close(serverFd)
	}()

	rateCh := make(chan struct{}, runtime.NumCPU())
	defer close(rateCh)

	var wg sync.WaitGroup

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
				func() {
					b := make([]byte, len(buf[:n]))
					copy(b, buf[:n])
					worker(&wg, logger, b, handleFn, errCh, rateCh)
				}()
			}
		}
	}()
	select {
	case err := <-errCh:
		return fmt.Errorf("ошибка graceful shutdown UDP: %w", err)
	}
}

func worker(wg *sync.WaitGroup, logger logging.Logger, bytes []byte, fn server.HandleFn, errCh chan<- error, rateCh chan struct{}) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			logger.Error(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		rateCh <- struct{}{}
		if err := fn(bytes); err != nil {
			select {
			case errCh <- err:
				logger.Error(err)
			default:
			}
		}
		<-rateCh
	}()
}
