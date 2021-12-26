package router

import (
	"quoteapi/service"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/quotes", service.Quotes).Methods("POST", "OPTIONS")

	return router
}
