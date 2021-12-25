package calculator

import (
	"common/errors"
	"context"
	"time"

	"common/requestid"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
	"hasherapi/pkg/grpcutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"common/hasherproto"
)

type GRPCCalculator struct {
	url     string
	timeout time.Duration
	logger  log.Logger
}

func NewGRPCCalculator(url string, timeout time.Duration, logger log.Logger) *GRPCCalculator {
	return &GRPCCalculator{
		url:     url,
		timeout: timeout,
		logger:  logger,
	}
}

func (c *GRPCCalculator) openConnection(ctx context.Context) (*grpcutil.Connection, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)

	clientConnection, err := grpc.DialContext(
		ctx,
		c.url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		cancel()

		return nil, err
	}

	requestID := requestid.Get(ctx)
	ctx = metadata.AppendToOutgoingContext(ctx,
		requestid.Header, requestID,
	)

	return grpcutil.NewConnection(ctx, clientConnection,
			func() {
				if err := clientConnection.Close(); err != nil {
					c.logger.LogWarn("failed to close an grpc connection: "+err.Error(), log.Details{
						log.FieldRequestID: requestID,
						log.FieldComponent: log.ComponentHashCalculator,
					})
				}

				cancel()
			}),
		nil
}

func (c *GRPCCalculator) Calculate(ctx context.Context, hashInputs []hash.Input) (hash.SHA3Hashes, error) {
	connection, err := c.openConnection(ctx)
	if err != nil {
		return []hash.SHA3Hash{}, errors.Errorf(
			"%s: failed to open a grpc connection: %w", log.ComponentHashCalculator, err,
		)
	}

	defer connection.Close()

	hasherClient := hasherproto.NewHasherServiceClient(connection.ClientConnection())

	protoSHA3Hashes, err := hasherClient.CalculateSha3Hashes(
		connection.Context(),
		hashInputsToProtoInputs(hashInputs),
	)
	if err != nil {
		return hash.SHA3Hashes{}, errors.Errorf(
			"%s: failed to fetch sha3 hashes: %w", log.ComponentHashCalculator, err,
		)
	}

	return hash.NewSHA3Hashes(protoSHA3Hashes.Hashes), nil
}
