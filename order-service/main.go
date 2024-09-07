package main

import (
	"order-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Load routes
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8083")
}
