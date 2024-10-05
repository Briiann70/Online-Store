package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
