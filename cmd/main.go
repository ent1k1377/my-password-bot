package main

import (
	"github.com/ent1k1377/my-password-bot/internal/bot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	loadEnvironment()
}

func main() {
	// initializing bot
	token := os.Getenv("BOT_API_TOKEN")
	bot, err := telegrambot.New(token)
	if err != nil {
		log.Panic(err)
	}

	if err := bot.SetConfig(); err != nil {
		log.Fatal(err)
	}

	bot.Start()
}

func loadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
