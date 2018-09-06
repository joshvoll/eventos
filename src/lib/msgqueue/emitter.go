package msgqueue

// EventEmitter interfaces for emitts events
type EventEmitter interface {
	Emit(e Event) error
}
