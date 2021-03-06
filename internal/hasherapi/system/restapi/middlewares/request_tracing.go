package middlewares

import (
	"common/requestid"
	"context"
	"net/http"
)

func TraceRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		requestID := r.Header.Get(requestid.Header)
		if requestID == "" {
			ctx = requestid.New(r.Context())
		} else {
			ctx = requestid.Set(r.Context(), requestID)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
