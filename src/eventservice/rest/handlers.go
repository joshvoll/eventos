package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/joshvoll/eventos/src/contracts"
	"github.com/joshvoll/eventos/src/lib/msgqueue"

	"github.com/gorilla/mux"

	"github.com/joshvoll/eventos/src/lib/persistence"
)

type eventServiceHandler struct {
	dbhandler    persistence.DatabaseHandler
	eventEmitter msgqueue.EventEmitter
}

// newEventHandler contructor functions
func newEventHandler(databasehandler persistence.DatabaseHandler, eventemitter msgqueue.EventEmitter) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler:    databasehandler,
		eventEmitter: eventemitter,
	}
}

// FindEventHandler implementation
func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	// creating the mux variables
	vars := mux.Vars(r)

	// defining the criteria
	criteria, ok := vars["SeachCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: No search criteria found, you can either search by id via /id/4 to search by name via /name/coldplay concert}`)
		return
	}

	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: No search key found, you can either serach by id via /id/4 to search by name vi /name/coldplayconcert}`)
		return
	}

	// defining properties for the searchkey
	var event persistence.Event
	var err error

	// check the criteria send it to the url
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchkey)
	case "id":
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}

	// handling errors
	if err != nil {
		fmt.Fprintf(w, "{error: %s}", err)
		return
	}

	// send everyting to the user
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	json.NewEncoder(w).Encode(&event)

}

// AllEventHandler implementation router
func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	// defining the local propierties
	events, err := eh.dbhandler.FindAllAvailableEvent()

	// error handling
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: error ocurred while trying to find all available events %s}`, err)
		return
	}

	// set headers and types
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: Error ocurred while trying to parse the json file %s}`, err)
	}
}

// newEventHandler implementation route
func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {
	// defining the routes
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: Error ocurred while trying to decode event data  %s}`, err)
		return
	}

	id, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: Error ocurred while tring to persistence event  %d %s}`, id, err)
		return
	}

	fmt.Fprint(w, `{"id":%d}`, id)

	// test the implementation of the brokers
	msg := contracts.EventCreatedEvent{
		ID:         hex.EncodeToString(id),
		Name:       event.Name,
		Start:      time.Unix(event.StartDate, 0),
		End:        time.Unix(event.EndDate, 0),
		LocationID: string(event.Location.ID),
	}

	// emit the event
	eh.eventEmitter.Emit(&msg)

	w.Header().Set("Content-Type", "application/json;charset=utf8")

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&event)

}
