package eventstream

import (
	"hasherapi/app/log"

	"github.com/segmentio/kafka-go"
)

type KafkaWriterConfig struct {
	Host  string
	Topic string
}

func NewKafkaWriter(cfg KafkaWriterConfig, logger log.Logger) (kafkaWriter *kafka.Writer, disconnectKafka func()) {
	kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers:     []string{cfg.Host},
		Topic:       cfg.Topic,
		ErrorLogger: logger,
	})

	return kafkaWriter, func() {
		if err := kafkaWriter.Close(); err != nil {
			logger.LogWarn("failed to close the kafka writer: "+err.Error(), log.Details{
				log.FieldComponent: log.ComponentEventStream,
			})
		}
	}
}
