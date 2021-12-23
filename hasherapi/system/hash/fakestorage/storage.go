package fakestorage

import (
	"context"
	"sync"

	"hasherapi/domain/hash"
)

func New() *HashStorage {
	return &HashStorage{}
}

type HashStorage struct {
	mutex  sync.RWMutex
	hashes hash.SHA3Hashes
}

func (s *HashStorage) nextID() hash.ID {
	return s.lastID() + 1
}

func (s *HashStorage) lastID() hash.ID {
	return hash.ID(len(s.hashes))
}

func (s *HashStorage) Save(_ context.Context, sha3Hashes hash.SHA3Hashes) ([]hash.IdentifiedSHA3Hash, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	identifiedHashes := make([]hash.IdentifiedSHA3Hash, 0, len(sha3Hashes))

	for _, sha3Hash := range sha3Hashes {
		identifiedHash := hash.NewIdentifiedSHA3Hash(s.nextID(), sha3Hash)

		s.hashes = append(s.hashes, sha3Hash)
		identifiedHashes = append(identifiedHashes, identifiedHash)
	}

	return identifiedHashes, nil
}

func (s *HashStorage) Get(_ context.Context, hashIDs []hash.ID) ([]hash.IdentifiedSHA3Hash, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	identifiedHashes := make([]hash.IdentifiedSHA3Hash, 0, len(hashIDs))

	for _, hashID := range hashIDs {
		if hashID > s.lastID() || hashID < 0 {
			continue
		}

		identifiedHashes = append(identifiedHashes, hash.NewIdentifiedSHA3Hash(
			hashID,
			s.hashes[hashIDToHashIndex(hashID)]),
		)
	}

	return identifiedHashes, nil
}

func hashIDToHashIndex(hashID hash.ID) int {
	return int(hashID) - 1
}
