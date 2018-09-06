package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joshvoll/eventos/src/lib/persistence/dblayer"
)

// DBTypeDefault all
var (
	DBTypeDefault            = dblayer.DBTYPE("mongodb")
	DBConnectionDefault      = "mongodb://127.0.0.1"
	RestfulEPDefault         = "localhost:8181"
	ResfulTLSEPDefault       = "localhost:9191"
	AMQPMessageBrokerDefault = "amqp://guest:guest@localhost:5672"
)

// ServiceConfig struct
type ServiceConfig struct {
	DatabaseType       dblayer.DBTYPE `json:"databasetype"`
	DatabaseConnection string         `json:"dbconnection"`
	RestfulEndpoint    string         `json:"restfulapi_endpoint"`
	RestfulTLSEndpoint string         `json:"restfulapi_tlsendpoint"`
	AMQPMessageBroker  string         `json:"amqp_message_broker"`
}

// ExtractConfiguration GET: string. Return: ServiceConfig, error
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	// define local properties
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
		ResfulTLSEPDefault,
		AMQPMessageBrokerDefault,
	}

	// read the file from the file system
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
