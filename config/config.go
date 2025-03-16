package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigLoader interface {
	Load() (*Config, error)
}

type (
	Config struct {
		DB  PostgresConfig
		App AppConfig
	}

	PostgresConfig struct {
		Host string `default:"localhost"`
		Port string `default:"5432"`
		User string `default:"postgres"`
		Pass string `default:"password"`
		Name string `default:"moviebase"`
	}

	AppConfig struct {
		AppName string `default:"moviebase"`
		Port    string `default:"8080"`
		Version string `default:"v1"`
		Prefix  string `default:"api"`
	}
)

// Loads config from .env
// If you want to load the config from some other source
// JCreate new struct and just implement the [ConfigLoader] into that
type EnvConfigLoader struct{}

func (envConfigLoader *EnvConfigLoader) Load() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	cfg.DB.Host = getEnv("DB_HOST", cfg.DB.Host)
	cfg.DB.Port = getEnv("DB_PORT", cfg.DB.Port)
	cfg.DB.User = getEnv("DB_USER", cfg.DB.User)
	cfg.DB.Pass = getEnv("DB_PASSWORD", cfg.DB.Pass)
	cfg.DB.Name = getEnv("DB_NAME", cfg.DB.Name)

	cfg.App.AppName = getEnv("APP_NAME", cfg.App.AppName)
	cfg.App.Port = getEnv("APP_PORT", cfg.App.Port)
	cfg.App.Version = getEnv("APP_VERSION", cfg.App.Version)
	cfg.App.Prefix = getEnv("APP_PREFIX", cfg.App.Prefix)

	return cfg, nil
}

func LoadNewConfig(loader ConfigLoader) (*Config, error) {
	return loader.Load()
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
