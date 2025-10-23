package service

import "apigateway/internal/models"

type CatalogService interface {
	GetCatalog() ([]models.Good, error)
	GetImage(productID string) ([]byte, error)
}
type AuthService interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
}
type CartService interface {
	GetCart(userID int64) (models.Cart, error)
	AddItem(userID int64, productID string, quantity int, price float64, category string) error
}
