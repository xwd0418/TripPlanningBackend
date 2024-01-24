/*
Managing user signup and login functionalities
*/
package service

import (
	"errors"
	"fmt"

	"tripPlanning/backend"
	"tripPlanning/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user model.User) error {
	// Check if the user with the same email already exists
	existingUserByEmail, err := GetUser(user.Email,true)
	if err != nil {
		return fmt.Errorf("error checking if user with the same email exists: %w", err)
	}

	// If user with the same email exists, return an error
	if existingUserByEmail != nil {
		return errors.New("user with the same email already exists")
	}

	// Check if the user with the same username already exists
	existingUserByUsername, err := GetUser(user.Username,false)
	if err != nil {
		return fmt.Errorf("error checking if user with the same username exists: %w", err)
	}

	// If user with the same username exists, return an error
	if existingUserByUsername != nil {
		return errors.New("user with the same username already exists")
	}

	// Generate hashed password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		fmt.Printf("Error hashing password: %v\n", err)
        return errors.New("Failed to register user")
	}

	// Save user with hashed password to database
    user.Id = uuid.New().String()
	user.HashedPassword = hashedPassword
    userData := map[string]interface{}{
		"userID":   user.Id,
        "username": user.Username,
		"password": user.HashedPassword,
		"email":    user.Email,
	}
	err = backend.InsertIntoDB("Users", userData)
	if err != nil {
		return fmt.Errorf("error saving user in DB: %w", err)
	}

	return nil
}

func hashPassword(password string) (string, error) {
	// Generate hashed password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckUser(identifier, password string, byEmail bool) (bool, error) {
	// Get user from the database by username or email
	user, err := GetUser(identifier, byEmail)
	if err != nil {
		return false, fmt.Errorf("error retrieving user: %v", err)
	}
	if user == nil {
		var identifierType string
		if byEmail {
			identifierType = "email"
		} else {
			identifierType = "username"
		}
		return false, fmt.Errorf("%s '%s' does not exist", identifierType, identifier)
	}

	// Compare hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return false, fmt.Errorf("passwords do not match")
	}

	// Passwords match
	return true, nil
}


func GetUser(identifier string, byEmail bool) (*model.User, error) {
	var user model.User

	// Specify the columns to read
	columnsToRead := []string{"userID", "username", "password", "email"}

	// Specify the condition for the query
	var conditions string
	var value interface{} // Define a value interface

	if byEmail {
		conditions = "email = $1"
		value = identifier
	} else {
		conditions = "username = $1"
		value = identifier 
	}

	// Call ReadFromDB to execute the query
	rows, err := backend.ReadFromDB_user("Users", columnsToRead, conditions, value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Extract data from the result set
	if rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.HashedPassword, &user.Email); err != nil {
			return nil, fmt.Errorf("error scanning user data: %w", err)
		}
        // Print the user information for debugging or logging
		fmt.Printf("User in GetUser: %+v\n", user)
	} else {
		// No user found
		return nil, nil
	}

	return &user, nil
}









