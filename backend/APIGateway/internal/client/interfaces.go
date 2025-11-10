package client

import "apigateway/internal/models"

type CatalogService interface {
	GetCatalog() ([]models.Good, error)
	GetImage(productID string) ([]byte, error)
	GetFavourites(userID string) ([]models.Favourites, error)
	AddFavourite(userID, productID string) error
}
type AuthService interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
}
type CartService interface {
	GetCart(userID string) (models.Cart, error)
	AddItem(userID string, productID string, quantity int, price float64, category string) error
	RemoveItem(userID string, productID string) error
	UpdateItem(userID string, productID string, typeOperation int) error
	ClearCart(userID string) error
}
