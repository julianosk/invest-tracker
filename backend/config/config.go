package config

import (
	"os"
)

type Config struct {
	PostgresDSN string
	Port        string
}

func LoadConfig() (*Config, error) {
	return &Config{
		PostgresDSN: getEnv("POSTGRES_DSN", "host=postgres user=postgres password=postgres dbname=investments port=5432 sslmode=disable"),
		Port:        getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
