package utils

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewDefaultHTTP2Server(path string, handler http.Handler, addr string) *http.Server {
	// service := &servers.ConnectEventService{}
	// path, handler := serviceconnect.NewEventServiceHandler(service)
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	h := &HandlerInterceptor{
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	server := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	http2.ConfigureServer(server, &http2.Server{})

	// if err := server.ListenAndServe(); err != nil {
	// 	slog.Error("ListenAndServe", slog.Any("error", err))
	// }

	return server
}

func NewDefaultHTTPServer(path string, handler http.Handler, addr string) *http.Server {

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	h := &HandlerInterceptor{
		Handler: mux,
	}

	server := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	return server
}
