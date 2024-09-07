package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserServiceURL    string
	ProductServiceURL string
	OrderServiceURL   string
}

var Cfg Config

func init() {
	loadEnv()
	fmt.Println("USER_SERVICE_URL", os.Getenv("USER_SERVICE_URL"))
	Cfg = Config{
		UserServiceURL:    os.Getenv("USER_SERVICE_URL"),
		ProductServiceURL: os.Getenv("PRODUCT_SERVICE_URL"),
		OrderServiceURL:   os.Getenv("ORDER_SERVICE_URL"),
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
