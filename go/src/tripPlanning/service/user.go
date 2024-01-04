/*
Managing user signup and login functionalities
*/
package service

import (
	"fmt"

	"tripPlanning/backend"
	"tripPlanning/model"

	"github.com/google/uuid"
)

func AddUser(user *model.User) (bool, error) {
    // Check if the username already exists
    existingUser, err := backend.GetUser(user.Username)
    if err != nil {
        return false, fmt.Errorf("error checking for existing user: %v", err)
    }
    if existingUser != nil {
        return false, fmt.Errorf("username '%s' already exists", user.Username)
    }

    // Save the new user
    //TODO: generate id 
    user.Id = uuid.New().String()
    err = backend.SaveUser(user)
    if err != nil {
        return false, fmt.Errorf("error saving new user: %v", err)
    }

    fmt.Printf("User added: %s\n", user.Username)
    return true, nil
}


func CheckUser(username, password string) (bool, error) { //Checkuser
    user, err := backend.GetUser(username)
    if err != nil {
        return false, fmt.Errorf("error retrieving user: %v", err)
    }
    if user == nil {
        return false, fmt.Errorf("username '%s' does not exist", username)
    }

    // Check if the provided password matches
    // Assuming passwords are stored in plain text 
    // TODO: encode password
    if user.Password != password {
        return false, fmt.Errorf("incorrect password for username '%s'", username)
    }

    // User is authorized
    return true, nil
}

