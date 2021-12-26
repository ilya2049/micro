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
	"runtime"
	"time"
)

func New() *Handler {
	aLogger := logrus.NewLogger(logrus.Config{
		GraylogHost: "graylog:12201",
		ServiceHost: "hasherapi",
	})

	var hashCalculator hash.Calculator
	hashCalculator = calculator.NewGRPCCalculator("hasher:8090", 1*time.Second, aLogger)
	hashCalculator = apphash.WrapCalculatorWithLogger(hashCalculator, aLogger)

	var hashStorage hash.Storage
	hashStorage, err := storage.New(storage.Config{
		Address:  "redis:6379",
		Password: "123456789",
	}, aLogger)

	if err != nil {
		aLogger.LogError("failed to create a hash storage: "+err.Error(), log.NoDetails())

		runtime.Goexit()
	}

	hashStorage = apphash.WrapStorageWithLogger(hashStorage, aLogger)
	hashService := hash.NewService(hashCalculator, hashStorage)

	errorResponderFactory := middlewares.NewResponderFactory(aLogger)

	return &Handler{
		hashHandler: newHashHandler(hashService, errorResponderFactory),
		aLogger:     aLogger,
	}
}

type Handler struct {
	*hashHandler

	aLogger log.Logger
}

func (h *Handler) ConfigureFlags(api *operations.HasherapiAPI) {}

func (h *Handler) ConfigureTLS(tlsConfig *tls.Config) {}

func (h *Handler) ConfigureServer(s *http.Server, scheme, addr string) {}

func (h *Handler) CustomConfigure(api *operations.HasherapiAPI) {
	api.Logger = h.aLogger.Printf
}

func (h *Handler) SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func (h *Handler) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	handler = middlewares.Logging(handler, h.aLogger)
	handler = middlewares.TraceRequest(handler)

	return handler
}
