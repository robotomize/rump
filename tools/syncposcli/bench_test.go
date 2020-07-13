package main

import (
	"context"
	"fmt"
	"github.com/valyala/fastrand"
	"rump/internal/client"
	"rump/internal/codec/msgpack"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkMultipleUDP(b *testing.B) {
	ctx := context.TODO()
	logger := logging.FromContext(ctx)
	c, err := client.NewUDPClient(ctx, "localhost:5555")
	if err != nil {
		logger.Fatal(err)
	}
	if err := c.Open(); err != nil {
		logger.Fatal(err)
	}
	var cnt uint64
	cd := msgpack.New()
	for i := 0; i < 4; i++ {
		go func() {
			for {
				player := gamestate.GeneratePlayer()
				bytes, err := cd.Encode(player)
				if err != nil {
					logger.Fatal(err)
				}
				if err := c.Write(bytes); err != nil {
					logger.Error("ошибка записи в сокет", err)
				}
				atomic.AddUint64(&cnt, 1)
				if fastrand.Uint32n(100000) == fastrand.Uint32n(100000) {
					fmt.Println(atomic.LoadUint64(&cnt))
				}
			}
		}()
	}
	time.Sleep(30 * time.Second)
}
