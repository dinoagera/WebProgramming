package storage

import (
	"CartService/internal/models"
	"context"
	"encoding/json"
	"log/slog"

	"github.com/go-redis/redis/v8"
)

type Storage struct {
	log    *slog.Logger
	client *redis.Client
}

func New(log *slog.Logger, client *redis.Client) *Storage {
	return &Storage{
		log:    log,
		client: client,
	}
}
func (s *Storage) GetCart(key string) (models.Cart, error) {
	data, err := s.client.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			s.log.Info("cart is not found")
			return models.Cart{
				UserId: key,
				Items:  []models.CartItem{},
				Total:  0,
			}, nil
		}
		s.log.Info("failed to get cart", "err", err)
		return models.Cart{}, err
	}
	var cart models.Cart
	if err = json.Unmarshal([]byte(data), &cart); err != nil {
		s.log.Info("failed to unmartshall data", "err", err)
		return models.Cart{}, err
	}
	return cart, nil
}
