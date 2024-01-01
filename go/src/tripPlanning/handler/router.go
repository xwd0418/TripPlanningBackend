package handler

import (
	"net/http"
	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	// This module lets you authenticate HTTP requests using JWT tokens in your Go Programming Language applications.
	jwt "github.com/form3tech-oss/jwt-go" // Package jwt is a Go implementation of JSON Web Tokens:
	// JWT json web token

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	jwtMiddleware := jwtMiddleware.New(jwtMiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            return []byte(mySigningKey), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })
	router := mux.NewRouter()
	router.Handle("/showDefaultPlaces", http.HandlerFunc(showDefaultPlacesHandler)).Methods("GET")

	 // when fisrt sign in and sign up, no token authentication
	 router.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST")
	 router.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST")
	 
	return router
}
