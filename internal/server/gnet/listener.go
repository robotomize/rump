package gnet

import (
	"context"
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"github.com/valyala/fastrand"
	"rump/internal/logging"
	"rump/internal/server"
	"rump/internal/srvenv"
	"rump/internal/utils"
	"sync/atomic"
	"syscall"
)

func New(env *srvenv.SrvEnv) (*Listener, error) {
	addr, err := utils.ParseAddr(env.SrvConfig.GetUDPAddr())
	if err != nil {
		return nil, err
	}
	return &Listener{sockAddr: addr, env: env, srv: &srv{pool: goroutine.Default()}}, nil
}

type Listener struct {
	env      *srvenv.SrvEnv
	sockAddr syscall.Sockaddr
	srv      *srv
}

type srv struct {
	*gnet.EventServer
	pool   *goroutine.Pool
	fn     []server.UDPHandler
	logger logging.Logger
	cnt    uint32
}

var cc uint32

func (es *srv) React(in []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	//fmt.Println(in)
	for _, f := range es.fn {
		f := f
		func() {
			//defer func() {
			//	if err := recover(); err != nil {
			//		const size = 64 << 10
			//		buf := make([]byte, size)
			//		buf = buf[:runtime.Stack(buf, false)]
			//		es.logger.Error(err)
			//	}
			//}()
			if err := f(es.cnt, in); err != nil {
				es.logger.Error(err)
			}
			if atomic.LoadUint32(&es.cnt) == 25*1000 {
				atomic.StoreUint32(&es.cnt, 0)
			}
			atomic.AddUint32(&es.cnt, 1)
			atomic.AddUint32(&cc, 1)
			if fastrand.Uint32n(100000) == fastrand.Uint32n(100000) {
				fmt.Println(atomic.LoadUint32(&cc))
			}
		}()
	}
	in = out
	return
}

func (s *Listener) Listen(ctx context.Context, fn ...server.UDPHandler) error {
	s.srv.fn = fn
	defer s.srv.pool.Release()
	logger := logging.FromContext(ctx)
	s.srv.logger = logger
	return gnet.Serve(s.srv, "udp://:5555", gnet.WithNumEventLoop(300), gnet.WithMulticore(true), gnet.WithReusePort(true))
}
