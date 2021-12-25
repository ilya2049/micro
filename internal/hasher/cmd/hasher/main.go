package main

import (
	"common/errors"
	"common/hasherproto"
	"common/log/logrus"
	"hasher/app/log"
	"hasher/system/grpcapi"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func main() {
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

	aLogger.LogInfo("grpc api server is ready to accept requests", log.NoDetails())

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				aLogger.LogError("grpc api server error: "+err.Error(), log.NoDetails())
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	aLogger.LogInfo("grpc api server is shut down gracefully", log.NoDetails())
}
