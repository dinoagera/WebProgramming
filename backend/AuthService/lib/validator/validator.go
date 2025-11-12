package validator

import (
	liberror "authservice/lib/errors"
	"errors"
	"strings"
)

func ValidateEmail(email string) error {
	if email == "" {
		return liberror.ErrEmailEmpty
	}
	if !strings.Contains(email, "@") {
		return liberror.ErrEmailNotAllowed
	}
	return nil
}
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password less 6 symbol")
	}
	return nil
}
