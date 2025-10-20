package service

import (
	"apigateway/internal/client"
	"apigateway/internal/models"
	"log/slog"
)

type Service struct {
	log            *slog.Logger
	catalogService client.CatalogService
	authServce     client.AuthService
}

func New(log *slog.Logger, catalogService client.CatalogService, authService client.AuthService) *Service {
	return &Service{
		log:            log,
		catalogService: catalogService,
		authServce:     authService,
	}
}
func (s *Service) GetCatalog() ([]models.Good, error) {
	goods, err := s.catalogService.GetCatalog()
	if err != nil {
		s.log.Info("failed to get catalog", "err", err)
		return nil, err
	}
	return goods, nil
}

func (s *Service) GetImage(productID string) ([]byte, error) {
	image, err := s.catalogService.GetImage(productID)
	if err != nil {
		s.log.Info("failed to get image", "err", err)
		return nil, err
	}
	return image, nil
}
func (s *Service) Register(email, password string) error {
	if err := s.authServce.Register(email, password); err != nil {
		s.log.Info("failed to register", "err", err)
		return err
	}
	return nil
}
func (s *Service) Login(email, password string) (string, error) {
	token, err := s.authServce.Login(email, password)
	if err != nil {
		s.log.Info("failed to register", "err", err)
		return "", err
	}
	return token, nil
}
