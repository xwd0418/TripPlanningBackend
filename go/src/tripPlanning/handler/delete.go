package handler

import (
	"fmt"
	"net/http"
	"tripPlanning/service"

	"github.com/gorilla/mux"
)

func DeleteTripHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one 'delete trip' request")

	params := mux.Vars(r)
	tripID := params["tripID"]

	err := service.DeleteTripWithAssociations(tripID)
	if err != nil {
		http.Error(w, "Failed to delete trip", http.StatusInternalServerError)
		fmt.Printf("Error deleting trip: %v\n", err)
		return
	}

	fmt.Println("Trip deleted successfully.")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trip deleted successfully"))
}
