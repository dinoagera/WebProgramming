package service

import "CartService/internal/models"

type GetCart interface {
	GetCart(key string) (models.Cart, error)
}
type AddItem interface {
	AddItem(key string, addItem models.AddItemRequest) error
}
type RemoveItem interface {
	RemoveItem(key string, removeItem models.CartItem) error
}
type UpdateItem interface {
	UpdateItem(key string, updateItem models.UpdateItemRequest) error
}
type ClearCart interface {
	ClearCart(key string) error
}
