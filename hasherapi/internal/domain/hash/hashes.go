package hash

import "strconv"

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

func NewIDFromString(s string) ID {
	// TODO: handle an error
	id, _ := strconv.Atoi(s)

	return ID(id)
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
