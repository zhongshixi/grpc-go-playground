package main

import (
	"flag"
	"log/slog"

	"github.com/zhongshixi/grpc-go-playground/gen/service"
	"github.com/zhongshixi/grpc-go-playground/servers"
	"github.com/zhongshixi/grpc-go-playground/utils"
)

func main() {

	addr := flag.String("addr", ":7003", "addr")
	flag.Parse()

	slog.Info("Initialize Twrip  HTTP 1 Server", slog.Any("Addr", *addr))

	svc := servers.NewTwripEventService(true)
	twrpSvc := service.NewEventServiceServer(svc)
	server := utils.NewDefaultHTTPServer(twrpSvc.PathPrefix(), twrpSvc, *addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("ListenAndServe", slog.Any("error", err))
	}
}
