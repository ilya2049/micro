package calculator

import (
	"context"
	"fmt"
	"time"

	"common/requestid"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
	"hasherapi/pkg/grpcutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

	return grpcutil.NewConnection(ctx, clientConnection,
			func() {
				if err := clientConnection.Close(); err != nil {
					c.logger.LogWarn(systemMessageComponent+": failed to close an grpc connection", log.Details{
						log.FieldError:     err.Error(),
						log.FieldRequestID: requestID,
					})
				}

				cancel()
			}),
		nil
}

func (c *GRPCCalculator) Calculate(ctx context.Context, hashInputs []hash.Input) (hash.SHA3Hashes, error) {
	connection, err := c.openConnection(ctx)
	if err != nil {
		return []hash.SHA3Hash{}, fmt.Errorf(
			"%s: failed to open a grpc connection: %w", systemMessageComponent, err,
		)
	}

	defer connection.Close()

	hasherClient := hasherproto.NewHasherServiceClient(connection.ClientConnection())

	protoSHA3Hashes, err := hasherClient.CalculateSha3Hashes(
		connection.Context(),
		hashInputsToProtoInputs(hashInputs),
	)
	if err != nil {
		return hash.SHA3Hashes{}, fmt.Errorf(
			"%s: failed to fetch sha3 hashes: %w", systemMessageComponent, err,
		)
	}

	return hash.NewSHA3Hashes(protoSHA3Hashes.Hashes), nil
}
