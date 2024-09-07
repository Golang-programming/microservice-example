package middlewares

import (
	"context"
	"net/http"
	"product-service/services"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("API-Key")

		// Check API key in Redis
		val, err := services.RedisClient.Get(ctx, apiKey).Result()
		if err != nil || val != "valid" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
