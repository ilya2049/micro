package hash

import (
	"common/requestid"
	"context"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
)

func WrapStorageWithLogger(storage hash.Storage, logger log.Logger) hash.Storage {
	return &storageLoggingWrapper{
		next:   storage,
		logger: logger,
	}
}

type storageLoggingWrapper struct {
	next hash.Storage

	logger log.Logger
}

func (s *storageLoggingWrapper) Save(ctx context.Context, sha3Hashes hash.SHA3Hashes) ([]hash.IdentifiedSHA3Hash, error) {
	identifiesSHA3Hashes, err := s.next.Save(ctx, sha3Hashes)

	requestID := requestid.Get(ctx)

	s.logger.LogDebug("save", log.Details{
		log.FieldRequestID:                requestID,
		log.FieldHashSHA3Hashes:           sha3Hashes,
		log.FieldComponent:                log.ComponentHashStorage,
		log.FieldHashIdentifiedSHA3Hashes: identifiesSHA3Hashes,
	})

	return identifiesSHA3Hashes, err
}

func (s *storageLoggingWrapper) Get(ctx context.Context, hashIDs []hash.ID) ([]hash.IdentifiedSHA3Hash, error) {
	identifiesSHA3Hashes, err := s.next.Get(ctx, hashIDs)

	requestID := requestid.Get(ctx)

	s.logger.LogDebug("get", log.Details{
		log.FieldRequestID:                requestID,
		log.FieldHashIDs:                  hashIDs,
		log.FieldComponent:                log.ComponentHashStorage,
		log.FieldHashIdentifiedSHA3Hashes: identifiesSHA3Hashes,
	})

	return identifiesSHA3Hashes, err
}
