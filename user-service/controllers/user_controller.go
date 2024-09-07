package controllers

import (
	"context"
	"net/http"
	"time"
	"user-service/services"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func CreateUser(c *gin.Context) {
	var userDto services.CreateUserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.CreateUser(userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GenerateAPIKey(c *gin.Context) {
	// For simplicity, generate a key (in a real app, use a more secure method)
	apiKey := "key-" + time.Now().Format("20060102150405")

	err := services.RedisClient.Set(ctx, apiKey, "valid", 24*time.Hour).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate API key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"api_key": apiKey})
}
