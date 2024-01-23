package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"tripPlanning/model"
	"tripPlanning/service"
)

func modifyTripHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to modify trip")

	//method check 
	if r.Method != "POST" {
		log.Println("Invalid request method: ", r.Method)
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	// Extracting the tripID from the URL path
	urlTripID := strings.TrimPrefix(r.URL.Path, "/modifyTrip/")
	if urlTripID == "" {
		log.Println("Trip ID not provided in the URL")
		http.Error(w, "Trip ID is required", http.StatusBadRequest)
		return
	}

	//request body decoding 
	var tripPlan model.TripPlan
	err := json.NewDecoder(r.Body).Decode(&tripPlan)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Decoded trip plan for modification: %v", tripPlan)

	//consistency check between URL tripID and tripID provided in the request body
	if urlTripID != tripPlan.TripPlanId {
		log.Println("Inconsistent trip IDs provided")
		http.Error(w, "Inconsistent trip IDs", http.StatusBadRequest)
		return
	}

	// Extracting placesOfAllDays from DayPlans
	var placesOfAllDays [][]model.Place
	for _, dayPlan := range tripPlan.DayPlans {
		placesOfAllDays = append(placesOfAllDays, dayPlan.PlacesToVisit)
	}

	newTripID, err := service.ModifyTrip(tripPlan.UserID, urlTripID, placesOfAllDays, tripPlan.StartDay, tripPlan.EndDay, tripPlan.Transportation, tripPlan.TripName) 
	if err != nil {
        log.Printf("Error modifying trip: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Trip modified successfully: newTripID = %s", newTripID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"new_trip_id": newTripID})
}