package fakestorage_test

import (
	"context"
	"hasherapi/domain/hash"
	"hasherapi/system/hash/fakestorage"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashStorage_Save(t *testing.T) {
	storage := fakestorage.New()

	sha3Hashes := hash.NewSHA3Hashes([]string{
		"hash-of-1",
		"hash-of-2",
		"hash-of-3",
	})

	identifiedHashes, _ := storage.Save(context.Background(), sha3Hashes)
	const (
		first         = 0
		third         = 2
		unknownHashID = hash.ID(10)
	)

	require.Len(t, identifiedHashes, len(sha3Hashes))

	receivedIdentifiedHashes, _ := storage.Get(context.Background(), []hash.ID{
		identifiedHashes[first].ID(),
		unknownHashID,
		identifiedHashes[third].ID(),
	})

	wantReceivedIdentifiedHashes := []hash.IdentifiedSHA3Hash{
		hash.NewIdentifiedSHA3Hash(identifiedHashes[first].ID(), sha3Hashes[first]),
		hash.NewIdentifiedSHA3Hash(identifiedHashes[third].ID(), sha3Hashes[third]),
	}

	assert.Equal(t, wantReceivedIdentifiedHashes, receivedIdentifiedHashes)
}
