package log

const (
	FieldRequestID  = "request_id"
	FieldHTTPQuery  = "query"
	FieldHTTPBody   = "body"
	FieldHTTPStatus = "status"
	FieldHTTPMethod = "method"
	FieldAddress    = "address"
	FieldUserAgent  = "ua"
	FieldLatency    = "latency"
)

const (
	FieldHashInputs               = "hash_inputs"
	FieldHashSHA3Hashes           = "sha3_hashes"
	FieldHashIdentifiedSHA3Hashes = "identified_sha3_hashes"
	FieldHashIDs                  = "hash_ids"
	FieldComponent                = "component"
	FieldStackTrace               = "stack_trace"
)

const (
	ComponentHashCalculator = "hash_calculator"
	ComponentHashStorage    = "hash_storage"
	ComponentHTTPAPI        = "http_api"
)
