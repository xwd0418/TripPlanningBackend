package model

type TripPlan struct {
	TripPlanId      string    `json:"tripID"`
	UserID          string    `json:"-"`
	TripName        string    `json:"TripName"`
	StartDay        string    `json:"StartDay"`
	EndDay          string    `json:"EndDay"`
	Transportation  string    `json:"Transportation"`
	DayPlans        []DayPlan `json:"-"`
	SamplePlaceName string    `json:"SamplePlaceName"`
}

type DayPlan struct {
	TripPlanId     string  `json:"-"`
	DayPlanID      string  `json:"-"`
	OrderInTrip    int     `json:"-"`
	PlacesToVisit  []Place `json:"places"`
	Date           string  `json:"date"`
	Transportation string  `json:"Transportation"`
}
