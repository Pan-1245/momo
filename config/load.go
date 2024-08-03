package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfigFromEnv() (*Config, error) {
	godotenv.Load()

	cfg := NewConfig()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalf("missing environment variable BOT_TOKEN")
	}
	cfg.Token = fmt.Sprintf("Bot %s", token)

	return cfg, nil
}
