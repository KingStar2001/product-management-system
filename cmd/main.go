// cmd/main.go
package main

import (
	"product-management-system/internal/api/routes"
	"product-management-system/internal/db"
)

func main() {
	// Set up database connection string (replace with your actual DB details)
	connStr := "user=postgres password=King@2001 dbname=products sslmode=disable"

	// Initialize the database
	db.InitDB(connStr)

	// Start the server
	routes.StartServer()
}
