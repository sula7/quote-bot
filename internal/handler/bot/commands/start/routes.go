package start

import (
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

func RegisterRoutes(logger *zap.Logger) telebot.HandlerFunc {
	handler := NewBotHandler(logger)
	return handler.start
}
