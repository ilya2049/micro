package hash

import "fmt"

func NewInputs(strings []string) []Input {
	inputs := make([]Input, 0, len(strings))

	for i, s := range strings {
		inputs = append(inputs, NewInput(i, s))
	}

	return inputs
}

func NewInput(id int, value string) Input {
	return Input{
		id:    id,
		value: []byte(value),
	}
}

type Input struct {
	id    int
	value []byte
}

func (in Input) String() string {
	return fmt.Sprintf("{id: %d, value: %s}", in.id, in.value)
}

func (in Input) NewSHA3(value string) SHA3 {
	return NewSHA3(in.id, value)
}

func NewSHA3(id int, value string) SHA3 {
	return SHA3{
		id:    id,
		value: string(value),
	}
}

type SHA3 struct {
	id    int
	value string
}

func (sha3 SHA3) String() string {
	return fmt.Sprintf("{id: %d, value: %s}", sha3.id, sha3.value)
}

type SHA3Hashes []SHA3

func (outputs SHA3Hashes) ToStrings() []string {
	chunks := make([]string, len(outputs))

	for _, output := range outputs {
		chunks[output.id] = output.value
	}

	return chunks
}
