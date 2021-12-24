package hash

import (
	"encoding/hex"
	"sync"

	"golang.org/x/crypto/sha3"
)

func CalculateSHA3HashSum(inputs []Input) SHA3Hashes {
	if len(inputs) == 0 {
		return []SHA3{}
	}

	if len(inputs) == 1 {
		return []SHA3{calculateSHA3HashSum256(inputs[0])}
	}

	return calculateSHA3HashSumParallel(inputs)
}

func calculateSHA3HashSum256(input Input) SHA3 {
	hash := sha3.New256()
	hash.Write(input.value)

	hashBytes := hash.Sum(nil)

	return input.NewSHA3(hex.EncodeToString(hashBytes))
}

func calculateSHA3HashSumParallel(inputs []Input) SHA3Hashes {
	var wg sync.WaitGroup
	sha3HashesChan := make(chan SHA3, len(inputs))

	calculate := func(input Input) {
		sha3HashesChan <- calculateSHA3HashSum256(input)

		wg.Done()
	}

	wg.Add(len(inputs))

	for _, input := range inputs {
		go calculate(input)
	}

	wg.Wait()
	close(sha3HashesChan)

	sha3Hashes := make(SHA3Hashes, 0, len(inputs))
	for sha3Hash := range sha3HashesChan {
		sha3Hashes = append(sha3Hashes, sha3Hash)
	}

	return sha3Hashes
}
