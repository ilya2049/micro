package requestid

import (
	"context"

	guid "github.com/satori/go.uuid"
)

type key string

const (
	contextKey key = "__ctx_requestId"
)

func Set(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, contextKey, requestID)
}

func New(ctx context.Context) context.Context {
	uuid := guid.Must(guid.NewV4(), nil)
	requestID := uuid.String()

	return Set(ctx, requestID)
}

func Get(ctx context.Context) string {
	if requestIDInContext := ctx.Value(contextKey); requestIDInContext != nil {
		if requestID, ok := requestIDInContext.(string); ok {
			return requestID
		}
	}

	return ""
}
