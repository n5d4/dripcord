package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DiscordBotToken string
	GeminiToken     string
	DiscordChannel  string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		DiscordBotToken: os.Getenv("DISCORD_BOT_TOKEN"),
		GeminiToken:     os.Getenv("GEMINI_TOKEN"),
		DiscordChannel:  os.Getenv("DISCORD_CHANNEL"),
	}
}
