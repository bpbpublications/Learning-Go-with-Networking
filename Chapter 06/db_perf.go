package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// Open a database connection
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set maximum idle connections and maximum open connections
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	// Enable statement caching
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	// Query users asynchronously
	results := make(chan []User)
	go getUsersAsync(db, results)

	// Process and print the retrieved users
	users := <-results
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}

func getUsersAsync(db *sql.DB, results chan<- []User) {
	// Prepare the SELECT statement
	stmt, err := db.Prepare("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the statement
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Fetch the users from the result set
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	results <- users
}

}

