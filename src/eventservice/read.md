package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// global properties
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672"
	}

	connection, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatal("Error ocurred : ", err)
	}

	defer connection.Close()

	// building the channle
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	// declaring message as struct
	message := amqp.Publishing{
		Body: []byte("hello world"),
	}

	// then use tu publish the message to the broker message
	err = channel.Publish("events", "some-routing-key", false, false, message)
	if err != nil {
		log.Fatal("Error publishing the message: ", err)
	}

	// declaring an exchange message to the broker
	_, err = channel.QueueDeclare("my_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Error ocurred creating the QueueDeclare method ", err)
	}

	err = channel.QueueBind("my_queue", "#", "events", false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// After having declared and bound a queue, you can now start consuming this queue. For this, use the channel's
	msgs, err := channel.Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// loop throw all the mesage
	for msg := range msgs {
		fmt.Println("Message received " + string(msg.Body))
		msg.Ack(false)
	}

}


// creating chanel to handler https or http
	// httpErrChan := make(chan error)
	// httpTLSErrChan := make(chan error)

	// // creating 2 go rutines to handler each request
	// go func() {
	// 	httpTLSErrChan <- http.ListenAndServeTLS(tlsendpoint, "cert.pem", "key.pem", r)
	// }()
	// go func() {
	// 	httpErrChan <- http.ListenAndServe(endpoint, r)
	// }()
	// return the http object with por
	//return httpErrChan, httpTLSErrChan