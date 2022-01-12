package eventstream

import (
	"context"
	"hasherapi/app/event"
	"hasherapi/app/log"

	"github.com/segmentio/kafka-go"
)

func New(kafkaWriter *kafka.Writer, logger log.Logger, pullSize int,
) (stream event.Stream, stopStream func()) {
	pull, stopStream := newStreamPull(
		&eventStream{
			kafkaWriter: kafkaWriter,
			logger:      logger,
		},
		pullSize,
	)

	return pull, stopStream
}

type eventStream struct {
	kafkaWriter *kafka.Writer
	logger      log.Logger
}

var componentHTTPAPI = []byte(log.ComponentHTTPAPI)

func (w *eventStream) send(ctx context.Context, e *event.Event) {
	err := w.kafkaWriter.WriteMessages(ctx,
		kafka.Message{
			Key:   componentHTTPAPI,
			Value: newKafkaEvent(e).Bytes(),
		},
	)

	if err != nil {
		w.logger.LogWarn("failed to send a message in the broker: "+err.Error(), log.Details{
			log.FieldComponent: log.ComponentEventStream,
		})
	}
}
