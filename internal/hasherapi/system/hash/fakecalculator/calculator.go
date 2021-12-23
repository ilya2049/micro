package fakecalculator

import (
	"context"
	"hasherapi/domain/hash"
)

func New() *HashCalculator {
	return &HashCalculator{}
}

type HashCalculator struct {
}

const hashPrefix = "hash-of-"

func (c *HashCalculator) Calculate(_ context.Context, inputs []hash.Input) (hash.SHA3Hashes, error) {
	hashes := make(hash.SHA3Hashes, 0, len(inputs))

	for _, input := range inputs {
		hashes = append(hashes, hash.SHA3Hash(hashPrefix+input))
	}

	return hashes, nil
}
