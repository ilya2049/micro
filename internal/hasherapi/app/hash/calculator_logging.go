package hash

import (
	"common/requestid"
	"context"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
)

func WrapCalculatorWithLogger(calculator hash.Calculator, logger log.Logger) hash.Calculator {
	return &calculatorLoggingWrapper{
		next:   calculator,
		logger: logger,
	}
}

type calculatorLoggingWrapper struct {
	next hash.Calculator

	logger log.Logger
}

func (c *calculatorLoggingWrapper) Calculate(ctx context.Context, inputs []hash.Input) (hash.SHA3Hashes, error) {
	sha3Hashes, err := c.next.Calculate(ctx, inputs)

	requestID := requestid.Get(ctx)

	c.logger.LogDebug("calculate", log.Details{
		log.FieldRequestID:      requestID,
		log.FieldHashInputs:     inputs,
		log.FieldComponent:      log.ComponentHashCalculator,
		log.FieldHashSHA3Hashes: sha3Hashes,
	})

	return sha3Hashes, err
}
