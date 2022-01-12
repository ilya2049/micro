package event

func NewSendCallEvent() *Event {
	return &Event{CallsSend: 1}
}

func NewCheckCallEvent() *Event {
	return &Event{CallsCheck: 1}
}

type Event struct {
	CallsSend  int8
	CallsCheck int8
}
