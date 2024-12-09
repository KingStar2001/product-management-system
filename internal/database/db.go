// internal/db/db.go
package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Global DB variable
var DB *sql.DB

// InitDB initializes the database connection
func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	fmt.Println("Database connected successfully")

	// Create schema if it does not exist
	// CreateSchema()
}

// CreateSchema creates the necessary tables if they don't exist
// func CreateSchema() {
// 	_, err := DB.Exec(`
// 		CREATE TABLE IF NOT EXISTS products (
// 			id SERIAL PRIMARY KEY,
// 			product_name TEXT NOT NULL,
// 			product_description TEXT,
// 			product_price DECIMAL(10, 2),
// 			created_at TIMESTAMP DEFAULT NOW()
// 		)
// 	`)
// 	if err != nil {
// 		log.Fatalf("Error creating schema: %v", err)
// 	}

// 	fmt.Println("Schema created or already exists.")
// }
