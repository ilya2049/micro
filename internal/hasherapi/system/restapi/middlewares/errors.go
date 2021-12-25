package middlewares

import (
	"common/errors"
	"common/requestid"
	"context"
	"hasherapi/app/log"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func NewResponderFactory(logger log.Logger) *ResponderFactory {
	return &ResponderFactory{
		logger: logger,
	}
}

type ResponderFactory struct {
	logger log.Logger
}

func (f *ResponderFactory) NewInternalErrorResponder(
	ctx context.Context,
	responder middleware.Responder,
	err error,
) middleware.Responder {
	return &Responder{
		next:       responder,
		err:        err,
		logger:     f.logger,
		httpStatus: http.StatusInternalServerError,
		ctx:        ctx,
	}
}

func (f *ResponderFactory) NewBadRequestErrorResponder(
	ctx context.Context,
	responder middleware.Responder,
	err error,
) middleware.Responder {
	return &Responder{
		next:       responder,
		err:        err,
		logger:     f.logger,
		httpStatus: http.StatusBadRequest,
		ctx:        ctx,
	}
}

type Responder struct {
	next middleware.Responder

	logger log.Logger

	ctx        context.Context
	err        error
	httpStatus int
}

func (r *Responder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	requestID := requestid.Get(r.ctx)

	errorMessage := r.err.Error()
	logDetails := log.Details{
		log.FieldComponent: log.ComponentHTTPAPI,
		log.FieldRequestID: requestID,
	}

	if stackTrace, ok := errors.StackTrace(r.err); ok {
		logDetails[log.FieldStackTrace] = stackTrace
	}

	switch r.httpStatus {
	case http.StatusBadRequest:
		r.logger.LogWarn(errorMessage, logDetails)

	default:
		r.logger.LogError(errorMessage, logDetails)
	}

	r.next.WriteResponse(w, p)
}
