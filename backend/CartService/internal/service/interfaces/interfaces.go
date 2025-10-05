package service

import (
	"CartService/internal/models"
)

type GetCart interface {
	GetCart(userID string) (models.Cart, error)
}
type AddItem interface {
	AddItem(userID string, addItem models.AddItemRequest) error
}
type RemoveItem interface {
	RemoveItem(userID string, cartItem models.CartItem) error
}
type UpdateItem interface {
	UpdateItem(userID string, updateItem models.UpdateItemRequest) error
}
type ClearCart interface {
	ClearCart(userID string) error
}
