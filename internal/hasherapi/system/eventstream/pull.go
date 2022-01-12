package eventstream

import (
	"context"
	"hasherapi/app/event"
)

type streamPull struct {
	stream *eventStream

	events chan *event.Event
}

func newStreamPull(stream *eventStream, size int) (pull *streamPull, stopPull func()) {
	pull = &streamPull{
		stream: stream,

		events: make(chan *event.Event, size),
	}

	for i := 0; i < size; i++ {
		go func() {
			for e := range pull.events {
				stream.send(context.Background(), e)
			}
		}()
	}

	return pull, func() {
		close(pull.events)
	}
}

func (w *streamPull) Send(e *event.Event) {
	w.events <- e
}
