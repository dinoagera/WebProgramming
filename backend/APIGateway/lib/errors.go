package lib

import "errors"

var (
	ErrImageNotFound  = errors.New("image not found")
	ErrCatalogIsEmpty = errors.New("catalog is empty")
)
