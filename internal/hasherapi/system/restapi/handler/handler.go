package handler

import (
	"common/log/logrus"
	"crypto/tls"
	apphash "hasherapi/app/hash"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
	"hasherapi/system/hash/calculator"
	"hasherapi/system/hash/storage"
	"hasherapi/system/restapi/middlewares"
	"hasherapi/system/restapi/operations"
	"net/http"
	"time"
)

func New() *Handler {
	logger := logrus.NewLogger(logrus.Config{
		GraylogHost: "graylog:12201",
		ServiceHost: "hasherapi",
	})

	var hashCalculator hash.Calculator
	hashCalculator = calculator.NewGRPCCalculator("hasher:8090", 1*time.Second, logger)
	hashCalculator = apphash.WrapCalculatorWithLogger(hashCalculator, logger)

	var hashStorage hash.Storage
	hashStorage, closeRedisConnectionsFunc, err := storage.New(storage.Config{
		Address:  "redis:6379",
		Password: "123456789",
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
	}
}

type Handler struct {
	*hashHandler

	closeRedisConnections func()
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
