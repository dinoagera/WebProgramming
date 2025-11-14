package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Address     string        `env:"SERVER_ADDRESS"`
	ReadTimeout time.Duration `env:"HTTPReadTimeout" env-default:"5s"`
	IdleTimeout time.Duration `env:"HTTPidleTimeout" env-default:"60s"`
}

func InitConfig(log *slog.Logger) *Config {
	cfgPath := ".env"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Info("config path is empty")
		os.Exit(1)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Info("failed to read config", "err", err)
		os.Exit(1)
	}
	return &cfg
}
