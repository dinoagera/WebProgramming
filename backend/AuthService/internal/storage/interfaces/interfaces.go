package storage

import "authservice/internal/domain/models"

type CreateUser interface {
	CreateUser(email, passHash string) error
}
type LoginUser interface {
	LoginUser(email string) (models.User, error)
}
