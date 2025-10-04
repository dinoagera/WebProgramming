package storage

import "CartService/internal/models"

type GetCart interface {
	GetCart(key string) (models.Cart, error)
}
type AddItem interface {
	AddItem(key string, addItem models.AddItemRequest) error
}
