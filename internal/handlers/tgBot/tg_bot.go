package tgBot

import (
	"github.com/SerjLeo/mlf_backend/internal/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotRouter struct {
	services handlers.Service
}

func NewBotRouter(services handlers.Service) *BotRouter {
	return &BotRouter{services: services}
}

func (r *BotRouter) HandleMessage(msg *tgbotapi.Message, api *tgbotapi.BotAPI) error {

	newMsg := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	newMsg.ReplyToMessageID = msg.MessageID

	api.Send(newMsg)
	return nil
}
