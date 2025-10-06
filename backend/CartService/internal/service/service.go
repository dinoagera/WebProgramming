package service

import (
	"CartService/internal/models"
	storage "CartService/internal/storage/interfaces"
	"log/slog"
)

type Service struct {
	log        *slog.Logger
	getCart    storage.GetCart
	addItem    storage.AddItem
	removeItem storage.RemoveItem
	updateItem storage.UpdateItem
	clearCart  storage.ClearCart
}

func New(log *slog.Logger, getCart storage.GetCart, addItem storage.AddItem, removeItem storage.RemoveItem, updateItem storage.UpdateItem, clearCart storage.ClearCart) *Service {
	return &Service{
		log:        log,
		getCart:    getCart,
		addItem:    addItem,
		updateItem: updateItem,
		removeItem: removeItem,
		clearCart:  clearCart,
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
	err := s.addItem.AddItem(userID, addItem)
	if err != nil {
		s.log.Info("failed to add item", "err", err)
		return err
	}
	return nil
}

func (s *Service) RemoveItem(userID string, removeItem models.CartItem) error {
	err := s.removeItem.RemoveItem(userID, removeItem)
	if err != nil {
		s.log.Info("failed to remove item", "err", err)
		return err
	}
	return nil
}
func (s *Service) UpdateItem(userID string, updateItem models.UpdateItemRequest) error {
	err := s.updateItem.UpdateItem(userID, updateItem)
	if err != nil {
		s.log.Info("failed to update item", "err", err)
		return err
	}
	return nil
}
func (s *Service) ClearCart(userID string) error {
	err := s.clearCart.ClearCart(userID)
	if err != nil {
		s.log.Info("failed to clear cart", "err", err)
		return err
	}
	return nil
}
