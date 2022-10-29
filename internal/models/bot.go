package models

import (
	"github.com/SerjLeo/mlf_backend/internal/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	botApi  *tgbotapi.BotAPI
	handler handlers.BotHandler
}

func NewTgBot(botApi *tgbotapi.BotAPI, handlers handlers.BotHandler) *TgBot {
	return &TgBot{
		botApi:  botApi,
		handler: handlers,
	}
}

func (b *TgBot) Run() error {
	return nil
}
