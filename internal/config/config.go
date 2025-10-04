package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Username         string
	Password         string
	ConnectionString string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		Username:         getEnv("DB_USERNAME", "default_user"),
		Password:         getEnv("DB_PASSWORD", "default_pass"),
		ConnectionString: getEnv("DB_CONNECTION_STRING", "localhost:5432"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
