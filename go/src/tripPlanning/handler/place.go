package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tripPlanning/service"
)

func showDefaultPlacesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received one request for search")
	w.Header().Set("Content-Type", "application/json")

	// 1. process request
	max_num_display_str := r.URL.Query().Get("max_num_display")
	
	// Set a default value for max_num_display
	defaultMaxNumDisplay := 50
	max_num_display, err := strconv.Atoi(max_num_display_str)
	if err != nil {
		log.Printf("Failed to convert max_num_display from showDefaultPlacesHandler to integer: %v. Using default value %d", err, defaultMaxNumDisplay)
		max_num_display = defaultMaxNumDisplay
	}

	// 2. call services to handle request
	default_places, err := service.GetDefaultPlaces(max_num_display)
	if err != nil {
		log.Printf("Failed to get default places: %v", err)
		return
	}

	// 3. construct response  : post => json
	js, err := json.Marshal(default_places)
	if err != nil {
		http.Error(w, "Failed to parse places into JSON format", http.StatusInternalServerError)
		log.Printf("Failed to parse places into JSON format %v.\n", err)
		return
	}
	w.Write(js)
}

func searchPlacesPlacesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received one request for search")
	w.Header().Set("Content-Type", "application/json")

	// 1. process request
	max_num_display_str := r.URL.Query().Get("max_num_display")
	user_input := r.URL.Query().Get("user_input")
	// 2. call services to handle request
	max_num_display, err := strconv.Atoi(max_num_display_str)
	if err != nil {
		log.Printf("Failed to convert max_num_display from showDefaultPlacesHandler to integer: %v", err)
		return
	}
	default_places, err := service.SearchPlaces(user_input, max_num_display)
	if err != nil {
		log.Printf("Failed to get searched places: %v", err)
		return
	}
	// 3. construct response  : post => json
	js, err := json.Marshal(default_places)
	if err != nil {
		http.Error(w, "Failed to parse places into JSON format", http.StatusInternalServerError)
		log.Printf("Failed to parse places into JSON format %v.\n", err)
		return
	}
	w.Write(js)
}
