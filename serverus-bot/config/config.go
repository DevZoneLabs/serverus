package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DiscordToken string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error while loading .env file")
		return nil, err
	}

	cfg := &Config{
		DiscordToken: getEnv("SERVERUS_TOKEN", ""),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}


