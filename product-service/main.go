package main

import (
	"product-service/database"
	"product-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()
	router := gin.Default()

	// Load routes
	routes.RegisterRoutes(router, db)

	// Start the server
	router.Run(":8082")
}
