package handler

import (
	"github.com/sula7/quote-bot/internal/handler/bot/commands/start"
)

func (b *Bot) registerRoutes() {
	b.api.Handle("/start", start.RegisterRoutes(b.logger))
}
