package grpcapi

import (
	"common/hasherproto"
	"context"
	"hasher/domain/hash"
)

func NewServer() *Server {
	return &Server{}
}

type Server struct {
	hasherproto.HasherServiceServer
}

func (s *Server) CalculateSha3Hashes(_ context.Context, inputs *hasherproto.Inputs) (*hasherproto.Sha3Hashes, error) {
	hashInputs := hash.NewInputs(inputs.Inputs)

	sha3Hashes := hash.CalculateSHA3HashSum(hashInputs)

	return &hasherproto.Sha3Hashes{
		Hashes: sha3Hashes.ToStrings(),
	}, nil
}
