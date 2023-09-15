package main

import (
	"flag"
	"log/slog"

	"github.com/zhongshixi/grpc-go-playground/gen/service/serviceconnect"
	"github.com/zhongshixi/grpc-go-playground/servers"
	"github.com/zhongshixi/grpc-go-playground/utils"
)

func main() {
	addr := flag.String("addr", ":7001", "host")
	flag.Parse()

	slog.Info("Initialize Connect HTTP 1 Server", slog.Any("Addr", *addr))

	service := servers.NewConnectEventService(true)
	path, handler := serviceconnect.NewEventServiceHandler(service)
	server := utils.NewDefaultHTTPServer(path, handler, *addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("ListenAndServe", slog.Any("error", err))
	}
}
