package handler

import (
	"encoding/json"
	"log"
	"net/http"
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

	// // Extracting the tripID from the URL path
	// urlTripID := strings.TrimPrefix(r.URL.Path, "/modifyTrip/")
	// if urlTripID == "" {
	// 	log.Println("Trip ID not provided in the URL")
	// 	http.Error(w, "Trip ID is required", http.StatusBadRequest)
	// 	return
	// }

	//request body decoding
	var tripPlan model.TripPlan
	err := json.NewDecoder(r.Body).Decode(&tripPlan)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := r.URL.Query().Get("username")

	// log.Printf("Decoded trip plan for modification: %v", tripPlan)

	// //consistency check between URL tripID and tripID provided in the request body
	// if urlTripID != tripPlan.TripPlanId {
	// 	log.Println("Inconsistent trip IDs provided")
	// 	http.Error(w, "Inconsistent trip IDs", http.StatusBadRequest)
	// 	return
	// }

	// Extracting placesOfAllDays from DayPlans
	placesOfAllDays := tripPlan.Places
	// for _, dayPlan := range tripPlan.DayPlans {
	// 	placesOfAllDays = append(placesOfAllDays, dayPlan.PlacesToVisit)
	// }

	newTripID, err := service.ModifyTrip(username, tripPlan.TripPlanId, placesOfAllDays, tripPlan.StartDay, tripPlan.EndDay, tripPlan.Transportation, tripPlan.TripName)
	if err != nil {
		log.Printf("Error modifying trip: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Trip modified successfully: newTripID = %s", newTripID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"new_trip_id": newTripID})
}

func generateExactTripHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to modify trip")

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var exactTripPlan model.TripPlan
	err := json.NewDecoder(r.Body).Decode(&exactTripPlan)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exactPlacesOfAllDays := exactTripPlan.Places
	// var exactPlacesOfAllDays [][]model.Place
	// for _, dayPlan := range exactTripPlan.DayPlans {
	// 	exactPlacesOfAllDays = append(exactPlacesOfAllDays, dayPlan.PlacesToVisit)
	// }

	tripID, err := service.GenerateExactTrip(exactTripPlan.UserID, exactPlacesOfAllDays, exactTripPlan.StartDay, exactTripPlan.EndDay, exactTripPlan.Transportation, exactTripPlan.TripName)
	if err != nil {
		log.Println("Error in generating trip:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"tripID": tripID})
}
