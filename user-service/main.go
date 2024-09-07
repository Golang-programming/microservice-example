package main

import (
	"user-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Load routes
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8081")
}
