package syscalludp

import (
	"context"
	"fmt"
	"rump/internal/logging"
	"rump/internal/server"
	"rump/internal/srvenv"
	"rump/internal/utils"
	"runtime"
	"sync/atomic"
	"syscall"
)

func New(env *srvenv.SrvEnv) (*Listener, error) {
	addr, err := utils.ParseAddr(env.SrvConfig.GetUDPAddr())
	if err != nil {
		return nil, err
	}
	return &Listener{sockAddr: addr, env: env}, nil
}

type Listener struct {
	env      *srvenv.SrvEnv
	sockAddr syscall.Sockaddr
}

func (s *Listener) Listen(ctx context.Context, fn ...server.UDPHandler) error {
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
	fmt.Println("server: сервер запущен на", s.env.SrvConfig.GetUDPAddr())
	rateCh := make(chan struct{}, runtime.NumCPU()*4)
	defer close(rateCh)

	buf := make([]byte, 1500)
	go func() {
		defer func() {
			fmt.Println("exit")
		}()
		var (
			cnt uint32
		)
		_ = cnt
		for {
			n, _, err := syscall.Recvfrom(serverFd, buf[:], 0)
			if err != nil {
				errCh <- err
				if err.Error() == "bad file descriptor" {
					return
				}
			}
			if n <= 0 {
				done()
				return
			}
			for _, f := range fn {
				f := f
				go func() {
					defer func() {
						if err := recover(); err != nil {
							const size = 64 << 10
							buf := make([]byte, size)
							buf = buf[:runtime.Stack(buf, false)]
							logger.Error(err)
						}
					}()
					b := make([]byte, len(buf[:n]))
					copy(b, buf[:n])

					if err := f(cnt, b); err != nil {
						select {
						case errCh <- err:
							logger.Error(err)
						default:
						}
					}
					if atomic.LoadUint32(&cnt) == 25*10000 {
						atomic.StoreUint32(&cnt, 0)
					}
					atomic.AddUint32(&cnt, 1)
				}()
			}
		}
	}()
	select {
	case err := <-errCh:
		return fmt.Errorf("ошибка graceful shutdown UDP: %w", err)
	}
}
