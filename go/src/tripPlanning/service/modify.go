package service

import (
	"fmt"
	"log"
	"tripPlanning/backend"
	"github.com/pborman/uuid"
	"tripPlanning/model"

)

//for the placesOfAllDays, we expect the place in the exact order as the user modifies
func ModifyTrip(userID string, tripID string, placesOfAllDays [][]model.Place, startDay string, endDay string, transportation string, tripName string) (string, error) {
	// Check if the trip exists
	exists, err := backend.CheckIfItemExistsInDB(backend.TableName_Trips, "tripID", tripID)
	if err != nil {
		log.Println("Error checking if the trip ID exists in the database:", err)
		return "", err
	}

	if !exists {
		// Return an error or handle the case where the trip doesn't exist
		return "", fmt.Errorf("Trip with ID %s does not exist", tripID)
	}

	// Delete the original trip
	err = DeleteTripWithAssociations(tripID)
	if err != nil {
		log.Println("Error deleting the original trip:", err)
		return "", err
	}

	// Generate a new trip with the provided parameters
	// with the place visit order exactly as passed in 
	newTripID, err := GenerateExactTrip(userID, placesOfAllDays, startDay, endDay, transportation, tripName)
	if err != nil {
		log.Println("Error generating a new trip:", err)
		return "", err
	}

	return newTripID, nil
}

//regenerate a trip with the day order and place visit order
//exactly as the user modifies at the frontend through [][]model.Place
func GenerateExactTrip(userID string, placesOfAllDays [][]model.Place, startDay string, endDay string, transportation string, tripName string) (string, error) {

	// 1. create a new TripPlan for this user
	tripID := uuid.New()
	tripTableEntry := map[string]interface{}{
		"tripID":         tripID,
		"userID":         userID,
		"tripName":       tripName,
		"startDay":       startDay,
		"endDay":         endDay,
		"transportation": transportation,
	}
	err := backend.InsertIntoDB(backend.TableName_Trips, tripTableEntry)
	if err != nil {
		log.Println("Error during store new trip plan: ", err)
		return "", err
	}
	
	// 2. get routes for each day
	var plannedRoutes [][]model.Place

	for _, placesEachDay := range placesOfAllDays {
		// No need to sort, directly use the order as passed in
		plannedRoutes = append(plannedRoutes, placesEachDay)
	}

	// 3. save routes to db
	for dayOrder, plannedRoute := range plannedRoutes {
		// 3.1 save each dayPlan to DB
		currentDayPlanId := uuid.New()
		tripTableEntry := map[string]interface{}{
			"dayPlanID": currentDayPlanId,
			"tripID":    tripID,
			"dayOrder":  dayOrder + 1,
		}
		err = backend.InsertIntoDB(backend.TableName_DayPlans, tripTableEntry)
		if err != nil {
			log.Println("Error during store new day-plan: ", err)
			return "", err
		}
	
		// 3.2 save each place of the day
		for visitOrder, place := range plannedRoute {
		
			placeID := place.Id
			placeIsInDB, err := backend.CheckIfItemExistsInDB(backend.TableName_PlaceDetails, "placeID", placeID)
			if err != nil {
				log.Println("Error during checking if place ID already exists: ", err)
				return "", err
			}
			if !placeIsInDB {
				err = SavePlaceToDB(place)
				if err != nil {
					log.Println("Error during store new trip place: ", err)
					return "", err
				}
			}
	
			// 3.2.2 save the day-place relation
			dayPlaceRelationEntry := map[string]interface{}{
				"placeID":    placeID,
				"dayPlanID":  currentDayPlanId,
				"visitOrder": visitOrder + 1,
			}
			err = backend.InsertIntoDB(backend.TableName_DayPlaceRelations, dayPlaceRelationEntry)
			if err != nil {
				log.Println("Error during store new day-place relation: ", err)
				return "", err
			}
		}
	}
	return tripID, nil
}

