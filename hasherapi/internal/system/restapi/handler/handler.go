package handler

import (
	"crypto/tls"
	"hasherapi/internal/domain/hash"
	"hasherapi/internal/system/hash/fakecalculator"
	"hasherapi/internal/system/hash/fakestorage"
	"hasherapi/internal/system/restapi/operations"
	"net/http"
)

func New() *Handler {
	hashCalculator := fakecalculator.New()
	hashStorage := fakestorage.New()
	hashService := hash.NewService(hashCalculator, hashStorage)

	return &Handler{
		hashHandler: newHashHandler(hashService),
	}
}

type Handler struct {
	*hashHandler
}

func (h *Handler) ConfigureFlags(api *operations.HasherapiAPI) {

}

func (h *Handler) ConfigureTLS(tlsConfig *tls.Config) {

}

func (h *Handler) ConfigureServer(s *http.Server, scheme, addr string) {

}

func (h *Handler) CustomConfigure(api *operations.HasherapiAPI) {

}

func (h *Handler) SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func (h *Handler) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
