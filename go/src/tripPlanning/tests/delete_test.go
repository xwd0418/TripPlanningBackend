package service_test

import (
	"log"
    "testing"
    "tripPlanning/backend"
    "tripPlanning/model"
    "tripPlanning/service"
)

func TestDeleteTripWithAssociations(t *testing.T) {
    err := backend.InitDB() 
	if err != nil {
		log.Printf("Failed to initialize DB: %v", err)
	}

	//create a new user and a new trip 
    newUser := &model.User{
        Username: "testUser20",
        Password: "testPassword",
        Email:    "testuser@example.com",
    }
    success, err := service.AddUser(newUser)
    if err != nil || !success {
        log.Println("Error adding new user: ", err)
    }

    // Use the newly added user's ID for the trip
    userID := newUser.Id

    // Generate a new trip plan and save it to the database
    placesOfAllDays := [][]model.Place{ /* ... populate with test data ... */ }
    startDay := "2024-02-10"
    endDay := "2024-02-12"
    transportation := "Bus"
    tripName := "Original Trip Name"
    tripID, err := service.GeneratePlanAndSaveToDB(userID, placesOfAllDays, startDay, endDay, transportation, tripName)
    if err != nil {
        log.Printf("Error during generating and storing trip: %v", err)
    }

    // Call the Delete function
    err = service.DeleteTripWithAssociations(tripID)
    if err != nil {
        log.Printf("Error deleting trip: %v", err)
    }

    // Verify that the trip has been deleted
    _, err = backend.CheckIfItemExistsInDB("trips", "tripid", tripID)
    if err == nil {
        log.Println("Trip was not deleted")
    }
}