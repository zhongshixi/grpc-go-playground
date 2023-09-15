package utils

import (
	"log/slog"
	"net/http"
)

type HandlerInterceptor struct {
	Handler http.Handler
}

func (h *HandlerInterceptor) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.Handler.ServeHTTP(rw, req)
	slog.Info("request intercepted", slog.Any("protocol", req.Proto))
}
