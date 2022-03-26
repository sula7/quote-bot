package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New returns new zap logger.
func New(level string) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true

	if err := config.Level.UnmarshalText([]byte(level)); err != nil && level != "" {
		log.Printf("failed to set log level %q: %v", level, err)
	}

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	return logger, nil
}
