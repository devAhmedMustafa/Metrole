package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	GRPC_URL             string
	PRODUCTION           bool
	Host                 string
	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
}

var AppSettings *Settings

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	AppSettings = &Settings{
		GRPC_URL:             os.Getenv("GRPC_URL"),
		PRODUCTION:           getEnv("PRODUCTION", "false") == "true",
		Host:                 getEnv("HOST", "http://localhost:8080"),
		GOOGLE_CLIENT_ID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GOOGLE_CLIENT_SECRET: getEnv("GOOGLE_CLIENT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
