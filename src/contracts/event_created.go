package contracts

import (
	"time"
)

// EventCreatedEvent defining the struct of each event is going to be created
type EventCreatedEvent struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LocationID string    `json:"location_id"`
	Start      time.Time `json:"start_date"`
	End        time.Time `json:"end_date"`
}

// EventName return the name of the event created
func (c *EventCreatedEvent) EventName() string {
	return "event.created"
}
