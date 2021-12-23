package middlewares

import (
	"hasherapi/app/log"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func NewResponderFactory(logger log.Logger) *ResponserFactory {
	return &ResponserFactory{
		logger: logger,
	}
}

type ResponserFactory struct {
	logger log.Logger
}

func (f *ResponserFactory) NewInternalErrorResponder(responder middleware.Responder, err error) middleware.Responder {
	return newResponder(responder, err, f.logger, http.StatusInternalServerError)
}

func (f *ResponserFactory) NewBadRequestErrorResponder(responder middleware.Responder, err error) middleware.Responder {
	return newResponder(responder, err, f.logger, http.StatusBadRequest)
}

func newResponder(
	responder middleware.Responder,
	err error,
	logger log.Logger,
	httpStatus int,
) *Responder {
	return &Responder{
		next:       responder,
		err:        err,
		logger:     logger,
		httpStatus: httpStatus,
	}
}

type Responder struct {
	next middleware.Responder

	logger log.Logger

	err        error
	httpStatus int
}

const errorMessage = "rest api error"

func (r *Responder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	switch r.httpStatus {
	case http.StatusBadRequest:
		r.logger.LogWarn(errorMessage, log.Details{
			log.FieldError: r.err.Error(),
		})

	default:
		r.logger.LogError(errorMessage, log.Details{
			log.FieldError: r.err.Error(),
		})
	}

	r.next.WriteResponse(w, p)
}
