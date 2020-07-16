package main

import (
	"context"
	"flag"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"rump/internal/client"
	"rump/internal/codec"
	"rump/internal/codec/msgpack"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"rump/internal/pb"
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
	conn, _ := grpc.Dial("127.0.0.1:5577", grpc.WithInsecure())
	client := pb.NewSyncStateClient(conn)

	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}
	var cnt uint32
	cd := msgpack.New()
	for j := 0; j < 1000; j++ {
		for cnt <= 25 {
			for i := 0; i <= 1000; i++ {
				emitFn(cd, client, logger, c)
			}
			atomic.AddUint32(&cnt, 1)
			time.Sleep(40 * time.Millisecond)
		}
		cnt = 0
	}
}

type emitter func(logging.Logger, *client.UDPClient)

func emitFn(codec codec.Codec, client pb.SyncStateClient, logger logging.Logger, c *client.UDPClient) {
	player := gamestate.GeneratePlayer()
	player.TimeStamp = time.Now()

	in := &pb.SyncPos{
		ID: player.ID,
		Pos: &pb.Vector3{
			X: player.Pos.X,
			Y: player.Pos.Y,
			Z: player.Pos.Z,
		},
		Timestamp: time.Now().UnixNano(),
	}
	bytes, err := proto.Marshal(in)
	if err != nil {
		logger.Fatal(err)
	}
	if err := c.Write(bytes); err != nil {
		logger.Fatal("ошибка записи в сокет", err)
	}
}
