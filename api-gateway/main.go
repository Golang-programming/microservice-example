package main

import (
	"api-gateway/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Load routes
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":7070")
}
