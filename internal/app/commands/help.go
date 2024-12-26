package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Help(inputMessage *tgbotapi.Update) {
	commandList := "/help - all comands \n/list - return list params"
	log.Printf("[%s] %s", inputMessage.Message.From.UserName, inputMessage.Message.Text)
	msg := tgbotapi.NewMessage(inputMessage.Message.Chat.ID, commandList)
	c.bot.Send(msg)
}
