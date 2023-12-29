package model

type TripPlan  struct{
	TripPlanId string
	UserID string 
	StartDay string
	EndDay string
	Transportation string
	DayPlans []DayPlan
}

type DayPlan struct{
	TripPlanId string
	DayPlanID string
	// OrderInTrip int
	PlacesToVisit []Place
}



