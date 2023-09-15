package servers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zhongshixi/grpc-go-playground/gen/service"
)

type TwirpEventService struct {
	debug bool
}

func NewTwripEventService(debug bool) *TwirpEventService {
	return &TwirpEventService{
		debug: debug,
	}
}

func (t *TwirpEventService) HandleRequest(ctx context.Context, req *service.EventRequest) (*service.EventResponse, error) {
	if t.debug {
		slog.Info("Twrip Request Recevied",
			slog.Any("id", req.OrderId))
	}

	return &service.EventResponse{
		OrderId: fmt.Sprintf("Hello, %d!", req.OrderId),
	}, nil
}

func (t *TwirpEventService) HandleClientStream(ctx context.Context, req *service.EventRequest) (*service.EventRequestResponse, error) {
	return nil, nil
}
