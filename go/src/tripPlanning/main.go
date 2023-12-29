package main

import (
	"fmt"
	"tripPlanning/backend"
)

func main(){
	fmt.Println("start service")
	backend.InitDB()
	// backend.InitGCSBackend()
	// log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))

}