package main

import (
	"fmt"
	"tripPlanning/service"
)

func main() {
    
    fmt.Println(string("start"))
	// service.SearchPlaces("pizza restaurants")
	json_response, err := service.GetDefaultPlaces(5)
	if err != nil {
        fmt.Println("Error sending request: ", err)
        return 
    }
	fmt.Println((json_response))
}