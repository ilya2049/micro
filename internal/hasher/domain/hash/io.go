package hash

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

type SHA3Hashes []SHA3

func (outputs SHA3Hashes) ToStrings() []string {
	chunks := make([]string, len(outputs))

	for _, output := range outputs {
		chunks[output.id] = output.value
	}

	return chunks
}
