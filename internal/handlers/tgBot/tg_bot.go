package tgBot

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotRouter struct {
	services handlers.Service
}

func NewBotRouter(services handlers.Service) *BotRouter {
	return &BotRouter{services: services}
}

func (r *BotRouter) HandleMessage(msg *tgbotapi.Message) error {
	fmt.Println(msg.Text)
	return nil
}
