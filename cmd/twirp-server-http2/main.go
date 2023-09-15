package main

import (
	"flag"
	"log/slog"

	"github.com/zhongshixi/grpc-go-playground/gen/proto"
	"github.com/zhongshixi/grpc-go-playground/servers"
	"github.com/zhongshixi/grpc-go-playground/utils"
)

func main() {

	addr := flag.String("addr", ":7001", "host")
	flag.Parse()

	slog.Info("Initialize Twrip HTTP 2 Server", slog.Any("Addr", *addr))
	svc := servers.NewTwripEventService()
	twrpSvc := proto.NewEventServiceServer(svc)
	server := utils.NewDefaultHTTP2Server(twrpSvc.PathPrefix(), twrpSvc, *addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("ListenAndServe", slog.Any("error", err))
	}
}
