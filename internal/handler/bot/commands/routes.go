package commands

import (
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"

	"github.com/sula7/quote-bot/internal/handler/bot/commands/start"
)

func RegisterRoutes(logger *zap.Logger) telebot.HandlerFunc {
	return start.RegisterRoutes(logger)
}
