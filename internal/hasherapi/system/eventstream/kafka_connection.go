package eventstream

import (
	"fmt"
	"time"

	"hasherapi/app/log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type KafkaWriterConfig struct {
	Host     string
	Username string
	Password string
	Topic    string
}

func NewKafkaWriter(cfg KafkaWriterConfig, logger log.Logger) (kafkaWriter *kafka.Writer, disconnectKafka func(), err error) {
	var saslMechanism sasl.Mechanism
	saslMechanism, err = scram.Mechanism(scram.SHA512, cfg.Username, cfg.Password)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: failed to create a scram mechanism: %w", log.ComponentEventStream, err)
	}

	kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{cfg.Host},
		Topic:   cfg.Topic,
		Dialer: &kafka.Dialer{
			Timeout:       10 * time.Second,
			DualStack:     true,
			SASLMechanism: saslMechanism,
		},
		ErrorLogger: logger,
	})

	return kafkaWriter, func() {
		if err := kafkaWriter.Close(); err != nil {
			logger.LogWarn("failed to close the kafka writer: "+err.Error(), log.Details{
				log.FieldComponent: log.ComponentEventStream,
			})
		}
	}, nil
}
