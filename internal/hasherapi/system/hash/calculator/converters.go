package calculator

import (
	"common/hasherproto"
	"hasherapi/domain/hash"
)

func hashInputsToProtoInputs(hashInputs []hash.Input) *hasherproto.Inputs {
	protoInputs := make([]string, 0, len(hashInputs))

	for _, hashInput := range hashInputs {
		protoInputs = append(protoInputs, string(hashInput))
	}

	return &hasherproto.Inputs{
		Inputs: protoInputs,
	}
}
