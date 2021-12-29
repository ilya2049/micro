package log

import "common/log"

const (
	FieldGRPCMetadata = "metadata"
)

const (
	FieldHashInputs     = "hash_inputs"
	FieldHashSHA3Hashes = "sha3_hashes"
)

const (
	FieldRequestID  = log.FieldRequestID
	FieldComponent  = log.FieldComponent
	FieldStackTrace = log.FieldStackTrace
)

const (
	ComponentRequestTracer  = "request_tracer"
	ComponentGRPCAPI        = "grpc_api"
	ComponentHasher         = "hasher"
	ComponentConfigurator   = "configurator"
	ComponentAppInitializer = "app_initializer"
)
