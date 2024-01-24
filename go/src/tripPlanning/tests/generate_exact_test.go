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
        Username: "testUser21",
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
        t.Fatalf("Error during generating and storing trip: %v", err)
    }

    // Modify the trip
    newStartDay := "2024-02-15"
    newEndDay := "2024-02-20"
    newTransportation := "Car"
    newTripName := "Modified Trip Name"
    modifiedTripID, err := service.ModifyTrip(userID, tripID, placesOfAllDays, newStartDay, newEndDay, newTransportation, newTripName)
    if err != nil {
        t.Fatalf("Error modifying trip: %v", err)
    }
    log.Printf("The trip has been successfully modified. New trip ID: %s", modifiedTripID)
    // Assertions and verifications
    // Verify that a new trip ID is generated, and other expected changes occurred
}