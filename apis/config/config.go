package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost           string
	DBPort           string
	DBUser           string
	DBPass           string
	DBName           string
	ServerPort       string
	APIGroup         string
	TWILIOApiSecret  string
	TWILIOAccountSID string
	TWILIOAuthToken  string
	IMDBApiKey       string
	IMDBApiUrl       string
}

func LoadConfig() (*Config, error) {

	// Load config from env vars
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return &Config{
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPass:           getEnv("DB_PASSWORD", ""),
		DBName:           getEnv("DB_NAME", "mydatabase"),
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		APIGroup:         getEnv("API_GROUP", "/api/v1"),
		TWILIOApiSecret:  getEnv("TWILIO_API_SECRET", ""),
		TWILIOAccountSID: getEnv("TWILIO_ACCOUNT_SID", ""),
		TWILIOAuthToken:  getEnv("TWILIO_AUTH_TOKEN", ""),
		IMDBApiKey:       getEnv("IMDB_API_KEY", ""),
		IMDBApiUrl:       getEnv("IMDB_API_URL", ""),
	}, nil
}

// getEnv follows Open/Closed Principle (OCP): Allows default values.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
