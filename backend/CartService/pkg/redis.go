package pkg

import (
	"CartService/internal/config"
	"context"
	"log/slog"

	"github.com/go-redis/redis/v8"
)

func NewClient(log *slog.Logger, cfg *config.Config) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:         cfg.RedisAddr,
		Password:     cfg.RedisPassword,
		DB:           cfg.DB,
		MaxRetries:   cfg.MaxRetries,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	})
	if err := db.Ping(context.Background()).Err(); err != nil {
		log.Info("failed to connect to Redis Base.", "err:", err)
		return db, err
	}
	return db, nil
}
