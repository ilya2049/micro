package middlewares

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

type Responser struct {
	next middleware.Responder

	err        error
	httpStatus int
}

func (r *Responser) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	message := r.err.Error()

	switch r.httpStatus {
	case http.StatusBadRequest:
		message = "warn: " + message

	default:
		message = "error: " + message
	}

	log.Println(message)

	r.next.WriteResponse(w, p)
}

func NewInternalErrorResponder(responder middleware.Responder, err error) middleware.Responder {
	return &Responser{
		next:       responder,
		err:        err,
		httpStatus: http.StatusInternalServerError,
	}
}

func NewBadRequestErrorResponder(responder middleware.Responder, err error) middleware.Responder {
	return &Responser{
		next:       responder,
		err:        err,
		httpStatus: http.StatusBadRequest,
	}
}
