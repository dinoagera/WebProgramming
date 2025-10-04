package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerAddr    string        `env:"SERVER_ADDRESS" env-default:":8081"`
	RedisAddr     string        `env:"REDIS_ADDRESS" env-default:"localhost:6379"`
	RedisPassword string        `env:"REDIS_PASSWORD"`
	DB            int           `env:"REDIS_DB"`
	CartTTL       time.Duration `env:"CART_TTL"`
	MaxRetries    int           `env:"MAX_RETRIES"`
	DialTimeout   time.Duration `env:"DialTimeout"`
	Timeout       time.Duration `env:"Timeout"`
	ReadTimeout   time.Duration `env:"HTTPReadTimeout"`
	IdleTimeout   time.Duration `env:"HTTPidleTimeout"`
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
