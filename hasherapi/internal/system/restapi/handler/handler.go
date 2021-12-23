package handler

import (
	"crypto/tls"
	apphash "hasherapi/internal/app/hash"
	"hasherapi/internal/app/log"
	"hasherapi/internal/domain/hash"
	"hasherapi/internal/system/hash/fakecalculator"
	"hasherapi/internal/system/hash/fakestorage"
	"hasherapi/internal/system/logger"
	"hasherapi/internal/system/restapi/middlewares"
	"hasherapi/internal/system/restapi/operations"
	"net/http"
)

func New() *Handler {
	aLogger := logger.New()

	hashCalculator := apphash.WrapCalculatorWithLogger(fakecalculator.New(), aLogger)
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
