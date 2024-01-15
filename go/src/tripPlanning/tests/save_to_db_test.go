package service_test

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

	userID := "backend_dev_user_id_01_14"
	fake_user_table := map[string]interface{}{
		"userID":   userID,
		"username": "backend_dev_01_14",
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
	day3Places, err := service.SearchPlaces("parks", 4)
	if err != nil {
		log.Fatal("failed to generaete park places, day3", err)
	}

	allPlaces := [][]model.Place{day1Places, day2Places, day3Places}
	fmt.Println("populating into DB")
	tripID, err := service.GeneratePlanAndSaveToDB(userID, allPlaces, "2024-02-10", "2024-02-11", "transit", "backend_test_loading_all_trips")

	if err != nil {
		log.Fatal("failed to generaete routes", err)
	}
	log.Printf("generated tripID is %s", tripID)
}
