package storage

import (
	"CartService/internal/config"
	"CartService/internal/models"
	"context"
	"encoding/json"
	"log/slog"

	"github.com/go-redis/redis/v8"
)

type Storage struct {
	log    *slog.Logger
	client *redis.Client
	cfg    *config.Config
}

func New(log *slog.Logger, client *redis.Client, cfg *config.Config) *Storage {
	return &Storage{
		log:    log,
		client: client,
		cfg:    cfg,
	}
}
func (s *Storage) saveCart(key string, cart models.Cart) error {
	data, err := json.Marshal(cart)
	if err != nil {
		s.log.Info("failed to marshall", "err", err)
		return err
	}
	err = s.client.Set(context.Background(), key, data, s.cfg.CartTTL).Err()
	if err != nil {
		s.log.Info("failed to set date in redis", "err", err)
		return err
	}
	s.log.Info("set date is successfully")
	return nil
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
func (s *Storage) AddItem(key string, addItem models.AddItemRequest) error {
	cart, err := s.GetCart(key)
	if err != nil {
		s.log.Info("failed to get cart", "err", err)
		return err
	}
	itemFound := false
	for i, item := range cart.Items {
		if item.ProductID == addItem.ProductID {
			cart.Items[i].Quantity += addItem.Quantity
			cart.Items[i].Price = addItem.Price
			cart.Items[i].Name = addItem.Name
			itemFound = true
			break
		}
	}
	if !itemFound {
		item := models.CartItem(addItem)
		cart.Items = append(cart.Items, item)
	}
	return s.saveCart(key, cart)
}
