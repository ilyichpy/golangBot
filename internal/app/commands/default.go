package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Default(inputMessage *tgbotapi.Update) {
	log.Printf("[%s] %s", inputMessage.Message.From.UserName, inputMessage.Message.Text)
	msg := tgbotapi.NewMessage(inputMessage.Message.Chat.ID, inputMessage.Message.Text)
	c.bot.Send(msg)
}
