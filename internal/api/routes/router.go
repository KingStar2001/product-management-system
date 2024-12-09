package routes

import (
	"product-management-system/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Product Management System API!",
		})
	})

	router.POST("/products", handlers.CreateProductHandler)
	router.GET("/products/:id", handlers.GetProductHandler)
	router.GET("/products", handlers.GetProductsHandler)

	return router
}
func StartServer() {
	r := gin.Default()

	// Define routes here
	r.POST("/products", handlers.CreateProductHandler)
	r.GET("/products/:id", handlers.GetProductHandler)
	r.GET("/products", handlers.GetProductsHandler)

	// Start the server on port 8080
	r.Run(":8080")
}
