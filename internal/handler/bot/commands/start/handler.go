package start

import (
	"fmt"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

type BotHandler struct {
	logger *zap.Logger
}

func (h *BotHandler) start(c telebot.Context) error {
	return c.Send(fmt.Sprint("Hello, ", c.Message().Sender.Username))
}

func newBotHandler(logger *zap.Logger) *BotHandler {
	return &BotHandler{logger: logger}
}
