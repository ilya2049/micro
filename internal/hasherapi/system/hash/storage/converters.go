package storage

import (
	"fmt"
	"hasherapi/domain/hash"
)

func sha3HashesToEmptyInterfaces(sha3Hashes hash.SHA3Hashes) []interface{} {
	emptyInterfaces := make([]interface{}, 0, len(sha3Hashes))

	for _, sha3Hash := range sha3Hashes {
		emptyInterfaces = append(emptyInterfaces, string(sha3Hash))
	}

	return emptyInterfaces
}

func sha3HashesToIdentifiedSHA3Hashes(sha3Hashes hash.SHA3Hashes, sha3HashIDs []int64) []hash.IdentifiedSHA3Hash {
	if len(sha3Hashes) != len(sha3HashIDs) {
		return []hash.IdentifiedSHA3Hash{}
	}

	identifiedHashes := make([]hash.IdentifiedSHA3Hash, 0, len(sha3Hashes))

	for i, sha3Hash := range sha3Hashes {
		identifiedHashes = append(identifiedHashes, hash.NewIdentifiedSHA3Hash(
			hash.ID(sha3HashIDs[i]),
			hash.SHA3Hash(sha3Hash),
		))
	}

	return identifiedHashes
}

func hashIDsToRedisKeys(hashIDs []hash.ID) []string {
	redisKeys := make([]string, 0, len(hashIDs))

	for _, hashID := range hashIDs {
		redisKeys = append(redisKeys, fmt.Sprintf("sha3:%d", hashID))
	}

	return redisKeys
}

func emptyInterfacesToIdentifiedHashes(emptyInterfaces []interface{}, hashIDs []hash.ID) []hash.IdentifiedSHA3Hash {
	if len(emptyInterfaces) != len(hashIDs) {
		return []hash.IdentifiedSHA3Hash{}
	}

	identifiedSHA3Hashes := make([]hash.IdentifiedSHA3Hash, 0, len(emptyInterfaces))

	for i, emptyInterface := range emptyInterfaces {
		if sha3HashAsString, ok := emptyInterface.(string); ok && sha3HashAsString != "" {
			identifiedSHA3Hashes = append(identifiedSHA3Hashes, hash.NewIdentifiedSHA3Hash(
				hash.ID(hashIDs[i]),
				hash.SHA3Hash(sha3HashAsString),
			))
		}
	}

	return identifiedSHA3Hashes
}
