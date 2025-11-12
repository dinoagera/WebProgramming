package lib

import "errors"

var (
	ErrImageNotFound  = errors.New("image not found")
	ErrCatalogIsEmpty = errors.New("catalog is empty")
	// ErrEmailEmpty      = errors.New("email is empty")
	// ErrEmailNotAllowed = errors.New("email is not allowed")
	// ErrEmailBusy       = errors.New("email is busy")
)
