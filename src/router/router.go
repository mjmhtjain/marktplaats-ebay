package router

import (
	"github.com/gorilla/mux"
	"github.com/mjmhtjain/marktplaats-ebay/src/handlers"
)

func MuxRouter() *mux.Router {
	r := mux.NewRouter()
	r.
		Path("/health").
		Methods("GET").
		HandlerFunc(handlers.HealthHandler)

	ecgRoute := r.PathPrefix("/ecg").Subrouter()
	ecgRoute.
		Path("/upload").
		Methods("POST").
		HandlerFunc(handlers.HelloHandler)

	return r
}
