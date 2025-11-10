package lib

import "errors"

var (
	ErrImageNotFound       = errors.New("image not found")
	ErrCatalogIsEmpty      = errors.New("catalog is empty")
	ErrFavouritesIsEmpty   = errors.New("favourites is empty")
	ErrUserIDIsEmpty       = errors.New("user id is empty")
	ErrAlreadyInFavourites = errors.New("favourite good already added")
	ErrAlreadyDeleted      = errors.New("favourite good already deleted")
)
