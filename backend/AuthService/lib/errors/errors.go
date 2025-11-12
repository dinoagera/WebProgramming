package liberror

import "errors"

var (
	ErrEmailEmpty      = errors.New("email is empty")
	ErrEmailNotAllowed = errors.New("email is not allowed")
	ErrEmailBusy       = errors.New("email is busy")
)
