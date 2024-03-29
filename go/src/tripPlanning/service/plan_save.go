/*
	 Manage user's plans, including:
		generate a plan from selected places and save it to database
		delete a plan
*/
package service

import (
	// "fmt"
	"log"
	"strings"
	"tripPlanning/backend"
	"tripPlanning/model"

	"github.com/pborman/uuid"
)

func GeneratePlanAndSaveToDB(username string, placesOfAllDays [][]model.Place,
	startDay string, endDay string, transportation string, tripName string) (model.TripPlan, error) {
	// params:
	// placesOfAllDays: Each sub-array represent the planned places to visit each day
	// return： Trip ID

	var plansToShow model.TripPlan
	// 1. create a new TirpPlan for this user
	tripID := uuid.New()
	user, err := backend.GetUser(username)
	if err != nil {
		log.Fatal("Error during reading user based on username from DB: ", err)
		return plansToShow, err
	}
	tripTableEntry := map[string]interface{}{
		"tripID":          tripID,
		"userID":          user.Id,
		"tripName":        tripName,
		"startDay":        startDay,
		"endDay":          endDay,
		"transportation":  transportation,
		"SamplePlaceName": placesOfAllDays[0][0].DisplayName.Text,
	}
	err = backend.InsertIntoDB(backend.TableName_Trips, tripTableEntry)
	if err != nil {
		log.Fatal("Error during store new trip plan: ", err)
		return plansToShow, err
	}

	// 2. plan route for each day,
	var plannedRoutes [][]model.Place
	for _, placesEachDay := range placesOfAllDays {
		// fmt.Println("placesEachDay:", placesEachDay)
		sortedPlaces, err := GenerateDayPlan(placesEachDay, transportation, "")
		if err != nil {
			log.Fatal("Error during sorting places for a day: ", err)
			return plansToShow, err
		}
		plannedRoutes = append(plannedRoutes, sortedPlaces)
	}
	// log.Printf("size of plannedRoutes is %d, %d", len(plannedRoutes), len(plannedRoutes[0]))
	// 3. save planned routes to db
	for dayOrder, planedRoute := range plannedRoutes {
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
			return plansToShow, err
		}
		// log.Printf("save to db dayplans with day_id %s tripID %s, order %d", currentDayPlanId, tripID, dayOrder+1)
		// 3.2 save each places of the day
		for visitOrder, place := range planedRoute {
			// 3.2.1 save the place detail if necessary
			placeID := place.Id
			placeIsInDB, err := backend.CheckIfItemExistsInDB(backend.TableName_PlaceDetails, "placeID", placeID)
			if err != nil {
				log.Println("Error during checking if place ID already exists: ", err)
				return plansToShow, err
			}
			if !placeIsInDB {
				err = SavePlaceToDB(place)
				if err != nil {
					log.Fatal("Error during store new trip place: ", err)
					return plansToShow, err
				}
			}
			// 3.2.1 save the day-place relation
			dayPlaceRelationEntry := map[string]interface{}{
				"placeID":    placeID,
				"dayPlanID":  currentDayPlanId,
				"visitOrder": visitOrder + 1,
			}
			err = backend.InsertIntoDB(backend.TableName_DayPlaceRelations, dayPlaceRelationEntry)
			if err != nil {
				log.Fatal("Error during store new day-place relation: ", err)
				return plansToShow, err
			}
		}
	}
	plansToShow, err = ReadAllDayPlansOfTripPlan(tripID)
	if err != nil {
		log.Println("Error reading the plan that is just saved to DB: ", err)
		return plansToShow, err
	}
	return plansToShow, nil
}

func SavePlaceToDB(place model.Place) error {
	// save place
	var photoURLs []string
	for _, p := range place.Photos {
		photoURLs = append(photoURLs, p.Id)
	}
	placeEntry := map[string]interface{}{
		"placeID":   place.Id,
		"name":      place.DisplayName.Text,
		"address":   place.Address,
		"placeType": place.PlaceType.Text,
		"photoURLs": strings.Join(photoURLs, "$$"),
		"longitude": place.Location.Longitude,
		"latitude":  place.Location.Latitude,
	}

	err := backend.InsertIntoDB(backend.TableName_PlaceDetails, placeEntry)
	if err != nil {
		log.Println("Error during store new place: ", err)
		return err
	}
	// save reviews of this place
	for _, review := range place.Reviews {
		reviewEntry := map[string]interface{}{
			"reviewID":    uuid.New(),
			"reviewText":  review.Text.Text,
			"rating":      review.Rating,
			"publishTime": review.PublishTime,
			"placeID":     place.Id,
		}
		err = backend.InsertIntoDB(backend.TableName_Reviews, reviewEntry)
		if err != nil {
			log.Fatal("Error during store new review: ", err)
			return err
		}
	}
	return nil
}
