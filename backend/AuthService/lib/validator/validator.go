package validator

import (
	"errors"
	"strings"
)

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is empty")
	}
	if !strings.Contains(email, "@") {
		return errors.New("email is not allowed")
	}
	return nil
}
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password less 6 symbol")
	}
	return nil
}
