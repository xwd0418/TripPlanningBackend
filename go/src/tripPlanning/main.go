package main

import (
	"fmt"
	"log"
	"net/http"
	"tripPlanning/backend"
	"tripPlanning/handler"
)

func main(){
	fmt.Println("start service")
	backend.InitDB()
	// backend.InitGCSBackend()
	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))

}