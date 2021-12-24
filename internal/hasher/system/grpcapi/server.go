package grpcapi

import (
	"common/hasherproto"
	"context"
	appHash "hasher/app/hash"
	"hasher/app/log"
	"hasher/domain/hash"
)

func NewServer(logger log.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

type Server struct {
	hasherproto.HasherServiceServer

	logger log.Logger
}

func (s *Server) CalculateSha3Hashes(ctx context.Context, inputs *hasherproto.Inputs) (*hasherproto.Sha3Hashes, error) {
	hashInputs := hash.NewInputs(inputs.Inputs)

	sha3Hashes := appHash.CalculateSHA3HashSum(ctx, s.logger)(hashInputs)

	return &hasherproto.Sha3Hashes{
		Hashes: sha3Hashes.ToStrings(),
	}, nil
}
