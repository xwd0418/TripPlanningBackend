package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"tripPlanning/model"
	"tripPlanning/service"

	jwt "github.com/form3tech-oss/jwt-go" // package is a Go implementation of JSON Web Tokens
)

var mySigningKey = []byte("secret")

// using input username and password to check if user exists and signin successful, if successful then generate token and return token to client
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one signin request")
	w.Header().Set("Content-Type", "text/plain")

	//  Get User information from client
	// process request -> user
	decoder := json.NewDecoder(r.Body) // NewDecoder returns a new decoder that reads from r.
	var user model.User                // define user type
	if err := decoder.Decode(&user); err != nil {
		// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest) // this is returned to client
		fmt.Printf("Cannot decode user data from client %v\n", err)                 // this is returned on service side
		return
	}

	// check if user exists,

	success, err := service.CheckUser(user.Username, user.Password,false)
	// if yes, return true, otherwise false,
	// if there is error message, something else is wrong
	if err != nil {
		http.Error(w, "Failed to read user from Database", http.StatusInternalServerError)
		fmt.Printf("Failed to read user from Database %v\n", err)
		return
	}

	if !success {
		http.Error(w, "User doesn't exists or wrong password", http.StatusUnauthorized)
		fmt.Printf("User doesn't exists or wrong password\n")
		return
	}

	// 2.1 if log in successfully first time, generate token for future use
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id": user.Id,
		// we didn't include password in this file because this token can be reversed, and it's not safe to inlude password
		"exp": time.Now().Add(time.Hour * 24).Unix(), // experition date
	})

	tokenString, err := token.SignedString(mySigningKey) // Get the complete, signed token
	//  if something is wrong
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		fmt.Printf("Failed to generate token %v\n", err)
		return
	}

	// 3. Generate response with token and user ID, return token to client
	// response := struct {
	// 	Token  string `json:"token"`
	// 	UserID string `json:"id"`
	// }{
	// 	Token:  tokenString,
	// 	UserID: user.Id,
	// }

	// responseData, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, "Failed to generate response data", http.StatusInternalServerError)
	// 	fmt.Printf("Failed to generate response data %v\n", err)
	// 	return
	// }

	// w.Write(responseData)

	w.Write([]byte(tokenString))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one signup request")
	w.Header().Set("Content-Type", "text/plain")

	// process request -> user
	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode user data from client %v\n", err)
		return
	}

	if user.Username == "" || user.Password == "" || regexp.MustCompile(`^[a-z0-9]$`).MatchString(user.Username) { // MatchString reports whether the string s contains any match of the regular expression re.
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		fmt.Printf("Invalid username or password\n")
		return
	}
	
	err := service.RegisterUser(user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		fmt.Printf("Failed to register user %v\n", err)
		return
	}

	fmt.Printf("User added successfully: %s.\n", user.Username)
}
