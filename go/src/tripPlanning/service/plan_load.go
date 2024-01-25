/*
	 reading user's plans, including:

		read user's all trip_plans overivew
		read each tipe_plan's daily plan
*/
package service

import (
	"fmt"
	"log"
	"sort"
	"time"
	"tripPlanning/backend"
	"tripPlanning/model"
)

var date_format_layout = "2006-01-02"

type idWithOder struct {
	ID    string
	order int
}
type byID []idWithOder

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func ReadUserGeneralTripPlans(username string) ([]model.TripPlan, error) {
	user, err := backend.GetUser(username)
	if err != nil {
		log.Fatal("Error during reading user based on username from DB: ", err)
		return nil, err
	}
	rows, err := backend.ReadFromDB(backend.TableName_Trips,
		[]string{"tripID", "tripname", "startday", "endday", "transportation", "SamplePlaceName"},
		fmt.Sprintf("userid='%s'", user.Id))
	if err != nil {
		log.Println("Error during reading all plans' overview: ", err)
		return nil, err
	}
	defer rows.Close()
	var tripPlans []model.TripPlan
	for rows.Next() {
		var p model.TripPlan
		if err := rows.Scan(&p.TripPlanId, &p.TripName, &p.StartDay, &p.EndDay, &p.Transportation, &p.SamplePlaceName); err != nil {
			return nil, err
		}
		tripPlans = append(tripPlans, p)
	}

	return tripPlans, nil
}

func ReadAllDayPlansOfTripPlan(tripID string) (model.TripPlan, error) {
	//the goal is to Return: a list of [date, transportation, [place_1, place_2, place_3]]
	// where Place: {place_id, lat, long}

	// get transportation and start_day
	var detailedTripPlan model.TripPlan
	sql_row_query := fmt.Sprintf("SELECT transportation, startDay, TripName FROM Trips WHERE tripID = '%s'", tripID)
	backend.ReadRowFromDB(sql_row_query).Scan(&detailedTripPlan.Transportation, &detailedTripPlan.StartDay, &detailedTripPlan.TripName)
	detailedTripPlan.TripPlanId = tripID
	// get all day_plan ID with their order
	rows, err := backend.ReadFromDB(backend.TableName_DayPlans, []string{"dayPlanID", "dayOrder"}, fmt.Sprintf("tripID='%s'", tripID))

	if err != nil {
		log.Println("Error during reading all day_plans' overview: ", err)
		return detailedTripPlan, err
	}
	defer rows.Close()
	var dayIDsWithOrder []idWithOder
	for rows.Next() {
		var day_datum idWithOder
		if err := rows.Scan(&day_datum.ID, &day_datum.order); err != nil {
			log.Println("Error during sacnning sql info to day_datum: ", err)
			return detailedTripPlan, err
		}
		dayIDsWithOrder = append(dayIDsWithOrder, day_datum)
	}
	sort.Sort(byID(dayIDsWithOrder))

	// construct dayplan for each day
	current_date, err := time.Parse(date_format_layout, detailedTripPlan.StartDay)
	if err != nil {
		log.Println("Error during parsing date string: ", err)
		return detailedTripPlan, err
	}

	// log.Printf("intial date string %s", start_day)

	for _, currentDay := range dayIDsWithOrder {
		// dayPlan.Date = current_date.Format(date_format_layout)
		current_date = current_date.Add(24 * time.Hour)
		// log.Printf("day id is %s", day.ID)
		currentDayPlaces, err := getPlacesOfDay(currentDay.ID)
		if err != nil {
			log.Println("Error during getting places details from a place id ", err)
			return detailedTripPlan, err
		}
		detailedTripPlan.Places = append(detailedTripPlan.Places, currentDayPlaces)
	}
	current_date = current_date.Add(-24 * time.Hour)
	detailedTripPlan.EndDay = current_date.Format(date_format_layout)
	return detailedTripPlan, nil
}

func getPlacesOfDay(dayID string) ([]model.Place, error) {
	// use sql query to find the day-places relation of each day
	rows, err := backend.ReadFromDB(backend.TableName_DayPlaceRelations, []string{"placeID", "visitOrder"}, fmt.Sprintf("dayPlanID='%s'", dayID))

	if err != nil {
		log.Println("Error during reading all place-day relations with given trip id: ", err)
		return nil, err
	}
	defer rows.Close()
	//  sort based on visited order
	var placesWithOrder []idWithOder

	// log.Printf("day id is %s ", dayID)
	for rows.Next() {
		var day_place_relation_datum idWithOder
		if err := rows.Scan(&day_place_relation_datum.ID, &day_place_relation_datum.order); err != nil {
			return nil, err
		}
		// log.Printf("place id is %s and order is %d", day_place_relation_datum.ID, day_place_relation_datum.order)
		placesWithOrder = append(placesWithOrder, day_place_relation_datum)
		// log.Printf("current place is %s", day_place_relation_datum.ID)
	}
	sort.Sort(byID(placesWithOrder))

	// place should inlude place_id, lat, long
	var detailedPlaces []model.Place
	for _, place := range placesWithOrder {
		// log.Println("place id", place.ID)
		var detailedPlace model.Place
		detailedPlace.Id = place.ID
		location_query := fmt.Sprintf("SELECT name, longitude, latitude FROM PlaceDetails WHERE placeID = '%s'", place.ID)
		backend.ReadRowFromDB(location_query).Scan(&detailedPlace.DisplayName.Text, &detailedPlace.Location.Longitude, &detailedPlace.Location.Latitude)
		detailedPlaces = append(detailedPlaces, detailedPlace)
	}
	return detailedPlaces, nil
}
