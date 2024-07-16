package main

import (
	"taskManagement/database"
	"taskManagement/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	database.Connect()

	// Set up Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Run the server
	router.Run(":8080")
}
