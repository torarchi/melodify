package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ReplicateAPIToken string
	ServerPort        string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{
		ReplicateAPIToken: os.Getenv("REPLICATE_API_TOKEN"),
		ServerPort:        os.Getenv("SERVER_PORT"),
	}

	if config.ReplicateAPIToken == "" {
		return nil, fmt.Errorf("REPLICATE_API_TOKEN is required")
	}

	if config.ServerPort == "" {
		config.ServerPort = "8080"
	}

	return config, nil
}
