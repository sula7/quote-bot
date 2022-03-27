package main

import (
	"log"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/sula7/quote-bot/internal/config"
	"github.com/sula7/quote-bot/internal/handler"
	ll "github.com/sula7/quote-bot/internal/logger"
	"github.com/sula7/quote-bot/internal/service/domain"
	"github.com/sula7/quote-bot/internal/storage/sqlite"
)

func main() {
	appPath := "./"
	configFilepath := filepath.Join(appPath, "configs", "config.yml")

	appConfig, err := config.NewConfig(configFilepath)
	if err != nil {
		log.Fatal("failed to get app config: ", err)
	}

	logger, err := ll.New(appConfig.Logger.Level)
	if err != nil {
		log.Fatal("failed to init logger: ", err)
	}

	db, err := sqlite.New(appConfig.Database.GetDBPath(appPath))
	if err != nil {
		logger.Fatal("failed to init database", zap.Error(err))
	}

	defer db.Close()

	if err := sqlite.Migrate(db, filepath.Join(appPath, "scripts", "migrations")); err != nil {
		log.Fatal("failed to exec migrations", zap.Error(err))
	}

	repository := sqlite.NewRepository()
	services := domain.NewService(repository)

	appBot := handler.NewBot(appConfig, services, logger)

	if err := appBot.Start(); err != nil {
		logger.Fatal("failed to init and prepare bot", zap.Error(err))
	}

	logger.Info("shutting down the logger")
	logger.Sync() // nolint:errcheck
}
