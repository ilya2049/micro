package log

const (
	FieldRequestID    = "request_id"
	FieldGRPCMetadata = "metadata"
	FieldComponent    = "component"
	FieldStackTrace   = "stack_trace"
)

const (
	FieldHashInputs     = "hash_inputs"
	FieldHashSHA3Hashes = "sha3_hashes"
)

const (
	ComponentRequestTracer = "request_tracer"
	ComponentGRPCAPI       = "grpc_api"
	ComponentHasher        = "hasher"
)
