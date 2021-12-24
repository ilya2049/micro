package hash_test

import (
	"context"
	"hasherapi/domain/hash"
	inmemoryCalculator "hasherapi/system/hash/calculator/inmemory"
	inmemoryStorage "hasherapi/system/hash/storage/inmemory"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_CreateHashes(t *testing.T) {
	storage := inmemoryStorage.NewHashStorage()
	calculator := inmemoryCalculator.NewHashCalculator()

	service := hash.NewService(calculator, storage)

	identifiedSHA3Hashes, err := service.CreateHashes(
		context.Background(),
		hash.NewInputs([]string{"1", "2", "3"}),
	)

	require.NoError(t, err)

	wantIdentifiedSHA3Hashes := []hash.IdentifiedSHA3Hash{
		hash.NewIdentifiedSHA3Hash(hash.ID(1), hash.SHA3Hash("hash-of-1")),
		hash.NewIdentifiedSHA3Hash(hash.ID(2), hash.SHA3Hash("hash-of-2")),
		hash.NewIdentifiedSHA3Hash(hash.ID(3), hash.SHA3Hash("hash-of-3")),
	}

	assert.Equal(t, wantIdentifiedSHA3Hashes, identifiedSHA3Hashes)
}

func TestService_FindHashes(t *testing.T) {
	storage := inmemoryStorage.NewHashStorage()
	calculator := inmemoryCalculator.NewHashCalculator()

	service := hash.NewService(calculator, storage)
	ctx := context.Background()

	_, err := service.CreateHashes(ctx, hash.NewInputs([]string{"1", "2", "3"}))

	require.NoError(t, err)

	wantIdentifiedSHA3Hashes := []hash.IdentifiedSHA3Hash{
		hash.NewIdentifiedSHA3Hash(hash.ID(1), hash.SHA3Hash("hash-of-1")),
		hash.NewIdentifiedSHA3Hash(hash.ID(2), hash.SHA3Hash("hash-of-2")),
	}

	identifiedSHA3Hashes, err := service.FindHashes(ctx, []hash.ID{hash.ID(1), hash.ID(2), hash.ID(10)})
	require.NoError(t, err)

	assert.Equal(t, wantIdentifiedSHA3Hashes, identifiedSHA3Hashes)
}
