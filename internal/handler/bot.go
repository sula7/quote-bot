package handler

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"

	"github.com/sula7/quote-bot/internal/config"
	"github.com/sula7/quote-bot/internal/service"
)

type Bot struct {
	api     *telebot.Bot
	logger  *zap.Logger
	config  *config.AppConfig
	service *service.Service
}

func NewBot(config *config.AppConfig, service *service.Service, logger *zap.Logger) *Bot {
	return &Bot{
		logger:  logger,
		config:  config,
		service: service,
	}
}

func (b *Bot) Start() error {
	botConfig := telebot.Settings{
		Token:  b.config.Telegram.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	var err error
	if b.api, err = telebot.NewBot(botConfig); err != nil {
		return fmt.Errorf("failed to init bot: %w", err)
	}

	b.registerRoutes()

	defer b.api.Stop()

	b.logger.Info("bot is running")

	b.api.Start()

	return nil
}
