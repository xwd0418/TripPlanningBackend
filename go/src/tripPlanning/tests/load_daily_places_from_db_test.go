package service_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"tripPlanning/backend"
	"tripPlanning/service"
)

func TestLoadPlacesFromDB(t *testing.T) {

	backend.InitDB()

	tripPlan, err := service.ReadAllDayPlansOfTripPlan(`970e9ae8-7e15-4ce3-b048-183126ed635f`)
	if err != nil {
		panic(" ReadAllDayPlansOfTripPlan failed ")
	}

	tripPlans_json, err := json.Marshal(tripPlan)
	if err != nil {
		panic(" marshling error")
	}
	fmt.Print("the trip plan is :", (string(tripPlans_json)))
}
