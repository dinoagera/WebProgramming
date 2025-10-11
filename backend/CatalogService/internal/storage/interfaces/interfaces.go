package storage

import "catalogservice/internal/models"

type GetCatalog interface {
	GetCatalog() ([]models.Good, error)
}
type GetImage interface {
	GetImage(productID string) ([]byte, error)
}
