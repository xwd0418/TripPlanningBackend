/* search places for user */
package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"tripPlanning/constants"
	"tripPlanning/model"
)

// This gives a list of place infomations based on user's search.
// The infomation contains id, Name, address,photos_urls,places.reviews")
func SearchPlaces(searchQuery string, maxOutputNum int)  ([]model.Place, error) {
	// TODO: first check if this has been searched before, a cache will be used here

	//  if not found in cache, call google map API to get the places in json format
	var jsonData = []byte(fmt.Sprintf(`{
		"textQuery": "%s in %s" 
	}`, searchQuery, constants.CITY))

    url := "https://places.googleapis.com/v1/places:searchText"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error creating request: ", err)
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Goog-Api-Key", constants.GOOGLE_MAP_API_KEY) 
    req.Header.Set("X-Goog-FieldMask", "places.id,places.displayName,places.formattedAddress,places.photos,places.reviews")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request: ", err)
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading body: ", err)
        return nil, err
    }

	fmt.Println(string(body))

	// convert the json to map, and then get the places objects 
	// 1. Unmarshal into an interface{}
    var result map[string]interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        log.Fatal(err)
    }
	// 2. Extract the sub-JSON (e.g., the "address" part)
	places, ok := result["places"].([]interface{})
    if !ok {
        log.Fatal("Error extracting sub-JSON (list of places)")
    }
	// 3. Marshal the sub-JSON back into []byte if needed
    placesJsonData, err := json.Marshal(places)
    if err != nil {
        log.Fatal(err)
    }

	var searchedPlaces []model.Place
	err = json.Unmarshal(placesJsonData, &searchedPlaces)
    if err != nil {
        log.Fatal(err)
    }
	
	return searchedPlaces[:min(maxOutputNum, len(searchedPlaces))], nil

}
	

// default list for users to select
func GetDefaultPlaces(maxOutputNum int)  ([]model.Place, error) {
	return SearchPlaces("tourist attractions", maxOutputNum) 
}