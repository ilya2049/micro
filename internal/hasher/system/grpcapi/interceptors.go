package grpcapi

import (
	"common/requestid"
	"context"

	"hasher/app/log"
	"hasher/pkg/grpcutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func InterceptorTraceRequest(logger log.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		grpcMetadata, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			ctx, requestID := requestid.NewGet(ctx)
			logger.LogWarn(
				"no metadata (a new request id generated)",
				log.Details{
					log.FieldRequestID: requestID,
					log.FieldComponent: log.ComponentRequestTracer,
				})

			return handler(ctx, req)
		}

		requestID := grpcutil.MetadataGetFirst("X-Request-ID", grpcMetadata)

		if requestID == "" {
			ctx, requestID = requestid.NewGet(ctx)
			logger.LogWarn(
				"no request id in metadata (a new request id generated)",
				log.Details{
					log.FieldRequestID: requestID,
					log.FieldComponent: log.ComponentRequestTracer,
				})
		} else {
			ctx = requestid.Set(ctx, requestID)
		}

		return handler(ctx, req)
	}
}

func InterceptorLogRequest(logger log.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		requestID := requestid.Get(ctx)

		logger.LogDebug(info.FullMethod, log.Details{
			log.FieldRequestID: requestID,
			log.FieldComponent: log.ComponentGRPCAPI,
		})

		return handler(ctx, req)
	}
}
