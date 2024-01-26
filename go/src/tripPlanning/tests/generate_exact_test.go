package service_test

import (
    "log"
    "testing"
    "tripPlanning/backend"
    "tripPlanning/model"
    "tripPlanning/service"
)

func TestGenerateExactTrip(t *testing.T) {
    backend.InitDB()

    // Create and add a new user
	// Username for testing needs to be unique every time
    newUser := &model.User{
        Username: "testUser23",
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
    tripName := "Trip Name"
    tripID, err := service.GenerateExactTrip(userID, placesOfAllDays, startDay, endDay, transportation, tripName)
    if err != nil {
        t.Fatalf("Error during generating and storing trip: %v", err)
    }
    if tripID == "" {
		log.Println("Failed to generate the exact trip plan as requested")
	}

    log.Printf("The trip has been successfully generated as requested. Trip ID: %s", tripID)

}