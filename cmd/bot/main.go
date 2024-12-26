package main

import (
	"github.com/joho/godotenv"
	"golangBot/internal/app/commands"
	"golangBot/internal/services/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	service := product.NewService()
	commander := commands.NewCommander(bot, service)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "help":
				commander.Help(&update)
			case "list":
				commander.List(&update)
			default:
				commander.Default(&update)
			}
		}
	}
}
