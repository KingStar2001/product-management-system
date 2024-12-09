package main

import (
	"log"
	"product-management-system/internal/api/routes"
)

func main() {
	// Initialize routes
	router := routes.SetupRouter()

	// Start the server
	log.Println("Server running on port 8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
