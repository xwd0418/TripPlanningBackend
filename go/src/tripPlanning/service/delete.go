// // at the service level -> needs a delete handler

package service

// import (
// 	"fmt"
// 	"log"
// 	"tripPlanning/backend"
// )

// // Service function to delete a trip and associated day plans and place relations
// func DeleteTripWithAssociations(tripID string) error {
// 	// Check if the trip exists
// 	exists, err := backend.CheckIfItemExistsInDB(backend.TableName_Trips, "tripID", tripID)
// 	if err != nil {
// 		log.Println("Error checking if the trip ID exists in the database:", err)
// 		return err
// 	}

// 	if !exists {
// 		// Return an error or handle the case where the trip doesn't exist
// 		return fmt.Errorf("Trip with ID %s does not exist", tripID)
// 	}

// 	// Delete associated day plans
// 	err = DeleteDayPlansForTrip(tripID)
// 	if err != nil {
// 		log.Println("Error deleting associated day plans:", err)
// 		return err
// 	}

// 	// Delete the trip entry
// 	err = backend.DeleteFromDB(backend.TableName_Trips, "tripID", tripID)
// 	if err != nil {
// 		log.Println("Error deleting the trip:", err)
// 		return err
// 	}

// 	return nil
// }

// func DeleteDayPlansForTrip(tripID string) error {
// 	// Get all day plans associated with the trip
// 	dayPlans, err := backend.GetDayPlansForTrip(tripID)
// 	if err != nil {
// 		log.Println("Error getting day plans for trip:", err)
// 		return err
// 	}

// 	// Delete each day plan and associated place relations
// 	for _, dayPlanID := range dayPlans {
// 		err = DeleteDayPlanWithAssociations(dayPlanID)
// 		if err != nil {
// 			log.Println("Error deleting day plan with associations:", err)
// 			return err
// 		}
// 	}

// 	return nil
// }

// func DeleteDayPlanWithAssociations(dayPlanID string) error {
// 	// Check if the day plan exists
// 	exists, err := backend.CheckIfItemExistsInDB(backend.TableName_DayPlans, "dayPlanID", dayPlanID)
// 	if err != nil {
// 		log.Println("Error checking if the day plan ID exists in the database:", err)
// 		return err
// 	}

// 	if !exists {
// 		// Return an error or handle the case where the day plan doesn't exist
// 		return fmt.Errorf("Day plan with ID %s does not exist", dayPlanID)
// 	}

// 	// Delete associated place relations
// 	err = backend.DeleteFromDB(backend.TableName_DayPlaceRelations, "dayPlanID", dayPlanID)
// 	if err != nil {
// 		log.Println("Error deleting associated place relations:", err)
// 		return err
// 	}

// 	// Delete the day plan entry
// 	err = backend.DeleteFromDB(backend.TableName_DayPlans, "dayPlanID", dayPlanID)
// 	if err != nil {
// 		log.Println("Error deleting the day plan:", err)
// 		return err
// 	}

// 	return nil
// }
