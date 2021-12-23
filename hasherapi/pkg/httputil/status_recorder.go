package httputil

import "net/http"

func NewStatusRecorder(responseWriter http.ResponseWriter) *StatusRecorder {
	return &StatusRecorder{
		ResponseWriter: responseWriter,
		status:         http.StatusOK,
	}
}

type StatusRecorder struct {
	http.ResponseWriter

	status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *StatusRecorder) Status() int {
	return r.status
}
