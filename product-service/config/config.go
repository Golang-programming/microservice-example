package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

var Cfg Config

func init() {
	loadEnv()
	fmt.Println("DB_URL", os.Getenv("DB_URL"))
	Cfg = Config{
		DBUrl: os.Getenv("DB_URL"),
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
