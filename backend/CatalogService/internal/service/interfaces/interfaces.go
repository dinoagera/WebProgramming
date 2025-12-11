package service

import "catalogservice/internal/models"

type GetCatalog interface {
	GetCatalog() ([]models.Good, error)
}
type GetImage interface {
	GetImage(productID string) ([]byte, error)
}
type GetFavourites interface {
	GetFavourites(userID string) ([]models.Favourites, error)
}
type AddFavourite interface {
	AddFavourite(userID, productID string) error
}
type RemoveFavourite interface {
	RemoveFavourite(userID, productID string) error
}
type GetPol interface {
	GetMale() ([]models.Good, error)
	GetFemale() ([]models.Good, error)
}
