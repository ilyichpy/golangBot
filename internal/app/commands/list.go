package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Update) {
	outputMessage := "Here's all products:\n"
	productList := c.service.List()
	for _, product := range productList {
		outputMessage += product.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Message.Chat.ID, outputMessage)
	c.bot.Send(msg)
}
