package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tripPlanning/model"
	"tripPlanning/service"
)

// RequestData represents the JSON data structure expected in the request body.
type RequestData struct {
	Username        string          `json:"username"`
	StartDay        string          `json:"StartDay"`
	EndDay          string          `json:"EndDay"`
	PlacesOfEachDay [][]model.Place `json:"places"`
	Transportation  string          `json:"Transportation"`
	TripName        string          `json:"TripName"`
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
	tripPlans, err := service.GeneratePlanAndSaveToDB(requestData.Username, requestData.PlacesOfEachDay, requestData.StartDay,
		requestData.EndDay, requestData.Transportation, requestData.TripName)
	if err != nil {
		log.Printf("Failed to GeneratePlanAndSaveToDB, error: %v", err)
		return
	}

	// 3. construct response  : post => json
	js, err := json.Marshal(tripPlans)
	if err != nil {
		http.Error(w, "Failed to parse day-plans into JSON format", http.StatusInternalServerError)
		log.Printf("Failed to parse day-plans into JSON format %v.\n", err)
		return
	}
	w.Write(js)
}
