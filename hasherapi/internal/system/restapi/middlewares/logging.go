package middlewares

import (
	"hasherapi/internal/app/log"
	"hasherapi/internal/pkg/httputil"
	"hasherapi/internal/pkg/httputil/requestid"
	"net/http"
	"time"
)

func Logging(next http.Handler, logger log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path
		query := r.URL.RawQuery
		requestID := requestid.Get(r.Context())

		statusRecorder := httputil.NewStatusRecorder(w)

		next.ServeHTTP(statusRecorder, r)

		logger.LogInfo(path, log.Details{
			log.FieldRequestID:  requestID,
			log.FieldHTTPQuery:  query,
			log.FieldHTTPStatus: statusRecorder.Status(),
			log.FieldHTTPMethod: r.Method,
			log.FieldAddress:    r.RemoteAddr,
			log.FieldUserAgent:  r.UserAgent(),
			log.FieldLatency:    time.Since(start).String(),
		})
	})
}
