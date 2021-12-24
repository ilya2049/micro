package hash

import (
	"common/requestid"
	"context"
	"hasher/app/log"
	"hasher/domain/hash"
)

func CalculateSHA3HashSum(ctx context.Context, logger log.Logger) func([]hash.Input) hash.SHA3Hashes {
	return func(inputs []hash.Input) hash.SHA3Hashes {
		sha3Hashes := hash.CalculateSHA3HashSum(inputs)

		requestID := requestid.Get(ctx)

		logger.LogDebug("calculate_sha3_hash_sum", log.Details{
			log.FieldComponent:      log.ComponentHasher,
			log.FieldRequestID:      requestID,
			log.FieldHashInputs:     inputs,
			log.FieldHashSHA3Hashes: sha3Hashes,
		})

		return sha3Hashes
	}
}
