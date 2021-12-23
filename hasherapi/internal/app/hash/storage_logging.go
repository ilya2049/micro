package hash

import (
	"context"
	"hasherapi/internal/app/log"
	"hasherapi/internal/domain/hash"
	"hasherapi/internal/pkg/httputil/requestid"
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

const (
	messageHashStorageSave = "hashStorage.Save"
	messageHashStorageGet  = "hashStorage.Get"
)

func (s *storageLoggingWrapper) Save(ctx context.Context, sha3Hashes hash.SHA3Hashes) ([]hash.IdentifiedSHA3Hash, error) {
	requestID := requestid.Get(ctx)

	s.logger.LogDebug(messageHashStorageSave, log.Details{
		log.FieldRequestID:      requestID,
		log.FieldHashSHA3Hashes: sha3Hashes,
	})

	identifiesSHA3Hashes, err := s.next.Save(ctx, sha3Hashes)

	if err == nil {
		s.logger.LogDebug(messageHashStorageSave, log.Details{
			log.FieldRequestID:                requestID,
			log.FieldHashIdentifiedSHA3Hashes: identifiesSHA3Hashes,
		})
	}

	return identifiesSHA3Hashes, err
}

func (s *storageLoggingWrapper) Get(ctx context.Context, hashIDs []hash.ID) ([]hash.IdentifiedSHA3Hash, error) {
	requestID := requestid.Get(ctx)

	s.logger.LogDebug(messageHashStorageGet, log.Details{
		log.FieldRequestID: requestID,
		log.FieldHashIDs:   hashIDs,
	})

	identifiesSHA3Hashes, err := s.next.Get(ctx, hashIDs)

	if err == nil {
		s.logger.LogDebug(messageHashStorageGet, log.Details{
			log.FieldRequestID:                requestID,
			log.FieldHashIdentifiedSHA3Hashes: identifiesSHA3Hashes,
		})
	}

	return identifiesSHA3Hashes, err
}
