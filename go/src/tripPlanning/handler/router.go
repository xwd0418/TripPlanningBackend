package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/showDefaultPlaces", http.HandlerFunc(showDefaultPlacesHandler)).Methods("GET")

	//savePlaces could be a "put" OR potentially a "post" request
	router.Handle("/savePlace", http.HandleFunc(saveHandler)).Methods("PUT")
	fmt.Println("ready to receive requests")
	return router
}
