package main

import (
	"dripcord/config"
	"dripcord/handlers"
	"log"
)

func main() {
	cfg := config.Load()
	bot, err := handlers.NewBot(cfg)
	if err != nil {
		log.Fatal("Failed to start bot:", err)
	}
	bot.Run()
}
