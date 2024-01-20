package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"tripPlanning/model"
	"tripPlanning/service"
)

//This handler is used handle any user requests to 'save' a place
//might need to implement the jwt authentication for saveHandler later

func saveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one 'save' request")

	//Note: the frontend will need to make sure the http.Request body adheres to the expected format/structure
	//which strictly follows the model.Post struct
	dec := json.NewDecoder(r.Body)

	var p model.Place

	if err := dec.Decode(&p); err != nil {
		http.Error(w, "Failed to decode the place struct from the request body", http.StatusBadRequest)
		fmt.Printf("Failed to decode the place struct from the request body: %v\n", err)
		return
	}

	err := service.SavePlaceToDB(p)
	if err != nil {
		http.Error(w, "Failed to save place to backend", http.StatusInternalServerError)
		fmt.Printf("Failed to save to backend %v\n", err)
		return
	}

	fmt.Println("The place is saved successfully.")
}
