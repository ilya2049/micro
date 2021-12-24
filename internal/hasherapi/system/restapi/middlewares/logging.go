package middlewares

import (
	"common/requestid"
	"hasherapi/app/log"
	"hasherapi/pkg/httputil"
	"net/http"
	"time"
)

func Logging(next http.Handler, logger log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if logger.Level() < log.LevelDebug {
			next.ServeHTTP(w, r)

			return
		}

		start := time.Now()
		path := r.URL.Path
		query := r.URL.RawQuery
		requestID := requestid.Get(r.Context())

		bodyBytes, err := httputil.ScanBody(r)
		if err != nil {
			logger.LogWarn("loggingMiddleware: failed to scan body bytes", log.Details{
				log.FieldRequestID:  requestID,
				log.FieldHTTPQuery:  query,
				log.FieldHTTPMethod: r.Method,
				log.FieldAddress:    r.RemoteAddr,
				log.FieldUserAgent:  r.UserAgent(),
			})
		}

		statusRecorder := httputil.NewStatusRecorder(w)

		next.ServeHTTP(statusRecorder, r)

		logger.LogDebug(path, log.Details{
			log.FieldRequestID:  requestID,
			log.FieldHTTPBody:   string(bodyBytes),
			log.FieldHTTPQuery:  query,
			log.FieldHTTPStatus: statusRecorder.Status(),
			log.FieldHTTPMethod: r.Method,
			log.FieldAddress:    r.RemoteAddr,
			log.FieldUserAgent:  r.UserAgent(),
			log.FieldLatency:    time.Since(start).String(),
		})
	})
}
