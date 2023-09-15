package utils

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewDefaultHTTP2Server(path string, handler http.Handler, addr string) *http.Server {
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
