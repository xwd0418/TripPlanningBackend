/*
	 reading user's plans, including:

		read user's all trip_plans overivew
		read each tipe_plan's daily plan
*/
package service

import (
	"fmt"
	"log"
	"tripPlanning/backend"
	"tripPlanning/model"
)

func ReadAllTripPlanInfo(userID string) ([]model.TripPlan, error) {
	rows, err := backend.ReadFromDB(backend.TableName_Trips,
		[]string{"tripname", "startday", "endday", "transportation"},
		fmt.Sprintf("userid=%s", userID))
	if err != nil {
		log.Println("Error during reading all plans' overview: ", err)
		return nil, err
	}
	defer rows.Close()
	var tripPlans []model.TripPlan
	for rows.Next() {
		var p model.TripPlan
		if err := rows.Scan(&p.TripName, &p.StartDay, &p.EndDay, &p.Transportation); err != nil {
			return nil, err
		}
		tripPlans = append(tripPlans, p)
	}

	return tripPlans, nil
}

func readAllDayPlansOfTripPlan(tripID string) ([]model.DayPlan, error) {
	return nil, nil
}
