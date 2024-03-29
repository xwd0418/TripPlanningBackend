package handler

import (
	"fmt"
	"net/http"

	// jwtMiddleware "github.com/auth0/go-jwt-middleware"

	// // This module lets you authenticate HTTP requests using JWT tokens in your Go Programming Language applications.
	// jwt "github.com/form3tech-oss/jwt-go" // Package jwt is a Go implementation of JSON Web Tokens:
	// JWT json web token
	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {
	// jwtMiddleware := jwtMiddleware.New(jwtMiddleware.Options{
	//     ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	//         return []byte(mySigningKey), nil
	//     },
	//     SigningMethod: jwt.SigningMethodHS256,
	// })

	router := mux.NewRouter()

	// places routing
	router.Handle("/showDefaultPlaces", http.HandlerFunc(showDefaultPlacesHandler)).Methods("GET")
	router.Handle("/searchPlaces", http.HandlerFunc(searchPlacesPlacesHandler)).Methods("GET")

	//  DB loading routing
	router.Handle("/getAllPlansOfUser", http.HandlerFunc(readUserGeneralTripsHandler)).Methods("GET")
	router.Handle("/getTripInfo", http.HandlerFunc(readAllDayPlansOfTripPlanHandler)).Methods("GET")

	// DB saving routing yc commented for testing
	router.Handle("/generateTripPlan", http.HandlerFunc(GeneratePlanAndSaveHandler)).Methods("POST")
	// NEED CHANGE: HOW TO SEND INPUT AS JSON

	//savePlaces could be a "put" OR potentially a "post" request
	// router.Handle("/savePlace", http.HandleFunc(saveHandler)).Methods("PUT")

	// New delete route for a trip
	router.Handle("/deleteTrip", http.HandlerFunc(DeleteTripHandler)).Methods("DELETE")

	// modify routing
	router.Handle("/modifyTrip", http.HandlerFunc(modifyTripHandler)).Methods("POST")

	// generate/save plan with visit order exactly as user requested
	router.Handle("/generateExactTrip", http.HandlerFunc(generateExactTripHandler)).Methods("POST")
	fmt.Println("ready to receive requests")

	// when fisrt sign in and sign up, no token authentication
	router.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST")
	router.Handle("/login", http.HandlerFunc(loginHandler)).Methods("POST")
	router.Handle("/recommendation", http.HandlerFunc(AiGeneratedPlanHandler)).Methods("GET")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})

	return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}
