package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/showDefaultPlaces", http.HandlerFunc(showDefaultPlacesHandler)).Methods("GET")
	fmt.Println("ready to receive requests")
	return router
}
