package main

import (
	"context"
	"flag"
	"fmt"
	"rump/internal/client"
	"rump/internal/codec"
	"rump/internal/codec/encbinary"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"sync/atomic"
	"time"
)

func main() {
	addr := flag.String("addr", ":5555", "устанавливает значение addr к которому необходимо подключится")
	flag.Parse()
	ctx := context.Background()
	logger := logging.FromContext(ctx)
	c, err := client.NewUDPClient(ctx, *addr)
	if err != nil {
		logger.Fatal(err)
	}
	if err := c.Open(); err != nil {
		logger.Fatal(err)
	}
	cd := encbinary.New()
	var cnt uint32
	for j := 0; j < 1000; j++ {
		for cnt <= 25 {
			for i := 0; i <= 1000; i++ {
				b, err := encodeCodec(cd)
				if err != nil {
					logger.Fatal(err)
				}
				if err := send(b, c); err != nil {
					logger.Fatal(err)
				}
			}
			atomic.AddUint32(&cnt, 1)
			time.Sleep(40 * time.Millisecond)
		}
		cnt = 0
	}
}

func encodePb() ([]byte, error) {
	return nil, nil
}

func encodeCodec(cd codec.Codec) ([]byte, error) {
	player := gamestate.GeneratePlayer()
	player.TimeStamp = time.Now().UnixNano()
	in := gamestate.Player{
		ID: player.ID,
		Pos: gamestate.Vector{
			X: player.Pos.X,
			Y: player.Pos.Y,
			Z: player.Pos.Z,
		},
		TimeStamp: player.TimeStamp,
	}
	bytes, err := cd.Encode(in)
	if err != nil {
		return nil, fmt.Errorf("client: ошибка декодирования данных %w", err)
	}
	return bytes, nil
}

func send(bytes []byte, c *client.UDPClient) error {
	if err := c.Write(bytes); err != nil {
		return fmt.Errorf("client: ошибка записи в сокет %w", err)
	}
	return nil
}
