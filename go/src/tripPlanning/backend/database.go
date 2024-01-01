package backend

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"

    "tripPlanning/model"

	_ "github.com/lib/pq"
)

var  (
	db *sql.DB
	err error
)

func Start_Cloud_SQL_Auth_Proxy(){
	cmd := exec.Command("./cloud-sql-proxy", "--private-ip", "tripplanning-409112:us-west1:trip-plan-database")
	_, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
    }
}

func InitDB() {
	
    // Connection string
	var (
		dbUser    = "postgres"
		dbPwd     = "David418."   
		dbTCPHost = "10.120.64.3" 
		dbPort    = "5432"     // e.g. '5432'
		dbName    = "TripPlanningDatabase"       // e.g. 'my-database'
)
    connStr := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s sslmode=disable",
                dbTCPHost, dbUser, dbPwd, dbPort, dbName)
	// fmt.Println(connStr)
    // Open the connection
	
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create table SQL statement
    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username TEXT NOT NULL,
		password TEXT NOT NULL
    );`

    // Execute the SQL statement
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Table created successfully")
}

type DatabaseService struct {
    DB *sql.DB
}

// 
func (service *DatabaseService) GetUser(username string) (*model.User, error) {
    var user model.User
    query := `SELECT id, username, password FROM users WHERE username = $1`
    err := service.DB.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // No user found
        }
        return nil, fmt.Errorf("error querying user: %w", err)
    }
    return &user, nil
}

// SaveNewUser saves a new user with the given username to the database
func (service *DatabaseService) SaveUser(user *model.User) error {
    // SQL statement to insert a new user
    query := `INSERT INTO users (username, password) VALUES ($1, $2)`
    
    _, err := service.DB.Exec(query, user.Username, user.Password)
    if err != nil {
        return fmt.Errorf("error saving user: %w", err)
    }

    return nil
}


