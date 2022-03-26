package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Telegram *Telegram
		Database *Database
		Logger   *Logger
	}

	Telegram struct {
		Token  string
		Poller *Poller
	}

	Poller struct {
		Timeout time.Duration
	}

	Logger struct {
		Level string
	}

	Database struct {
		Name string
	}
)

func (c *Database) GetDBPath(appPath string) string {
	return filepath.Join(appPath, "db", c.Name)
}

// NewConfig parses & returns app config.
func NewConfig(configPath string) (*AppConfig, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read in: %w", err)
	}

	config := AppConfig{}
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return &config, nil
}
