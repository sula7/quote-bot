package handler

import "github.com/sula7/quote-bot/internal/handler/bot/commands"

func (b *Bot) registerRoutes() {
	b.api.Handle("/start", commands.RegisterRoutes(b.logger))
}
