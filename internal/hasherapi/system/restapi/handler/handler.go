package handler

import (
	"crypto/tls"
	apphash "hasherapi/app/hash"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
	"hasherapi/system/hash/calculator"
	"hasherapi/system/hash/fakestorage"
	"hasherapi/system/logger"
	"hasherapi/system/restapi/middlewares"
	"hasherapi/system/restapi/operations"
	"net/http"
	"time"
)

func New() *Handler {
	aLogger := logger.New()

	var hashCalculator hash.Calculator
	hashCalculator = calculator.NewGRPCCalculator("hasher:8090", 1*time.Second, aLogger)
	hashCalculator = apphash.WrapCalculatorWithLogger(hashCalculator, aLogger)

	hashStorage := apphash.WrapStorageWithLogger(fakestorage.New(), aLogger)
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
