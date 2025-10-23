package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Address         string        `env:"SERVER_ADDRESS" env-reqired:"true"`
	HTTPReadTimeout time.Duration `env:"HTTPReadTimeout" env-defautl:"5s"`
	HTTPidleTimeout time.Duration `env:"HTTPReadTimeout" env-defautl:"60s"`
	CatalogAddress  string        `env:"CatalogAddress"`
	AuthAddress     string        `env:"AuthAddress"`
	CartAddress     string        `env:"CartAddress"`
	JWTSecret       string        `env:"JWT_SECRET"`
}

func InitConfig(log *slog.Logger) *Config {
	cfgPath := ".env"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Info("config path is empty")
		os.Exit(1)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Info("failed to read config", "error:", err)
		os.Exit(1)
	}
	return &cfg

}
