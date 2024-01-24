package handler

import (
	"encoding/json"
	"net/http"
	"tripPlanning/service"
	//"tripPlanning/model"
	"tripPlanning/constants"
	"fmt"
)

// AiGeneratedPlanHandler handles requests for AI-generated travel plans
func AiGeneratedPlanHandler(w http.ResponseWriter, r *http.Request) {
	// Assuming the frontend sends a JSON payload with city and startDay and EndDay
	var request struct {
		City 		   string 	 `json:"City"`
		StartDay       string    `json:"StartDay"`
		EndDay         string    `json:"EndDay"`
	}

	// Decode the JSON payload from the request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the received data
	if request.City == "" || request.StartDay  == "" || request.EndDay == ""{
		http.Error(w, "City and date are required", http.StatusBadRequest)
		return
	}

	// Create an instance of TravelPlannerService
	travelPlanner := service.NewTravelPlannerService(constants.Openai_key)
	// Call the AiGeneratedPlan method on the TravelPlannerService instance
	plan, err := travelPlanner.AiGeneratedPlan(request.City, request.StartDay, request.EndDay)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating plan: %v", err), http.StatusInternalServerError)
		return
	}

	// Send the generated plan as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plan)
}

