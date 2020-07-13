package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"rump/internal/client"
	"rump/internal/codec"
	"rump/internal/codec/msgpack"
	"rump/internal/gamestate"
	"rump/internal/logging"
	"time"
)

func main() {
	addr := flag.String("addr", ":5555", "устанавливает значение addr к которому необходимо подключится")
	id := flag.Uint64("id", 0, "устанавливает id игрока")
	x := flag.Float64("x", 0, "устанавливает координату x")
	y := flag.Float64("y", 0, "устанавливает координату y")
	z := flag.Float64("z", 0, "устанавливает координату z")
	flag.Parse()
	var (
		player *gamestate.Player
		cd     codec.Codec
	)
	ctx := context.TODO()
	logger := logging.FromContext(ctx)

	fmt.Println(x, y)
	if *addr == "" {
		log.Fatalf("client: необходимый аргумент --addr не задан")
	}
	if *id == 0 {
		logger.Infof("id игрока не передан, генерируем нового %s", *addr)
		player = gamestate.GeneratePlayer()
	} else {
		player = gamestate.NewPlayer(uint32(*id), gamestate.NewVector(*x, *y, *z))
	}
	c, err := client.NewUDPClient(ctx, *addr)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("подключаюсь к %s", *addr)
	if err := c.Open(); err != nil {
		logger.Fatal(err)
	}
	logger.Infof("отправляю позицию игрока id: %d, позиция x: %f, y: %f, z: %f", player.ID, player.Pos.X, player.Pos.Y, player.Pos.Z)
	cd = msgpack.New()
	player.TimeStamp = time.Now()
	t := time.Now()
	bytes, err := cd.Encode(player)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(time.Since(t))
	if err := c.Write(bytes); err != nil {
		logger.Fatal("ошибка записи в сокет", err)
	}
}
