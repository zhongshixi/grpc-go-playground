package servers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zhongshixi/grpc-go-playground/gen/proto"
)

type TwirpEventService struct {
}

func NewTwripEventService() *TwirpEventService {
	return &TwirpEventService{}
}

func (t *TwirpEventService) HandleRequest(ctx context.Context, req *proto.EventRequest) (*proto.EventResponse, error) {

	slog.Info("Twrip Request Recevied",
		slog.Any("id", req.OrderId))

	return &proto.EventResponse{
		OrderId: fmt.Sprintf("Hello, %d!", req.OrderId),
	}, nil
}

func (t *TwirpEventService) HandleClientStream(ctx context.Context, req *proto.EventRequest) (*proto.EventRequestResponse, error) {
	return nil, nil
}
