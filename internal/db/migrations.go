// internal/db/migrations.go
package db

import (
	"fmt"
	"log"
)

// Migrate runs database migrations
func Migrate() {
	// Ensure that DB is already initialized
	if DB == nil {
		log.Fatalf("DB is not initialized")
	}

	// Example query to create the products table
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			product_name TEXT NOT NULL,
			product_description TEXT,
			product_price DECIMAL(10, 2),
			created_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	fmt.Println("Migrations executed successfully!")
}
