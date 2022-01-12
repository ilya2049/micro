package log

import "common/log"

const (
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
)

const (
	FieldRequestID  = log.FieldRequestID
	FieldComponent  = log.FieldComponent
	FieldStackTrace = log.FieldStackTrace
)

const (
	ComponentHashCalculator = "hash_calculator"
	ComponentHashStorage    = "hash_storage"
	ComponentEventStream    = "event_stream"
	ComponentHTTPAPI        = "http_api"
	ComponentAppInitializer = log.ComponentAppInitializer
	ComponentConfigurator   = log.ComponentConfigurator
)
