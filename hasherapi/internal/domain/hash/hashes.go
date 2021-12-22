package hash

import (
	"errors"
	"fmt"
	"strconv"
)

type Input string

func NewInputs(strings []string) []Input {
	calculationUnits := make([]Input, 0, len(strings))

	for _, s := range strings {
		calculationUnits = append(calculationUnits, Input(s))
	}

	return calculationUnits
}

type SHA3Hash string

type SHA3Hashes []SHA3Hash

func NewSHA3Hashes(strings []string) SHA3Hashes {
	hashes := make(SHA3Hashes, 0, len(strings))

	for _, s := range strings {
		hashes = append(hashes, SHA3Hash(s))
	}

	return hashes
}

func (sha3Hashes SHA3Hashes) NewIdentifiedSHA3Hashes(hashIDs []ID) []IdentifiedSHA3Hash {
	identifiedHashes := make([]IdentifiedSHA3Hash, 0, len(sha3Hashes))

	for i, sha3Hash := range sha3Hashes {
		identifiedHashes = append(identifiedHashes, IdentifiedSHA3Hash{
			id:       hashIDs[i],
			sha3Hash: sha3Hash,
		})
	}

	return identifiedHashes
}

type ID int

var errIDMustBeInt = errors.New("hash id must be an integer value")

func newIDFromString(s string) (ID, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", errIDMustBeInt, s)
	}

	return ID(id), nil
}

func NewIDsFromStrings(strings []string) ([]ID, error) {
	var (
		id  ID
		ids = make([]ID, 0, len(strings))
		err error
	)

	for _, s := range strings {
		id, err = newIDFromString(s)
		if err != nil {
			return []ID{}, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

type IdentifiedSHA3Hash struct {
	id       ID
	sha3Hash SHA3Hash
}

func NewIdentifiedSHA3Hash(id ID, sha3Hash SHA3Hash) IdentifiedSHA3Hash {
	return IdentifiedSHA3Hash{
		id:       id,
		sha3Hash: sha3Hash,
	}
}

func (h IdentifiedSHA3Hash) SHA3Hash() SHA3Hash {
	return h.sha3Hash
}

func (h IdentifiedSHA3Hash) ID() ID {
	return h.id
}
