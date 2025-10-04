package service

import "CartService/internal/models"

type GetCart interface {
	GetCart(userID string) (models.Cart, error)
}
type AddItem interface {
	AddItem(userID string, addItem models.AddItemRequest) error
}
type RemoveItem interface {
	RemoveItem(userID string) error
}
