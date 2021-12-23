package hash

import (
	"context"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
	"hasherapi/pkg/httputil/requestid"
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

const messageHashCalculatorCalculate = "hashCalculator.Calculate"

func (c *calculatorLoggingWrapper) Calculate(ctx context.Context, inputs []hash.Input) (hash.SHA3Hashes, error) {
	requestID := requestid.Get(ctx)

	c.logger.LogDebug(messageHashCalculatorCalculate, log.Details{
		log.FieldRequestID:  requestID,
		log.FieldHashInputs: inputs,
	})

	sha3Hashes, err := c.next.Calculate(ctx, inputs)

	if err == nil {
		c.logger.LogDebug(messageHashCalculatorCalculate, log.Details{
			log.FieldRequestID:      requestID,
			log.FieldHashSHA3Hashes: sha3Hashes,
		})
	}

	return sha3Hashes, err
}
