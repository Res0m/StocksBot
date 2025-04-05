package main

import (
	"StocksBot/main/cmd/api"
	"StocksBot/main/cmd/bot"
	"StocksBot/main/configs"
	"StocksBot/main/internal/services"
	"log"
	"os"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := services.InitDatabase(config.DatabaseURL); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("Application started successfully!")
	mode := os.Getenv("APP_MODE")
	if mode == "bot" {
		log.Println("Starting Telegram bot...")
		bot.StartBot()
	} else if mode == "api" {
		log.Println("Starting REST API server...")
		api.StartServer()
	} else {
		log.Fatal("Invalid APP_MODE. Use 'bot' or 'api'.")
	}
}
