package mongolayer

import (
	"github.com/joshvoll/eventos/src/lib/persistence"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Global properties
const (
	DB     = "eventos"
	USERS  = "users"
	EVENTS = "events"
)

// MongoDBLayer struct definition
type MongoDBLayer struct {
	session *mgo.Session
}

// NewMongoDBLayer contructor function, this will make the connection of the db
func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	// Dial up to the mongo db database
	s, err := mgo.Dial(connection)
	return &MongoDBLayer{
		session: s,
	}, err
}

// interfaces definition
func (mgolayer *MongoDBLayer) getFresshSession() *mgo.Session {
	return mgolayer.session.Copy()
}

// AddEvent implementation interface
func (mgolayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	// Get fresh session
	s := mgolayer.getFresshSession()
	defer s.Close()

	// define the id of location and event
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	// return byte and error
	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)

}

// FindEvent implementation interface
func (mgolayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	// get fresh session
	s := mgolayer.getFresshSession()
	defer s.Close()

	// properties declaration
	e := persistence.Event{}
	var err error

	// query db
	if err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e); err != nil {
		return e, err
	}

	// return struct and error
	return e, err
}

// FindEventByName implementation interface
func (mgolayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	// get fresh session
	s := mgolayer.getFresshSession()
	defer s.Close()

	// properties declaration
	e := persistence.Event{}
	var err error

	// query db
	if err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e); err != nil {
		return e, err
	}

	// return everythingr
	return e, err

}

// FindAllAvailableEvent implementation interface
func (mgolayer *MongoDBLayer) FindAllAvailableEvent() ([]persistence.Event, error) {
	// get fresh session
	s := mgolayer.getFresshSession()

	// local properties
	e := []persistence.Event{}
	var err error

	// query db
	if err := s.DB(DB).C(EVENTS).Find(nil).All(&e); err != nil {
		return e, err
	}

	// regurn struct and error
	return e, err
}
