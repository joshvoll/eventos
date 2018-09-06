package persistence

import (
	"gopkg.in/mgo.v2/bson"
)

// Event definition
type Event struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

// Location definition
type Location struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}

// Hall definition
type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location, omitempty"`
	Capacity int    `json:"capacity"`
}
