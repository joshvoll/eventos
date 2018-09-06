package msgqueue

// Event hold the event interface
type Event interface {
	EventName() string
}
