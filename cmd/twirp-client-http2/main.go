package main

import (
	"context"

	"log/slog"

	"github.com/zhongshixi/grpc-go-playground/data"
	"github.com/zhongshixi/grpc-go-playground/gen/service"
	"github.com/zhongshixi/grpc-go-playground/utils"
)

func main() {
	slog.Info("Initialize Twrip Client", slog.Any("endpoint addr", ":7004"))

	client := service.NewEventServiceProtobufClient(
		"http://localhost:7004",
		utils.NewDefaultHTTP2Client(),
	)

	utils.SpikeWithFunc(1000, func(id int64) (int64, error) {
		_, err := client.HandleRequest(context.Background(), &service.EventRequest{
			OrderId: id,
			Data:    data.Bid,
		})

		return int64(len([]rune(data.Bid))), err
	})
}
