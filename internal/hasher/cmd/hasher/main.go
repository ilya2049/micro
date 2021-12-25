package main

import (
	"common/hasherproto"
	"common/log/logrus"
	"fmt"
	"hasher/system/grpcapi"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting server ...")
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	aLogger := logrus.NewLogger(logrus.Config{
		GraylogHost: "graylog:12201",
		ServiceHost: "hasher",
	})

	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpcapi.InterceptorTraceRequest(aLogger),
		grpcapi.InterceptorLogRequest(aLogger),
	))

	grpcAPIServer := grpcapi.NewServer(aLogger)
	hasherproto.RegisterHasherServiceServer(grpcServer, grpcAPIServer)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
