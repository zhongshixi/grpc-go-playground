package main

import (
	"context"
	"flag"

	"log/slog"

	"connectrpc.com/connect"
	"github.com/zhongshixi/grpc-go-playground/data"
	"github.com/zhongshixi/grpc-go-playground/gen/proto"
	"github.com/zhongshixi/grpc-go-playground/gen/proto/protoconnect"
	"github.com/zhongshixi/grpc-go-playground/utils"
)

func main() {
	host := flag.String("host", "http://localhost:7100", "host")
	requests := flag.Int64("requests", 1, "requests")
	flag.Parse()

	slog.Info("Initialize Connect Client", slog.Any("host", *host), slog.Any("request", *requests))

	client := protoconnect.NewEventServiceClient(
		utils.NewDefaultHTTP2Client(),
		*host, // in production example, it is https
		// connect.WithGRPC(),
		// connect.WithSendCompression("gzip"),
	)

	function := func(id int64) (int64, error) {
		req := connect.NewRequest(&proto.EventRequest{
			OrderId: id,
			Data:    data.Bid,
		})

		_, err := client.HandleRequest(
			context.Background(),
			req,
		)

		return int64(len([]rune(data.Bid))), err
	}

	utils.SpikeWithFunc(*requests, function)
}
