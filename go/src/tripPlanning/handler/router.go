package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/showDefaultPlaces", http.HandlerFunc(showDefaultPlacesHandler)).Methods("GET")
	return router
}
