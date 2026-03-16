package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{},
			fmt.Errorf("error loading .env file: %v", err)
	}

	mongoURI, err := extractEnv("MONGO_DB_URI", "")
	if err != nil {
		return Config{},
			fmt.Errorf("error loading .env file: %v", err)
	}

	mongoDB, err := extractEnv("MONGO_DB_NAME", "")
	if err != nil {
		return Config{},
			fmt.Errorf("error loading .env file: %v", err)
	}

	port, err := extractEnv("PORT", "")
	if err != nil {
		return Config{},
			fmt.Errorf("error loading .env file: %v", err)
	}

	return Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: port,
	}, nil
}

func extractEnv(key string, defaultValue string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("missing required env var: %s", key)
	}
	return val, nil
}
