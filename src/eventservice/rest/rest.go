package rest

import (
	"net/http"

	"github.com/joshvoll/eventos/src/lib/msgqueue"

	"github.com/gorilla/mux"
	"github.com/joshvoll/eventos/src/lib/persistence"
)

// ServeAPI get endpoint string, return the http server and port
func ServeAPI(endpoint string, dbhandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) error {
	// add the handler from handlers
	handler := newEventHandler(dbhandler, eventEmitter)

	// defining the router object
	r := mux.NewRouter()

	// defining the subrouter "/evetnts"
	eventrouter := r.PathPrefix("/events").Subrouter()

	// defining the routes
	eventrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(endpoint, r)

}
