package main

import (
	"context"
	"flag"
	"github.com/valyala/fastrand"
	"google.golang.org/grpc"
	"log"
	"rump/internal/logging"
	"rump/internal/pb"
	"sync"
)

var (
	addr = flag.String("addr", ":5577", "устанавливает значение addr к которому необходимо подключится")
)

func main() {
	ctx := context.TODO()
	logger := logging.FromContext(ctx)

	flag.Parse()

	if *addr == "" {
		log.Fatalf("client: необходимый аргумент --addr не задан")
	}

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
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				resp, err := c.RcvPosition(context.TODO(), &pb.RcvPositionRequest{ID: fastrand.Uint32n(1000)})
				if err != nil {
					logger.Error(err)
				}
				_ = resp
			}
		}()
	}
	wg.Wait()
}
