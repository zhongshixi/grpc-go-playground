package servers

import (
	"context"
	"fmt"
	"time"

	"log/slog"

	"connectrpc.com/connect"
	"github.com/zhongshixi/grpc-go-playground/gen/proto"
)

type ConnectEventService struct {
}

func NewConnectEventService() *ConnectEventService {
	return &ConnectEventService{}
}

func (s *ConnectEventService) HandleRequest(
	ctx context.Context,
	req *connect.Request[proto.EventRequest],
) (*connect.Response[proto.EventResponse], error) {

	slog.Info("Event Request Recevied",
		slog.Any("id", req.Msg.OrderId),
		slog.String("method", req.HTTPMethod()),
		slog.Any("peer", req.Peer()),
		slog.Any("spec", req.Spec()),
		slog.Any("headers", req.Header()))

	res := connect.NewResponse(&proto.EventResponse{
		OrderId: fmt.Sprintf("Hello, %d!", req.Msg.OrderId),
	})

	return res, nil
}

func (s *ConnectEventService) HandleClientStream(ctx context.Context, stream *connect.ClientStream[proto.EventRequest]) (*connect.Response[proto.EventRequestResponse], error) {

	slog.Info("Stream Start",
		slog.Any("request_header", stream.RequestHeader()),
		slog.Any("peer", stream.Peer()),
		slog.Any("spec", stream.Spec()),
		slog.Any("err", stream.Err()))

	startTime := time.Now()

	resp := &proto.EventRequestResponse{
		Count:     0,
		TotalSize: 0,
	}

	defer func() {
		slog.Info("time elaspsed:", time.Since(startTime).Seconds(), "seconds")
	}()

	count := int64(0)

	for stream.Receive() {
		count = count + 1
		req := stream.Msg()
		resp.Count++
		resp.TotalSize = resp.TotalSize + int64(len(req.Data))
		if count == 1 {
			slog.Info("Stream Request", slog.Any("header", stream.RequestHeader()))
		}

		slog.Info("Stream Received", slog.Any("request", req))
	}

	if e := stream.Err(); e != nil {
		return nil, connect.NewError(connect.CodeUnknown, e)
	}

	res := connect.NewResponse(resp)
	return res, nil

}
