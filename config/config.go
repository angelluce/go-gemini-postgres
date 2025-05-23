package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL          string
	GeminiAPIKey         string
	Port                 string
	GoogleCloudTTSAPIKey string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DatabaseURL:          os.Getenv("DATABASE_URL"),
		GeminiAPIKey:         os.Getenv("GEMINI_API_KEY"),
		Port:                 os.Getenv("PORT"),
		GoogleCloudTTSAPIKey: os.Getenv("GOOGLE_CLOUD_TTS_API_KEY"),
	}
}
