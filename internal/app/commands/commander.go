package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golangBot/internal/services/product"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	service *product.Service,
) *Commander {
	return &Commander{
		bot:     bot,
		service: service,
	}
}
