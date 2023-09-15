package servers

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"log/slog"

	"connectrpc.com/connect"
	"github.com/zhongshixi/grpc-go-playground/gen/service"
)

type ConnectEventService struct {
	count atomic.Int64
	debug bool
}

func NewConnectEventService(debug bool) *ConnectEventService {
	return &ConnectEventService{
		debug: debug,
	}
}

func (s *ConnectEventService) HandleRequest(
	ctx context.Context,
	req *connect.Request[service.EventRequest],
) (*connect.Response[service.EventResponse], error) {
	s.count.Add(1)

	if s.debug {
		slog.Info("Event Request Recevied",
			slog.Any("id", req.Msg.OrderId),
			slog.String("method", req.HTTPMethod()),
			slog.Any("peer", req.Peer()),
			slog.Any("spec", req.Spec()),
			slog.Any("headers", req.Header()))
	}

	res := connect.NewResponse(&service.EventResponse{
		OrderId: fmt.Sprintf("Hello, %d!", req.Msg.OrderId),
	})

	return res, nil
}

func (s *ConnectEventService) GetCount() int64 {
	return s.count.Load()
}

func (s *ConnectEventService) HandleClientStream(ctx context.Context, stream *connect.ClientStream[service.EventRequest]) (*connect.Response[service.EventRequestResponse], error) {

	if s.debug {
		slog.Info("Stream Start",
			slog.Any("request_header", stream.RequestHeader()),
			slog.Any("peer", stream.Peer()),
			slog.Any("spec", stream.Spec()),
			slog.Any("err", stream.Err()))
	}

	startTime := time.Now()

	resp := &service.EventRequestResponse{
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
			log.Println("Stream response headers: ", stream.RequestHeader())
		}

		log.Println("Stream Recv", req.OrderId)
	}

	if e := stream.Err(); e != nil {
		return nil, connect.NewError(connect.CodeUnknown, e)
	}

	res := connect.NewResponse(resp)
	return res, nil

}
