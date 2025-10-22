package client

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
	GetCart(userID string) (models.Cart, error)
}
