package storage

import "catalogservice/internal/models"

type GetCatalog interface {
	GetCatalog() (models.Goods, error)
}
