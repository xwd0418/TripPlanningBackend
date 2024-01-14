/*
	 Generate user's plans:
	 	given a list of places user wants to visit within a day
		return a list of reordered places that optimizes travel distance
*/
package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strings"
	"tripPlanning/constants"
	"tripPlanning/model"
)

func GenerateDayPlan(places []model.Place, transportation string, date string) ([]model.Place, error) {
	// Step 1: Create a matrix of distances between all places
    distanceMatrix, err := GetDistanceMatrix(places, transportation)
    if err != nil {
        return nil, err
    }

	/// Step 2: Find the shortest route
    shortestRouteIndices, err := FindShortestRoute(distanceMatrix)
    if err != nil {
        return nil, err
    }

	// Step 3: Reorder the places according to the shortest route
    reorderedPlaces := make([]model.Place, len(places))
    for i, idx := range shortestRouteIndices {
        reorderedPlaces[i] = places[idx]
    }

	return places, nil
}

// Use Google Maps Distance Matrix API to get distances between places
func GetDistanceMatrix(places []model.Place, transportation string) ([][]int, error) {
    // Prepare the origins and destinations in the required format
    var origins, destinations []string
    for _, place := range places {
        formattedAddress := url.QueryEscape(place.Address)
        origins = append(origins, formattedAddress)
        destinations = append(destinations, formattedAddress)
    }

    // Construct the API request URL
    apiUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%s&destinations=%s&mode=%s&key=%s",
        strings.Join(origins, "|"),
        strings.Join(destinations, "|"),
        transportation,
        constants.GOOGLE_MAP_API_KEY)

    // Make the request
    resp, err := http.Get(apiUrl)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Read and parse the response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    // Print the raw JSON response
    fmt.Println("Raw JSON response:", string(body))

	var matrixResponse model.DistanceMatrixResponse
    if err := json.Unmarshal(body, &matrixResponse); err != nil {
        return nil, err
    }

	// Check if response status is OK
    if matrixResponse.Status != "OK" {
        return nil, errors.New("API response status is not OK")
    }

	// Extract the distance matrix
    numOrigins := len(matrixResponse.OriginAddresses)
    numDestinations := len(matrixResponse.DestinationAddresses)
    distanceMatrix := make([][]int, numOrigins)

    // fmt.Println("numOrigins:", numOrigins)
    // fmt.Println("len(matrixResponse.Rows):", len(matrixResponse.Rows))
    // fmt.Println("numDestinations:", numDestinations)

    if numOrigins != len(matrixResponse.Rows) {
        return nil, errors.New("mismatch between number of origins and number of rows in response")
    }

    if numOrigins == 0 {
        return nil, errors.New("no origins provided or received in response")
    }    
    

    for i, row := range matrixResponse.Rows {
        distanceMatrix[i] = make([]int, numDestinations)
        for j, element := range row.Elements {
            if element.Status != "OK" {
                return nil, fmt.Errorf("Error with element status: %s", element.Status)
            }
            distanceMatrix[i][j] = element.Distance.Value
        }
    }
    
    return distanceMatrix, nil
}

// Function to find the shortest route
func FindShortestRoute(matrix [][]int) ([]int, error) {
    n := len(matrix)
    if n == 0 {
        return nil, errors.New("distance matrix is empty")
    }

    // Create a slice of destination indices
    destinations := make([]int, n)
    for i := range destinations {
        destinations[i] = i
    }

    // Generate all permutations of destinations
    permutations := permute(destinations)

    minDistance := math.MaxInt32
    var shortestRoute []int

    // Iterate over all permutations to find the shortest route
    for _, perm := range permutations {
        distance := calculateDistance(perm, matrix)
        if distance < minDistance {
            minDistance = distance
            shortestRoute = perm
        }
    }

	if len(shortestRoute) == 0 {
        return nil, errors.New("no route found")
    }

    return shortestRoute, nil
}

// Function to generate all permutations of a slice
func permute(nums []int) [][]int {
    var result [][]int
    var backtrack func(first int)
    backtrack = func(first int) {
        if first == len(nums) {
            temp := make([]int, len(nums))
            copy(temp, nums)
            result = append(result, temp)
        }
        for i := first; i < len(nums); i++ {
            nums[i], nums[first] = nums[first], nums[i]
            backtrack(first + 1)
            nums[i], nums[first] = nums[first], nums[i] // backtrack
        }
    }
    backtrack(0)
    return result
}

// Function to calculate the total distance of a route
func calculateDistance(route []int, matrix [][]int) int {
    totalDistance := 0
    for i := 0; i < len(route)-1; i++ {
        totalDistance += matrix[route[i]][route[i+1]]
    }
    return totalDistance
}
