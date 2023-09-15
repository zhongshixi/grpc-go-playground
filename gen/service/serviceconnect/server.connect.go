// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: service/server.proto

package serviceconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	service "github.com/zhongshixi/grpc-go-playground/gen/service"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// EventServiceName is the fully-qualified name of the EventService service.
	EventServiceName = "service.EventService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EventServiceHandleRequestProcedure is the fully-qualified name of the EventService's
	// HandleRequest RPC.
	EventServiceHandleRequestProcedure = "/service.EventService/HandleRequest"
	// EventServiceHandleClientStreamProcedure is the fully-qualified name of the EventService's
	// HandleClientStream RPC.
	EventServiceHandleClientStreamProcedure = "/service.EventService/HandleClientStream"
)

// EventServiceClient is a client for the service.EventService service.
type EventServiceClient interface {
	HandleRequest(context.Context, *connect.Request[service.EventRequest]) (*connect.Response[service.EventResponse], error)
	HandleClientStream(context.Context) *connect.ClientStreamForClient[service.EventRequest, service.EventRequestResponse]
}

// NewEventServiceClient constructs a client for the service.EventService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEventServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) EventServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &eventServiceClient{
		handleRequest: connect.NewClient[service.EventRequest, service.EventResponse](
			httpClient,
			baseURL+EventServiceHandleRequestProcedure,
			opts...,
		),
		handleClientStream: connect.NewClient[service.EventRequest, service.EventRequestResponse](
			httpClient,
			baseURL+EventServiceHandleClientStreamProcedure,
			opts...,
		),
	}
}

// eventServiceClient implements EventServiceClient.
type eventServiceClient struct {
	handleRequest      *connect.Client[service.EventRequest, service.EventResponse]
	handleClientStream *connect.Client[service.EventRequest, service.EventRequestResponse]
}

// HandleRequest calls service.EventService.HandleRequest.
func (c *eventServiceClient) HandleRequest(ctx context.Context, req *connect.Request[service.EventRequest]) (*connect.Response[service.EventResponse], error) {
	return c.handleRequest.CallUnary(ctx, req)
}

// HandleClientStream calls service.EventService.HandleClientStream.
func (c *eventServiceClient) HandleClientStream(ctx context.Context) *connect.ClientStreamForClient[service.EventRequest, service.EventRequestResponse] {
	return c.handleClientStream.CallClientStream(ctx)
}

// EventServiceHandler is an implementation of the service.EventService service.
type EventServiceHandler interface {
	HandleRequest(context.Context, *connect.Request[service.EventRequest]) (*connect.Response[service.EventResponse], error)
	HandleClientStream(context.Context, *connect.ClientStream[service.EventRequest]) (*connect.Response[service.EventRequestResponse], error)
}

// NewEventServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEventServiceHandler(svc EventServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	eventServiceHandleRequestHandler := connect.NewUnaryHandler(
		EventServiceHandleRequestProcedure,
		svc.HandleRequest,
		opts...,
	)
	eventServiceHandleClientStreamHandler := connect.NewClientStreamHandler(
		EventServiceHandleClientStreamProcedure,
		svc.HandleClientStream,
		opts...,
	)
	return "/service.EventService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EventServiceHandleRequestProcedure:
			eventServiceHandleRequestHandler.ServeHTTP(w, r)
		case EventServiceHandleClientStreamProcedure:
			eventServiceHandleClientStreamHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedEventServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEventServiceHandler struct{}

func (UnimplementedEventServiceHandler) HandleRequest(context.Context, *connect.Request[service.EventRequest]) (*connect.Response[service.EventResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("service.EventService.HandleRequest is not implemented"))
}

func (UnimplementedEventServiceHandler) HandleClientStream(context.Context, *connect.ClientStream[service.EventRequest]) (*connect.Response[service.EventRequestResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("service.EventService.HandleClientStream is not implemented"))
}
