package amqp

import (
	"encoding/json"
	"fmt"

	"github.com/joshvoll/eventos/src/lib/msgqueue"
	"github.com/streadway/amqp"
)

// ampq Event Emitter structure
type amqpEventEmitter struct {
	connection *amqp.Connection
	exchange   string
	events     chan *emittedEvents
}

type emittedEvents struct {
	event    msgqueue.Event
	erroChan chan error
}

/* implementing the event emitter for the struct */

// NewAMQPEventEmitter contructure function
func NewAMQPEventEmitter(conn *amqp.Connection, exChange string) (msgqueue.EventEmitter, error) {
	emitter := amqpEventEmitter{
		connection: conn,
		exchange:   exChange,
	}

	// check the error
	err := emitter.setup()
	if err != nil {
		return nil, err
	}

	return &emitter, nil

}

// setup the implementation of the listernet
func (a *amqpEventEmitter) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	// we need to return the channel exchange
	err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
	return err
}

// Emit method with json file format
func (a *amqpEventEmitter) Emit(e msgqueue.Event) error {
	// Next, we can create a new AMQP channel and publish our message to the events exchange
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	// defer the connection
	defer channel.Close()

	jsonDoc, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("could not JSON-serialize event: %s", err)
	}

	// defining the publication channel
	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": e.EventName()},
		ContentType: "application/json",
		Body:        jsonDoc,
	}

	// return eveything to publis
	err = channel.Publish(a.exchange, e.EventName(), false, false, msg)
	return err
}
