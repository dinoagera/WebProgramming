package service

import (
	"CartService/internal/models"
	storage "CartService/internal/storage/interfaces"
	"log/slog"
)

type Service struct {
	log      *slog.Logger
	getCart  storage.GetCart
	saveCart storage.AddItem
}

func New(log *slog.Logger, getCart storage.GetCart, saveCart storage.AddItem) *Service {
	return &Service{
		log:      log,
		getCart:  getCart,
		saveCart: saveCart,
	}
}
func (s *Service) GetCart(userID string) (models.Cart, error) {
	cart, err := s.getCart.GetCart(userID)
	if err != nil {
		s.log.Info("faield to get cart", "err", err)
		return cart, err
	}
	return cart, nil
}

func (s *Service) AddItem(userID string, addItem models.AddItemRequest) error {
	err := s.saveCart.AddItem(userID, addItem)
	if err != nil {
		s.log.Info("failed to save cart", "err", err)
		return err
	}
	return nil
}
