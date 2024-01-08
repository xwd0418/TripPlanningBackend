package main

import (
	"fmt"
	"tripPlanning/service"
)

// func main() {
// 	fmt.Println("start service")
// 	backend.InitDB()
// 	// backend.InitGCSBackend()
// 	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))

// }
func main() {
	fmt.Println("start service")
	service.GetDefaultPlaces(3)

}
