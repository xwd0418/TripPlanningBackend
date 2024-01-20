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

	dayPlans, err := service.ReadAllDayPlansOfTripPlan(`94625430-7351-4f6c-8c26-220e2c5e0ccf`)
	if err != nil {
		panic(" ReadAllDayPlansOfTripPlan failed ")
	}

	dayPlans_json, err := json.Marshal(dayPlans)
	if err != nil {
		panic(" marshling error")
	}
	fmt.Print("dayplans:", (string(dayPlans_json)))
}
