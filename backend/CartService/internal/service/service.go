package service

import (
	"CartService/internal/models"
	storage "CartService/internal/storage/interfaces"
	"log/slog"
	"strconv"
)

type Service struct {
	log     *slog.Logger
	getCart storage.GetCart
}

func New(log *slog.Logger, getCart storage.GetCart) *Service {
	return &Service{
		log:     log,
		getCart: getCart,
	}
}
func (s *Service) GetCart(userID int) (models.Cart, error) {
	key := strconv.Itoa(userID)
	cart, err := s.getCart.GetCart(key)
	if err != nil {
		s.log.Info("faield to get cart", "err", err)
		return cart, err
	}
	return cart, nil
}
