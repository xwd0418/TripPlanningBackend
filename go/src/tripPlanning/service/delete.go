// at the service level -> needs a delete handler

package service

import (
	"fmt"
	"log"
	"tripPlanning/backend"
)

// Service function to delete a trip and associated day plans and place relations
func DeleteTripWithAssociations(tripID string) error {
	// Check if the trip exists
	exists, err := backend.CheckIfItemExistsInDB("trips", "tripid", tripID)
	if err != nil {
		log.Println("Error checking if the trip ID exists in the database:", err)
		return err
	}

	if !exists {
		// Return an error or handle the case where the trip doesn't exist
		return fmt.Errorf("trip with ID %s does not exist", tripID)
	}

	// Delete associated day plans
	err = DeleteDayPlansForTrip(tripID)
	if err != nil {
		log.Println("Error deleting associated day plans:", err)
		return err
	}

	// Delete the trip entry
	err = backend.DeleteFromDB("trips", "tripid", tripID)
	if err != nil {
		log.Println("Error deleting the trip:", err)
		return err
	}

	return nil
}

func DeleteDayPlansForTrip(tripID string) error {
    columns := []string{"dayplanid"}
    conditions := fmt.Sprintf("tripid = '%s'", tripID)
    rows, err := backend.ReadFromDB("dayplans", columns, conditions)
    if err != nil {
        log.Println("Error reading day plans for trip:", err)
        return err
    }
    defer rows.Close()

	var dayPlans []string

	// Extract day plan IDs from the rows
	for rows.Next() {
		var dayPlanID string
		err := rows.Scan(&dayPlanID)
		if err != nil {
			log.Println("Error scanning row:", err)
			return err
		}
		dayPlans = append(dayPlans, dayPlanID)
	}

	// Delete each day plan and associated place relations
	for _, dayPlanID := range dayPlans {
		err := DeleteDayPlanWithAssociations(dayPlanID)
		if err != nil {
			log.Println("Error deleting day plan with associations:", err)
			return err
		}
	}

	return nil
}

func DeleteDayPlanWithAssociations(dayPlanID string) error {
	// Check if the day plan exists
	exists, err := backend.CheckIfItemExistsInDB("dayplans", "dayplanid", dayPlanID)
	if err != nil {
		log.Println("Error checking if the day plan ID exists in the database:", err)
		return err
	}

	if !exists {
		// Return an error or handle the case where the day plan doesn't exist
		return fmt.Errorf("day plan with ID %s does not exist", dayPlanID)
	}

	// Delete associated place relations
	err = backend.DeleteFromDB("dayplacerelations", "dayplanid", dayPlanID)
	if err != nil {
		log.Println("Error deleting associated place relations:", err)
		return err
	}

	// Delete the day plan entry
	err = backend.DeleteFromDB("dayplans", "dayplanid", dayPlanID)
	if err != nil {
		log.Println("Error deleting the day plan:", err)
		return err
	}

	return nil
}
