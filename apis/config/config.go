package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ConfigLoader is the interface for loading configuration
type ConfigLoader interface {
	Load() *Config
}

// Config holds the configuration values
type Config struct {
	DBType           string
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

// EnvConfigLoader loads configuration from environment variables
type EnvConfigLoader struct{}

// Load loads the configuration from environment variables
func (e *EnvConfigLoader) Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	return &Config{
		DBType:           getEnv("DB_TYPE", "postgres"),
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
	}
}

// getEnv follows Open/Closed Principle (OCP): Allows default values.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// LoadConfig loads the configuration using the provided loader
func LoadEnvConfig(loader ConfigLoader) *Config {
	return loader.Load()
}
