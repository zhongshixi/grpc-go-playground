package main

import (
	"context"

	"log/slog"

	"github.com/zhongshixi/grpc-go-playground/data"
	"github.com/zhongshixi/grpc-go-playground/gen/proto"
	"github.com/zhongshixi/grpc-go-playground/utils"
)

func main() {
	slog.Info("Initialize Twrip Client", slog.Any("endpoint addr", ":7004"))

	client := proto.NewEventServiceProtobufClient(
		"http://localhost:7004",
		utils.NewDefaultHTTP2Client(),
	)

	utils.SpikeWithFunc(1000, func(id int64) (int64, error) {
		_, err := client.HandleRequest(context.Background(), &proto.EventRequest{
			OrderId: id,
			Data:    data.Bid,
		})

		return int64(len([]rune(data.Bid))), err
	})
}
