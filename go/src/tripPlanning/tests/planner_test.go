package service_test

import (
	"reflect"
	"testing"
	"tripPlanning/model"
	"tripPlanning/service"

	"github.com/jarcoal/httpmock"
)

// func TestFindShortestRoute(t *testing.T) {
//     // Example distance matrix (replace with actual values)
//     matrix := [][]int{
//         {0, 2, 9},
//         {1, 0, 6},
//         {7, 3, 0},
//     }

//     expectedRoute := []int{2, 1, 0} // Expected shortest route
//     route, err := service.FindShortestRoute(matrix)
//     if err != nil {
//         t.Errorf("findShortestRoute returned an error: %v", err)
//     }
//     if !reflect.DeepEqual(route, expectedRoute) {
//         t.Errorf("findShortestRoute = %v; want %v", route, expectedRoute)
//     }
// }

func TestGetDistanceMatrix(t *testing.T) {
	// Activate the httpmock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the expected request
	httpmock.RegisterResponder("GET", "https://maps.googleapis.com/maps/api/distancematrix/json",
		httpmock.NewStringResponder(200, `{
        "status": "OK",
        "origin_addresses": ["Central Park, New York, USA", "Brooklyn Bridge, New York, USA", "New York University, New York, USA"],
        "destination_addresses": ["Central Park, New York, USA", "Brooklyn Bridge, New York, USA", "New York University, New York, USA"],
        "rows": [
            {
                "elements": [
                    {
                        "status": "OK",
                        "distance": {"text": "0 km", "value": 0},
                        "duration": {"text": "0 mins", "value": 0}
                    },
                    {
                        "status": "OK",
                        "distance": {"text": "2 km", "value": 2000},
                        "duration": {"text": "5 mins", "value": 300}
                    },
                    {
                        "status": "OK",
                        "distance": {"text": "3 km", "value": 3000},
                        "duration": {"text": "10 mins", "value": 600}
                    }
                ]
            },
            {
                "elements": [
                    {
                        "status": "OK",
                        "distance": {"text": "2 km", "value": 2000},
                        "duration": {"text": "5 mins", "value": 300}
                    },
                    {
                        "status": "OK",
                        "distance": {"text": "0 km", "value": 0},
                        "duration": {"text": "0 mins", "value": 0}
                    },
                    {
                        "status": "OK",
                        "distance": {"text": "4 km", "value": 4000},
                        "duration": {"text": "15 mins", "value": 900}
                    }
                ]
            },
            {
                "elements": [
                    {
                        "status": "OK",
                        "distance": {"text": "3 km", "value": 3000},
                        "duration": {"text": "10 mins", "value": 600}
                    },
                    {
                        "status": "OK",
                        "distance": {"text": "4 km", "value": 4000},
                        "duration": {"text": "15 mins", "value": 900}
                    },
                    {
                        "status": "OK",
                        "distance": {"text": "0 km", "value": 0},
                        "duration": {"text": "0 mins", "value": 0}
                    }
                ]
            }
        ]
    }`))



    // Define your input
    places := []model.Place{
        {DisplayName: model.Text{Text: "Central Park"}, Address: "New York, USA"},
        {DisplayName: model.Text{Text: "Brooklyn Bridge"}, Address: "New York, USA"},
        {DisplayName: model.Text{Text: "New York University"}, Address: "New York, USA"},
        // ... other places
    }
    transportation := "driving"

	// Call your function
	matrix, err := service.GetDistanceMatrix(places, transportation)

	// Check for errors
	if err != nil {
		t.Fatalf("YourMatrixFunction returned an unexpected error: %v", err)
	}

	// Define what you expect the matrix to look like
	expectedMatrix := [][]int{
		{0, 100, 200},
		{100, 0, 200},
		{100, 200, 0},
	}

	// Check if the matrix matches your expectations
	if !reflect.DeepEqual(matrix, expectedMatrix) {
		t.Errorf("YourMatrixFunction = %v, want %v", matrix, expectedMatrix)
	}

}
