package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tripPlanning/service"
)

func readUserGeneralTripsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received one request for read general trip info")
	w.Header().Set("Content-Type", "application/json")

	// 1. process request
	username := r.URL.Query().Get("username")
	// 2. call services to handle request

	tripPlans, err := service.ReadUserGeneralTripPlans(username)
	if err != nil {
		log.Printf("Failed to get tripPlans, error: %v", err)
		return
	}
	// log.Println(len(tripPlans))

	// 3. construct response  : post => json
	js, err := json.Marshal(tripPlans)
	if err != nil {
		http.Error(w, "Failed to parse trip-plans into JSON format", http.StatusInternalServerError)
		log.Printf("Failed to parse trip-plans into JSON format %v.\n", err)
		return
	}
	w.Write(js)
}

func readAllDayPlansOfTripPlanHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received one request for read day plans of each day, with given trip id")
	w.Header().Set("Content-Type", "application/json")

	// 1. process request
	trip_id := r.URL.Query().Get("trip_id")
	// 2. call services to handle request

	tripPlans, err := service.ReadAllDayPlansOfTripPlan(trip_id)
	if err != nil {
		log.Printf("Failed to get the day-plans, error: %v", err)
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
