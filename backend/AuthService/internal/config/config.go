package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	StoragePath string        `env:"DB_URL" env-required:"true"`
	SecretKey   string        `env:"JWT_SECRET" env-required:"true"`
	TokenTTL    time.Duration `env:"TOKEN_TTL" env-default:"30m"`
	Address     string        `env:"SERVER_ADDRESS" env-required:"true"`
	ReadTimeout time.Duration `env:"HTTPReadTimeout" env-default:"5s"`
	IdleTimeout time.Duration `env:"HTTPidleTimeout" env-default:"60s"`
}

var cfg *Config

func InitConfig(log *slog.Logger) *Config {
	cfgPath := ".env"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Info("config path is empty")
		os.Exit(1)
	}
	var cfgLocal Config
	if err := cleanenv.ReadConfig(cfgPath, &cfgLocal); err != nil {
		log.Info("failed to read config", "error:", err)
		os.Exit(1)
	}
	cfg = &cfgLocal
	return cfg
}
func GetConfig() *Config {
	if cfg == nil {
		panic("cfg is not init")
	}
	return cfg
}
