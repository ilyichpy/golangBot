package main

import (
	"github.com/joho/godotenv"
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

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, &update)
			default:
				defaultTextInput(bot, &update)
			}

		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Update) {
	log.Printf("[%s] %s", inputMessage.Message.From.UserName, inputMessage.Message.Text)
	msg := tgbotapi.NewMessage(inputMessage.Message.Chat.ID, "/help command return all commands")
	bot.Send(msg)
}

func defaultTextInput(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Update) {
	log.Printf("[%s] %s", inputMessage.Message.From.UserName, inputMessage.Message.Text)
	msg := tgbotapi.NewMessage(inputMessage.Message.Chat.ID, inputMessage.Message.Text)
	bot.Send(msg)
}
