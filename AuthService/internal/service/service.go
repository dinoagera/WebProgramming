package service

import (
	"authservice/internal/handler"
	storage "authservice/internal/storage/interfaces"
	lib "authservice/lib/jwt"
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	log        *slog.Logger
	createUser storage.CreateUser
	loginUser  storage.LoginUser
}

func New(log *slog.Logger, createUser storage.CreateUser, loginUser storage.LoginUser) *Service {
	return &Service{
		log:        log,
		createUser: createUser,
		loginUser:  loginUser,
	}
}
func (s *Service) Register(email, password string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Info("Failed to generate passhash", "err", err)
		return errors.New("failed to generate passhash")
	}
	err = s.createUser.CreateUser(email, string(passHash))
	if err != nil {
		if errors.Is(err, handler.ErrEmailBusy) {
			s.log.Info("email is busy", "email", email)
			return handler.ErrEmailBusy
		} else {
			return err
		}
	}
	return nil
}
func (s *Service) Login(email, password string) (string, error) {
	user, err := s.loginUser.LoginUser(email)
	if err != nil {
		s.log.Info("failed to login user", "err:", err)
		return "", errors.New("failed to login user")
	}
	token, err := lib.GenerateJWT(user)
	if err != nil {
		s.log.Info("failed to generate token", "err", err)
		return "", errors.New("failed to generate token")
	}
	return token, nil
}
