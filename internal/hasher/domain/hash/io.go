package hash

import "fmt"

func NewInputs(strings []string) []Input {
	inputs := make([]Input, 0, len(strings))

	for i, s := range strings {
		inputs = append(inputs, NewInput(i, s))
	}

	return inputs
}

func NewInput(index int, value string) Input {
	return Input{
		index: index,
		value: []byte(value),
	}
}

type Input struct {
	index int
	value []byte
}

func (in Input) String() string {
	return fmt.Sprintf("{index: %d, value: %s}", in.index, in.value)
}

func (in Input) NewSHA3(value string) SHA3 {
	return NewSHA3(in.index, value)
}

func NewSHA3(index int, value string) SHA3 {
	return SHA3{
		index: index,
		value: string(value),
	}
}

type SHA3 struct {
	index int
	value string
}

func (sha3 SHA3) String() string {
	return fmt.Sprintf("{index: %d, value: %s}", sha3.index, sha3.value)
}

type SHA3Hashes []SHA3

func (outputs SHA3Hashes) ToStrings() []string {
	chunks := make([]string, len(outputs))

	for _, output := range outputs {
		chunks[output.index] = output.value
	}

	return chunks
}
