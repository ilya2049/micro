package event

type Stream interface {
	Send(*Event)
}
