package config

import (
	"os"
)

type Config struct {
	MongoURI        string
	MongoDB         string
	MongoCollection string
	Port            string
}

func LoadConfig() (*Config, error) {
	return &Config{
		MongoURI:        getEnv("MONGO_URI", "mongodb://localhost:27017/investment_management"),
		MongoDB:         getEnv("MONGO_DB", "invest-tracker"),
		MongoCollection: getEnv("MONGO_COLLECTION", "investments"),
		Port:            getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
