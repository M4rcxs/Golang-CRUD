// Package middleware is responsible for handling requests and responses
// between the server and database.
package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// The response struct defines the shape of the response data that will be
// sent back to the client as JSON.
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message, omitempty"`
}

// createConnection() establishes a connection to the PostgreSQL database.
// It loads environment variables from the .ENV file, initializes the
// database driver, and pings the database to ensure connectivity.
func createConnection() *sql.DB {
	// Load environment variables from .ENV file
	err := godotenv.Load(".ENV")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize database driver
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// Ping the database to ensure connectivity
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// Return the database connection
	return db
}
