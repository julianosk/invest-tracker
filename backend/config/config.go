package config

import (
	"os"
)

type Config struct {
	DatabaseURI string
	Port        string
}

func LoadConfig() (*Config, error) {
	return &Config{
		DatabaseURI: getEnv("DATABASE_URI", "mongodb://localhost:27017/investment_management"),
		Port:        getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}