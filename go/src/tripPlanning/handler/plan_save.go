package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tripPlanning/model"
	"tripPlanning/service"
)

// RequestData represents the JSON data structure expected in the request body.
type RequestData struct {
	UserID          string          `json:"user_id"`
	StartDay        string          `json:"start_day"`
	EndDay          string          `json:"end_day"`
	PlacesOfEachDay [][]model.Place `json:"place_ids_of_each_day"`
	Transportation  string          `json:"transportation"`
	TripName        string          `json:"trip_name"`
}

// this will return the tripId in database
func GeneratePlanAndSaveHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST.
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON data from the request body.
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to decode JSON data.", http.StatusBadRequest)
		return
	}

	// call services
	tripID, err := service.GeneratePlanAndSaveToDB(requestData.UserID, requestData.PlacesOfEachDay, requestData.StartDay,
		requestData.EndDay, requestData.Transportation, requestData.TripName)
	if err != nil {
		log.Printf("Failed to GeneratePlanAndSaveToDB : %v", err)
		return
	}
	// Send a response back to the client.
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tripID))
	// fmt.Fprintln(w, tripID)
	fmt.Fprintln(w, "GeneratePlanAndSave request processed successfully.")
}
