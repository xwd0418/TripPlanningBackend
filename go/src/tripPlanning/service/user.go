/*
Managing user signup and login functionalities
*/
package service

import (
	"fmt"
	"log"
	"reflect"

	"database/sql"
	"tripPlanning/backend"
	"tripPlanning/constants"
	"tripPlanning/model"

	"github.com/olivere/elastic/v7"
	// "github.com/olivere/elastic/v7"
)

func AddUser(user *model.User) (bool, error) {
    // Check if the username already exists
    existingUser, err := backend.service.GetUser(user.Username)
    if err != nil {
        return false, fmt.Errorf("error checking for existing user: %v", err)
    }
    if existingUser != nil {
        return false, fmt.Errorf("username '%s' already exists", user.Username)
    }

    // Save the new user
    err = backend.service.SaveUser(user)
    if err != nil {
        return false, fmt.Errorf("error saving new user: %v", err)
    }

    fmt.Printf("User added: %s\n", user.Username)
    return true, nil
}

// method 1: search ES using username then compare password
// method 2: search ES using username + passowrd, totalhits() > -> success
func CheckUser(username, password string) (bool, error) { //Checkuser
    query := elastic.NewBoolQuery() // Creates a new bool query.
    query.Must(elastic.NewTermQuery("username", username)) // must contain this username
    query.Must(elastic.NewTermQuery("password", password)) // must contain this password
    searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
    if err != nil {
        return false, err
    }
    return false, nil
}

