package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joshvoll/eventos/src/eventservice/rest"
	"github.com/joshvoll/eventos/src/lib/configuration"
	msgqueue_amqp "github.com/joshvoll/eventos/src/lib/msgqueue/amqp"
	"github.com/joshvoll/eventos/src/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	// path of the configuration file
	confPath := flag.String("conf", `../../src/lib/configuration/config.json`, "flag to set the path of the configuration file")
	flag.Parse()

	// extract configuration file
	config, err := configuration.ExtractConfiguration(*confPath)
	if err != nil {
		log.Fatal(err)
	}

	// extract the amqp message broker
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		log.Fatal("ERROR AMQP: ", err)
	} else {
		fmt.Println("Connection succesfully...")
	}

	// connect the configuration file
	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
	if err != nil {
		log.Fatal("ERROR CONNECTING RABBIT MQ: ", err)
	}

	fmt.Println("Connecting to the Database...")
	// get dbtype, dbconnection and rest api
	dbhandler, err := dblayer.NewPersistenceLayer(config.DatabaseType, config.DatabaseConnection)
	if err != nil {
		log.Fatal("DATABASE ERROR: ", err)
	}

	// RestFullt APi start
	//httpErrChannel, httpTLSErrChannel := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndpoint, dbhandler, emitter)

	// select {
	// case err := <-httpErrChannel:
	// 	log.Fatal("HTTP Error: ", err)
	// case err := <-httpTLSErrChannel:
	// 	log.Fatal("HTTPS Error: ", err)
	// }

	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler, emitter))
}
