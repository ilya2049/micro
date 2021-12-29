package main

import (
	"common/cleanup"
	commonConfig "common/config"
	"common/errors"
	"common/hasherproto"
	"common/log/logrus"
	"hasher/app/log"
	"hasher/system/config"
	"hasher/system/grpcapi"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func main() {
	configProvider, err := config.NewProvider()
	if err != nil {
		stdlog.Fatalf("%s: failed to create a configurator: %s", log.ComponentAppInitializer, err.Error())
	}

	logger, updateLogLevel := logrus.NewLogger(logrus.Config{
		GraylogHost:   configProvider.Logger().Graylog.Host,
		GraylogSource: configProvider.Logger().Graylog.Source,
		LogLevel:      log.Level(configProvider.Logger().Level),
	})

	cleanupFuncs := cleanup.Funcs{}

	stopConfigWatching := config.Watch(configProvider, logger, []commonConfig.Trigger{
		commonConfig.Trigger(func() {
			updateLogLevel(log.Level(configProvider.Logger().Level))
		}),
	})

	cleanupFuncs = append(cleanupFuncs, stopConfigWatching)

	listener, err := net.Listen("tcp", configProvider.GRPC().Host)
	if err != nil {
		logger.LogFatal("failed to start server: "+err.Error(), log.Details{
			log.FieldComponent: log.ComponentGRPCAPI,
		})
	}

	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpcapi.InterceptorTraceRequest(logger),
		grpcapi.InterceptorLogRequest(logger),
	))

	grpcAPIServer := grpcapi.NewServer(logger)
	hasherproto.RegisterHasherServiceServer(grpcServer, grpcAPIServer)

	logger.LogInfo("server is ready to accept requests", log.Details{
		log.FieldComponent: log.ComponentGRPCAPI,
	})

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.LogFatal("server error: "+err.Error(), log.Details{
					log.FieldComponent: log.ComponentGRPCAPI,
				})
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	grpcServer.GracefulStop()
	cleanupFuncs.Invoke()

	logger.LogInfo(log.ComponentGRPCAPI+" is shut down gracefully", log.NoDetails())
}
