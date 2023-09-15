package main

import (
	"context"
	"errors"
	"flag"
	"io"
	"log"
	"sync"

	"log/slog"

	"github.com/zhongshixi/grpc-go-playground/data"
	"github.com/zhongshixi/grpc-go-playground/gen/service"
	"github.com/zhongshixi/grpc-go-playground/gen/service/serviceconnect"
	"github.com/zhongshixi/grpc-go-playground/utils"
)

func main() {

	host := flag.String("host", "http://localhost:7002", "host")
	requests := flag.Int64("requests", 1, "requests")
	flag.Parse()

	slog.Info("Initialize Connect Stream Client", slog.Any("host", *host), slog.Any("request", *requests))

	client := serviceconnect.NewEventServiceClient(
		utils.NewDefaultHTTP2Client(),
		*host, // in production example, it is https
		// connect.WithGRPC(),
		// connect.WithSendCompression("gzip"),
	)

	stream := client.HandleClientStream(context.Background())
	var mutex sync.Mutex

	utils.SpikeWithFunc(*requests, func(id int64) (int64, error) {
		mutex.Lock()
		defer mutex.Unlock()
		err := stream.Send(&service.EventRequest{
			OrderId: id,
			Data:    data.Bid,
		})

		if err != nil {
			log.Println("Err:", err)
			if errors.Is(err, io.EOF) {
				_, err := stream.CloseAndReceive()
				if err != nil {
					log.Println("Receive Err:", err)
				}
			}
			return -1, err
		}

		return int64(len([]rune(data.Bid))), nil
	})
}
