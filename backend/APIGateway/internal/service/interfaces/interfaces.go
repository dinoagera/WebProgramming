package service

import (
	"apigateway/internal/models"
)

type CatalogService interface {
	GetCatalog() ([]models.Good, error)
	GetImage(productID string) ([]byte, error)
	GetFavourites(userID int64) ([]models.Favourites, error)
	AddFavourite(userID int64, productID string) error
	RemoveFavourite(userID int64, productID string) error
	GetMale() ([]models.Good, error)
	GetFemale() ([]models.Good, error)
}
type AuthService interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
}
type CartService interface {
	GetCart(userID int64) (models.Cart, error)
	AddItem(userID int64, productID string, quantity int, price float64, category string) error
	RemoveItem(userID int64, productID string) error
	UpdateItem(userID int64, productID string, typeOperation int) error
	ClearCart(userID int64) error
}
