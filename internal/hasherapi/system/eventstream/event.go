package eventstream

import (
	"fmt"
	"hasherapi/app/event"
	"time"
)

type kafkaEvent struct {
	TimeStamp time.Time

	*event.Event
}

func newKafkaEvent(e *event.Event) kafkaEvent {
	return kafkaEvent{
		Event:     e,
		TimeStamp: time.Now(),
	}
}

func (e kafkaEvent) Bytes() []byte {
	return []byte(fmt.Sprintf("%s,%d,%d",
		e.TimeStamp.Format("2006-01-02 15:04:05"),
		e.CallsSend,
		e.CallsCheck,
	))
}
