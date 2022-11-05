package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler interface {
	HandleMessage(msg *tgbotapi.Message, api *tgbotapi.BotAPI) error
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
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.botApi.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			err := b.handler.HandleMessage(update.Message, b.botApi)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
