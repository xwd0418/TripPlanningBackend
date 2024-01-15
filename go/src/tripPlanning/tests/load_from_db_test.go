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

	dayPlans, err := service.ReadAllDayPlansOfTripPlan(`'9b138350-ae23-4401-b3c9-d7489c54cef5'`)
	if err != nil {
		panic(" ReadAllDayPlansOfTripPlan failed ")
	}

	dayPlans_json, err := json.Marshal(dayPlans)
	if err != nil {
		panic(" marshling error")
	}
	fmt.Print("dayplans:", string(dayPlans_json))
}
