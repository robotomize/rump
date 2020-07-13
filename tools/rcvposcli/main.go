package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"rump/internal/pb"
	"time"
)

var (
	addr = flag.String("addr", ":5577", "устанавливает значение addr к которому необходимо подключится")
	id   = flag.Uint64("id", 0, "устанавливает id игрока")
	x    = flag.Float64("x", 0, "устанавливает координату x")
	y    = flag.Float64("y", 0, "устанавливает координату y")
	z    = flag.Float64("z", 0, "устанавливает координату z")
)

func main() {
	var (
		player *gamestate.Player
	)
	ctx := context.TODO()
	logger := logging.FromContext(ctx)

	flag.Parse()

	if *addr == "" {
		log.Fatalf("client: необходимый аргумент --addr не задан")
	}
	if *id == 0 {
		logger.Infof("id игрока не передан, генерируем нового %s", *addr)
		player = gamestate.GeneratePlayer()
	} else {
		player = gamestate.NewPlayer(uint32(*id), gamestate.NewVector(*x, *y, *z))
	}

	logger.Infof("отправляю позицию игрока id: %d, позиция x: %f, y: %f, z: %f", player.ID, player.Pos.X, player.Pos.Y, player.Pos.Z)
	player.TimeStamp = time.Now()

	logger.Infof("подключаюсь к %s", *addr)
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("ошибка соединения с grpc: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Error(err)
		}
	}()

	c := pb.NewSyncStateClient(conn)

	resp, err := c.RcvPosition(context.TODO(), &pb.RcvPositionRequest{ID: uint32(*id)})
	if err != nil {
		logger.Error(err)
	}
	fmt.Println(resp)
}
