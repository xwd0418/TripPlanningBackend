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

	dayPlans, err := service.ReadAllDayPlansOfTripPlan(`'87e51db2-0f14-4962-a56b-b44c3048732a'`)
	if err != nil {
		panic(" ReadAllDayPlansOfTripPlan failed ")
	}

	dayPlans_json, err := json.Marshal(dayPlans)
	if err != nil {
		panic(" marshling error")
	}
	fmt.Print("dayplans:", string(dayPlans_json))
}
