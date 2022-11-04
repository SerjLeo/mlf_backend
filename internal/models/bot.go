package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler interface {
	HandleMessage(msg *tgbotapi.Message) error
}

type TgBot struct {
	botApi  *tgbotapi.BotAPI
	handler BotHandler
}

func NewTgBot(botApi *tgbotapi.BotAPI, handlers BotHandler) *TgBot {
	return &TgBot{
		botApi:  botApi,
		handler: handlers,
	}
}

func (b *TgBot) Run() error {
	return nil
}
