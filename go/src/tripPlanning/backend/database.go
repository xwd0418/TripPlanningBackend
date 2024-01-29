/*
This file provides functions to  connect to DB, init DB tables, insert data to DB,
and delete data from DB
*/
package backend

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"tripPlanning/constants"

	"tripPlanning/model"

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

func InsertIntoDB(tableName string, entry map[string]interface{}, additional_query_config ...string) error {

	if len(entry) == 0 {
		return fmt.Errorf("entry is empty")
	}

	// Build the SQL statement
	var keys []string
	var placeholders []string
	var values []interface{}

	counter := 1
	for k, v := range entry {
		keys = append(keys, k)
		placeholders = append(placeholders, "$"+strconv.Itoa(counter))
		values = append(values, v)
		counter += 1
	}

	suffix := ""
	if len(additional_query_config) > 0 {
		suffix = additional_query_config[0]
	}

	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) %s",
		tableName,
		strings.Join(keys, ", "),
		strings.Join(placeholders, ", "),
		suffix,
	)

	// Prepare the statement
	prepStmt, err := db.Prepare(stmt)
	if err != nil {
		log.Println("db statement prepare failed, statement is ", stmt, err)
		return err
	}
	defer prepStmt.Close()

	// Execute the statement
	_, err = prepStmt.Exec(values...)
	return err
}

func ReadRowFromDB(query string) *sql.Row {
	// log.Printf("Row query statement is %s", query)
	return db.QueryRow(query)
}

func ReadFromDB(tableName string, columns_to_read []string, conditions string) (*sql.Rows, error) {
	queryStatement := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns_to_read, ", "), tableName)
	if conditions != "" {
		queryStatement += " WHERE " + conditions
	}
	// log.Printf("query statement is %s", queryStatement)
	rows, err := db.Query(queryStatement)
	if err != nil {
		log.Println("Query "+queryStatement+"fails: ", err)
		return nil, err
	}
	return rows, nil
}

func ReadFromDB_user(tableName string, columnsToRead []string, conditions string, values ...interface{}) (*sql.Rows, error) {
	queryStatement := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columnsToRead, ", "), tableName)
	if conditions != "" {
		queryStatement += " WHERE " + conditions
	}
	rows, err := db.Query(queryStatement, values...)
	if err != nil {
		log.Println("Query "+queryStatement+" fails: ", err)
		return nil, err
	}
	return rows, nil
}

func QueryRowsFromDB(queryStatement string) (*sql.Rows, error) {
	rows, err := db.Query(queryStatement)
	if err != nil {
		log.Println("Query "+queryStatement+"fails: ", err)
		return nil, err
	}
	return rows, nil
}

// CheckIfItemExistsInDB checks if an item exists in a specified column in a table
func CheckIfItemExistsInDB(tableName string, columnName string, itemValue interface{}) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = $1)", tableName, columnName)
	// log.Printf("check duplicated place query is %s", query)
	// log.Println(itemValue)
	err := db.QueryRow(query, itemValue).Scan(&exists)
	if err != nil {
		log.Printf("error of query row during checking duplicated places %v", err)
		return false, err
	}
	return exists, nil
}

func initAllTables() error {
	// Create user table
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS Users (
        userID TEXT PRIMARY KEY,
        username TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT
    );`
	_, err = db.Exec(createUserTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("User table created successfully or already exists")

	// Create Trip table
	createTripTableSQL := `CREATE TABLE IF NOT EXISTS Trips (
        tripID TEXT PRIMARY KEY,
        userID TEXT REFERENCES Users(userID),
		tripName TEXT NOT NULL,
		startDay TEXT NOT NULL,
		endDay Text NOT NULL,
		transportation TEXT ,
		SamplePlaceName TEXT
    );`
	_, err = db.Exec(createTripTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Trip table created successfully or already exists")

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
	fmt.Println("DayPlan table created successfully or already exists")

	// Create placeDetails table
	createPlaceDetailsTableSQL := `CREATE TABLE IF NOT EXISTS PlaceDetails (
        placeID TEXT PRIMARY KEY,
        name TEXT NOT NULL,
		address TEXT NOT NULL,
		placeType TEXT,
		photoURLs TEXT,
		longitude DECIMAL,
		latitude DECIMAL
    );`
	_, err = db.Exec(createPlaceDetailsTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("PlaceDetails table created successfully or already exists")

	// Create DayPlaceRelation table
	createDayPlaceRelationsTableSQL := `CREATE TABLE IF NOT EXISTS DayPlaceRelations (
        placeID TEXT REFERENCES PlaceDetails(placeID),
        dayPlanID TEXT REFERENCES DayPlans(dayPlanID),
		visitOrder INT,
		PRIMARY KEY (placeID, dayPlanID)
    );`
	_, err = db.Exec(createDayPlaceRelationsTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("DayPlaceRelation table created successfully or already exists")

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
	fmt.Println("PlaceDetails table created successfully or already exists")

	return nil
}

// F=GetUser returns model.User object from table Users with given username
func GetUser(username string) (*model.User, error) {
	var user model.User
	query := `SELECT userID, username, password FROM Users WHERE username = $1`
	err := db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("error querying user: %w", err)
	}
	log.Println("inside getting user function, user is", user.Username, user.Id)
	return &user, nil
}

// SaveUser saves a new user to table Users
func SaveUser(user *model.User) error {
	// SQL statement to insert a new user
	query := `INSERT INTO Users (username, password, userID, email) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, user.Username, user.Password, user.Id, user.Email)
	if err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}

	return nil
}

// DeleteFromDB
func DeleteFromDB(tableName, conditionColumn string, conditionValue interface{}) error {
	// Build the SQL statement
	stmt := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", tableName, conditionColumn)

	// Prepare the statement
	prepStmt, err := db.Prepare(stmt)
	if err != nil {
		log.Println("DB statement prepare failed, statement is ", stmt, err)
		return err
	}
	defer prepStmt.Close()

	// Execute the statement
	_, err = prepStmt.Exec(conditionValue)
	if err != nil {
		log.Println("Error deleting from table", tableName, ":", err)
		return err
	}

	return nil
}

// UpdateToDB
func UpdateToDB(tableName string, idField string, idValue interface{}, updatedFields map[string]interface{}) error {
	// Build the SQL statement for update
	var placeholders []string
	var values []interface{}

	//placeholders -> the fields that need to be updated in the SQL statement
	//values -> new values
	for k, v := range updatedFields {
		placeholders = append(placeholders, fmt.Sprintf("%s=$%d", k, len(values)+1))
		values = append(values, v)
	}

	stmt := fmt.Sprintf("UPDATE %s SET %s WHERE %s=$%d",
		tableName,
		strings.Join(placeholders, ", "),
		idField,
		len(values)+1,
	)

	// Prepare the statement
	prepStmt, err := db.Prepare(stmt)
	if err != nil {
		log.Println("db statement prepare failed, statement is ", stmt, err)
		return err
	}
	defer prepStmt.Close()

	// Execute the update statement
	_, err = prepStmt.Exec(values...)
	return err
}
