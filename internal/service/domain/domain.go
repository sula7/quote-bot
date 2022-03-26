package domain

import (
	"github.com/sula7/quote-bot/internal/service"
	"github.com/sula7/quote-bot/internal/storage"
)

func NewService(repository *storage.Repository) *service.Service {
	return &service.Service{}
}
