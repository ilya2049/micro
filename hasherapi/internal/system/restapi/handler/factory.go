package handler

import (
	"crypto/tls"
	"hasherapi/internal/system/restapi/operations"
	"net/http"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct {
	HashHandler
}

func (h Handler) ConfigureFlags(api *operations.HasherapiAPI) {

}

func (h Handler) ConfigureTLS(tlsConfig *tls.Config) {

}

func (h Handler) ConfigureServer(s *http.Server, scheme, addr string) {

}

func (h Handler) CustomConfigure(api *operations.HasherapiAPI) {

}

func (h Handler) SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func (h Handler) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
