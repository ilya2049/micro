package handler

import (
	"common/log/logrus"
	"crypto/tls"
	apphash "hasherapi/app/hash"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
	"hasherapi/system/config"
	"hasherapi/system/hash/calculator"
	"hasherapi/system/hash/storage"
	"hasherapi/system/restapi/middlewares"
	"hasherapi/system/restapi/operations"
	stdlog "log"
	"net/http"
	"time"
)

func New() *Handler {
	var logger log.Logger

	configProvider, stopConfigWatching, err := config.NewProvider(logger)
	if err != nil {
		stdlog.Fatal("failed to create a configurator: " + err.Error())
	}

	logger = logrus.NewLogger(logrus.Config{
		GraylogHost:   configProvider.Graylog().Host,
		GraylogSource: configProvider.Graylog().Source,
	})

	var hashCalculator hash.Calculator
	hashCalculator = calculator.NewGRPCCalculator(
		configProvider.Hasher().Host,
		time.Duration(configProvider.Hasher().TimeoutSec)*time.Second,
		logger,
	)
	hashCalculator = apphash.WrapCalculatorWithLogger(hashCalculator, logger)

	var hashStorage hash.Storage
	hashStorage, closeRedisConnectionsFunc, err := storage.New(storage.Config{
		Address:  configProvider.Redis().Host,
		Password: configProvider.Redis().Password,
	}, logger)

	if err != nil {
		logger.LogFatal("failed to create a hash storage: "+err.Error(), log.Details{
			log.FieldComponent: log.ComponentAppInitializer,
		})
	}

	hashStorage = apphash.WrapStorageWithLogger(hashStorage, logger)
	hashService := hash.NewService(hashCalculator, hashStorage)

	errorResponderFactory := middlewares.NewResponderFactory(logger)

	return &Handler{
		hashHandler:           newHashHandler(hashService, errorResponderFactory),
		logger:                logger,
		closeRedisConnections: closeRedisConnectionsFunc,
		stopConfigWatching:    stopConfigWatching,
	}
}

type Handler struct {
	*hashHandler

	closeRedisConnections func()
	stopConfigWatching    func()
	logger                log.Logger
}

func (h *Handler) ConfigureFlags(api *operations.HasherapiAPI) {}

func (h *Handler) ConfigureTLS(tlsConfig *tls.Config) {}

func (h *Handler) ConfigureServer(s *http.Server, scheme, addr string) {}

func (h *Handler) CustomConfigure(api *operations.HasherapiAPI) {
	api.Logger = h.logger.Printf

	api.ServerShutdown = func() {
		h.closeRedisConnections()
	}
}

func (h *Handler) SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func (h *Handler) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	handler = middlewares.Logging(handler, h.logger)
	handler = middlewares.TraceRequest(handler)

	return handler
}
