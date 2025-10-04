package lib

import "errors"

var (
	ErrCartNotFound  = errors.New("cart not found")
	ErrUserIDIsEmpty = errors.New("user id is empty")
)
