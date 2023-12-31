package main

import (
	"flag"
	"log/slog"

	"github.com/zhongshixi/grpc-go-playground/gen/proto/protoconnect"
	"github.com/zhongshixi/grpc-go-playground/servers"
	"github.com/zhongshixi/grpc-go-playground/utils"
)

func main() {
	addr := flag.String("addr", ":7002", "host")
	flag.Parse()

	slog.Info("Initialize Connect HTTP 2 Server", slog.Any("Addr", *addr))

	service := servers.NewConnectEventService()
	path, handler := protoconnect.NewEventServiceHandler(service)
	server := utils.NewDefaultHTTP2Server(path, handler, *addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("ListenAndServe", slog.Any("error", err))
	}
}
