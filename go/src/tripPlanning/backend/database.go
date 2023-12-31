/*
This file provides functions to  connect to DB, init DB tables, insert data to DB,
and delete data from DB
*/
package backend

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"tripPlanning/constants"

	_ "github.com/lib/pq"
)

var (
	db                          *sql.DB
	err                         error
	TableName_Users             = "Users"
	TableName_Trips             = "Trips"
	TableName_DayPlans          = "DayPlans"
	TableName_DayPlaceRelations = "DayPlaceRelations"
	TableName_PlaceDetails      = "PlaceDetails"
	TableName_Reviews           = "Reviews"
)

func InitDB() error {

	// Connection string
	var (
		dbUser    = constants.DB_USER
		dbPwd     = constants.DB_PWD
		dbTCPHost = constants.DB_TCP_HOST // e.g. '127.0.0.1'
		dbPort    = constants.DB_PORT     // e.g. '5432'
		dbName    = constants.DB_NAME     // e.g. 'my-database'
	)
	connStr := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s sslmode=disable",
		dbTCPHost, dbUser, dbPwd, dbPort, dbName)
	// Open the connection

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// defer db.Close() It is commented because we want the DB to keep running
	fmt.Println("successfully connected DB")

	err = initAllTables()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func InsertIntoDB(tableName string, entry map[string]interface{}) error {

	if len(entry) == 0 {
		return fmt.Errorf("entry is empty")
	}

	// Build the SQL statement
	var keys []string
	var placeholders []string
	var values []interface{}

	for k, v := range entry {
		keys = append(keys, k)
		placeholders = append(placeholders, "?")
		values = append(values, v)
	}

	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(keys, ", "),
		strings.Join(placeholders, ", "),
	)

	// Prepare the statement
	prepStmt, err := db.Prepare(stmt)
	if err != nil {
		return err
	}
	defer prepStmt.Close()

	// Execute the statement
	_, err = prepStmt.Exec(values...)
	return err
}

func initAllTables() error {
	// Create user table
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS Users (
        userID TEXT PRIMARY KEY,
        username TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT ,
    );`
	_, err = db.Exec(createUserTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("User table created successfully")

	// Create Trip table
	createTripTableSQL := `CREATE TABLE IF NOT EXISTS Trips (
        tripID TEXT PRIMARY KEY,
        userID TEXT REFERENCES Users(userID),
		tripName TEXT NOT NULL,
		startDay TEXT NOT NULL,
		endDay Text NOT NULL,
		transportation TEXT 
    );`
	_, err = db.Exec(createTripTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Trip table created successfully")

	// Create DayPlan table
	createDayplanTableSQL := `CREATE TABLE IF NOT EXISTS DayPlans (
        dayPlanID TEXT PRIMARY KEY,
        tripID TEXT REFERENCES Trips(tripID),
		dayOrder INT
    );`
	_, err = db.Exec(createDayplanTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("DayPlan table created successfully")

	// Create DayPlaceRelation table
	createDayPlaceRelationsTableSQL := `CREATE TABLE IF NOT EXISTS DayPlaceRelations (
        placeID TEXT PRIMARY KEY,
        dayPlanID TEXT REFERENCES DayPlan(dayPlanID),
		visitOrder INT
    );`
	_, err = db.Exec(createDayPlaceRelationsTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("DayPlaceRelation table created successfully")

	// Create placeDetails table
	createPlaceDetailsTableSQL := `CREATE TABLE IF NOT EXISTS PlaceDetails (
        placeID TEXT PRIMARY KEY,
        name TEXT NOT NULL,
		address TEXT NOT NULL,
		placeType TEXT,
		photoURLs TEXT
    );`
	_, err = db.Exec(createPlaceDetailsTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("PlaceDetails table created successfully")

	// Create Reviews table
	createReviewsTableSQL := `CREATE TABLE IF NOT EXISTS Reviews (
        reviewID TEXT PRIMARY KEY,
		reviewText TEXT,
        rating INT,
		publishTime TEXT,
		placeID TEXT REFERENCES PlaceDetails(placeID)

    );`
	_, err = db.Exec(createReviewsTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("PlaceDetails table created successfully")

	return nil
}
