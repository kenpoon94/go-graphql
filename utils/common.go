package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env files")
	}
	return os.Getenv(key)
}