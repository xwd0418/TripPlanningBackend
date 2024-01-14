package tests

import (
	"fmt"
	"log"
	"testing"
	"tripPlanning/backend"
	"tripPlanning/model"
	"tripPlanning/service"
)

func TestSavePlacesToDB(t *testing.T) {

	backend.InitDB()

	fake_user_table := map[string]interface{}{
		"userID":   "backend_dev_user_id",
		"username": "backend_dev",
		"password": "1234",
		"email":    "fake@123.com",
	}
	err := backend.InsertIntoDB(backend.TableName_Users, fake_user_table, "ON CONFLICT (userID) DO NOTHING;")
	if err != nil {
		log.Fatal("Error during store fake user: ", err)
	}

	day1Places, err := service.GetDefaultPlaces(3)
	if err != nil {
		log.Fatal("failed to generaete recommended places day 1", err)
	}
	day2Places, err := service.SearchPlaces("museums", 2)
	if err != nil {
		log.Fatal("failed to generaete museum places, day2", err)
	}

	allPlaces := [][]model.Place{day1Places, day2Places}
	fmt.Println("populating into DB")
	service.GeneratePlanAndSaveToDB("backend_dev_user_id", allPlaces, "2024-02-10", "2024-02-11", "transit", "backend_test_1")
}
