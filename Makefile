install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

init:
	buf mod init

gen:
	buf generate

connect-server-http1:
	go run ./cmd/connect-server-http1/main.go --addr :7001

connect-server-http2:
	go run ./cmd/connect-server-http2/main.go --addr :7002

connect-client-http1:
	go run ./cmd/connect-client-http1/main.go --host  http://localhost:7001 --requests 1000

connect-client-http2:
	go run ./cmd/connect-client-http2/main.go --host http://localhost:7002 --requests 10000	

connect-stream-client:
	go run ./cmd/connect-client-stream/main.go --host  http://localhost:7002 --requests 10000
	

twirp-client-http1:
	go run ./cmd/twirp-client-http1/main.go

twirp-client-http2:
	go run ./cmd/twirp-client-http2/main.go

twirp-server-http1:
	go run ./cmd/twirp-server-http1/main.go --addr :7003

twirp-server-http2:
	go run ./cmd/twirp-server-http2/main.go --addr :7004





# bidistream-client:
# go run ./cmd/stream-client/main.go  