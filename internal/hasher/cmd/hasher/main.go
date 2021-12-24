package main

import (
	"common/hasherproto"
	"fmt"
	"hasher/system/grpcapi"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting server ...")
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	server := grpcapi.NewServer()
	hasherproto.RegisterHasherServiceServer(s, server)

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
