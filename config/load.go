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

func LoadStorageFromEnv() (*Storage, error) {
	godotenv.Load()

	strg := NewStorage()

	sa := os.Getenv("GOOGLE_SERVICE_ACCOUNT")
	if sa == "" {
		log.Fatalf("missing environment variable GOOGLE_SERVICE_ACCOUNT")
	}
	strg.ServiceAccount = sa

	bucket := os.Getenv("GOOGLE_STORAGE_BUCKET")
	if bucket == "" {
		log.Fatalf("missing environment variable GOOGLE_STORAGE_BUCKET")
	}
	strg.Bucket = bucket

	return strg, nil
}
