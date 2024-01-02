package model

type TripPlan struct {
	TripPlanId     string    `json:"-"`
	UserID         string    `json:"-"`
	TripName       string    `json:"TripName"`
	StartDay       string    `json:"StartDay"`
	EndDay         string    `json:"EndDay"`
	Transportation string    `json:"Transportation"`
	DayPlans       []DayPlan `json:"-"`
	City           string    `json:"City"`
}

type DayPlan struct {
	TripPlanId string
	DayPlanID  string
	// OrderInTrip int
	PlacesToVisit []Place
}
