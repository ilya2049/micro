package handler

import (
	"hasherapi/domain/hash"
	"hasherapi/pkg/conv"
	"hasherapi/system/restapi/models"
	"hasherapi/system/restapi/operations"
)

func postSendParamsToHashInputs(params operations.PostSendParams) []hash.Input {
	hashInputs := make([]hash.Input, 0, len(params.Params))

	for _, s := range params.Params {
		hashInputs = append(hashInputs, hash.Input(s))
	}

	return hashInputs
}

func identifiedHashesToArrayOfHash(identifiedSHA3Hashes []hash.IdentifiedSHA3Hash) models.ArrayOfHash {
	arrayOfHashes := make(models.ArrayOfHash, 0, len(identifiedSHA3Hashes))

	for _, identifiedSHA3Hash := range identifiedSHA3Hashes {
		arrayOfHashes = append(arrayOfHashes, identifiedHashToHash(identifiedSHA3Hash))
	}

	return arrayOfHashes
}

func identifiedHashToHash(identifiedSHA3Hash hash.IdentifiedSHA3Hash) *models.Hash {
	return &models.Hash{
		Hash: conv.PointerString(string(identifiedSHA3Hash.SHA3Hash())),
		ID:   conv.PointerInt64(int64(identifiedSHA3Hash.ID())),
	}
}
