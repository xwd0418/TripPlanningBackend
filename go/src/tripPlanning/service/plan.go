/*
	 Manage user's plans, including:
		generate a plan from selected places and save it to database
		delete a plan
		read user's plans
*/
package service

import (
	"fmt"
	"tripPlanning/model"
)


func GeneratePlanAndSaveToDB(UserID, String, start_day string, end_day string, transportation string, places []model.Place) (error){

	
	groupedPlaces,err := groupPlacesByDate(start_day, end_day, transportation, places)
	if err != nil {
        fmt.Println("Error grouping places: ", err)
        return err
    }
	// save userID, PlanID

}

// core algorithm of backend: a method to group places for each day during the trip.
// The grouping strategy involves: place open hours, save transpotation time, rational schedule of restaurant and tourist attactions 
func groupPlacesByDate(start_day string, end_day string, transportation string, places []model.Place) ([][]model.Place, error){
	
}

func generateDayPlansAndSaveToDB(start_day string, end_day string, transportation string, grouped_places [][]model.Place){
	// save planID, day_plan_id
	// save plan_id, 
}