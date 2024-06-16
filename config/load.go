package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfigFromEnv() (*Config, error) {
	godotenv.Load()

	cfg := &Config{}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("missing environment variable BOT_TOKEN")
	}
	cfg.Token = token

	return cfg, nil
}
