package evioudp

import (
	"context"
	"fmt"
	"github.com/tidwall/evio"
	"github.com/valyala/fastrand"
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

var cc uint32

func (s *Listener) Listen(ctx context.Context, fn ...server.UDPHandler) error {
	logger := logging.FromContext(ctx)
	logger.Debugf("server: сервер UDP запущен на порту %s", s.env.SrvConfig.GetUDPAddr())
	rateCh := make(chan struct{}, runtime.NumCPU()*4)
	defer close(rateCh)

	var (
		loops  = 20
		cnt    uint32
		events evio.Events
	)
	events.NumLoops = loops
	_ = cnt
	go func() {
		<-ctx.Done()
		logger.Debugf("server: завершение по контексту UDP")
	}()
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		for _, f := range fn {
			func() {
				if err := f(cnt, in); err != nil {
					logger.Errorf("server: ошибка при вызове UDP хэндлера %w", err)
				}
				if atomic.LoadUint32(&cnt) == 25*1000 {
					atomic.StoreUint32(&cnt, 0)
				}
				atomic.AddUint32(&cnt, 1)
				atomic.AddUint32(&cc, 1)
				if fastrand.Uint32n(100000) == fastrand.Uint32n(100000) {
					fmt.Println(atomic.LoadUint32(&cc))
				}
			}()
		}
		out = in
		return
	}
	return evio.Serve(events, fmt.Sprintf("%s://%s?reuseport=%t", "udp", s.env.SrvConfig.GetUDPAddr(), true))
}
