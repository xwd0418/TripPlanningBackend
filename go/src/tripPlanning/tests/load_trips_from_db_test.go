package service_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"tripPlanning/backend"
	"tripPlanning/service"
)

func TestLoadAllTripsOfUser(t *testing.T) {

	backend.InitDB()

	tripPlans, err := service.ReadUserGeneralTripPlans(`'backend_dev_user_id_01_14'`)
	if err != nil {
		panic(" ReadUserGeneralTripPlans failed ")
	}

	tripPlans_json, err := json.Marshal(tripPlans)
	if err != nil {
		panic(" marshling error")
	}
	fmt.Print("dayplans:", string(tripPlans_json))
}
