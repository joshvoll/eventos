package dblayer

import (
	"github.com/joshvoll/eventos/src/lib/persistence"
	"github.com/joshvoll/eventos/src/lib/persistence/mongolayer"
)

// DBTYPE is the type of the database
type DBTYPE string

// MONGODB and other stff
const (
	MONGODB    DBTYPE = "mongodb"
	DOCUMENTDB DBTYPE = "documentdb"
	DYNAMODB   DBTYPE = "dynamodb"
)

// NewPersistenceLayer DBtye, connection. Return: persistence.DatabaseHandler, error
func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	// get all the option to send back the db layer
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}

	return nil, nil
}
