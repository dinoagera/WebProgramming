package service

import "CartService/internal/models"

type GetCart interface {
	GetCart(userID int) (models.Cart, error)
}
