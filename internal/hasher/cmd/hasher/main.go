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
	aLogger := logrus.NewLogger(logrus.Config{
		GraylogHost:   "graylog:12201",
		GraylogSource: "hasher",
	})

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		aLogger.LogFatal("failed to start server: "+err.Error(), log.Details{
			log.FieldComponent: log.ComponentGRPCAPI,
		})
	}

	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpcapi.InterceptorTraceRequest(aLogger),
		grpcapi.InterceptorLogRequest(aLogger),
	))

	grpcAPIServer := grpcapi.NewServer(aLogger)
	hasherproto.RegisterHasherServiceServer(grpcServer, grpcAPIServer)

	aLogger.LogInfo("server is ready to accept requests", log.Details{
		log.FieldComponent: log.ComponentGRPCAPI,
	})

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				aLogger.LogFatal("server error: "+err.Error(), log.Details{
					log.FieldComponent: log.ComponentGRPCAPI,
				})
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	grpcServer.GracefulStop()

	aLogger.LogInfo(log.ComponentGRPCAPI+" is shut down gracefully", log.NoDetails())
}
