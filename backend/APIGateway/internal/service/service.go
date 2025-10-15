package service

import (
	"apigateway/internal/client"
	"apigateway/internal/models"
	"log/slog"
)

type Service struct {
	log            *slog.Logger
	catalogService client.CatalogService
}

func New(log *slog.Logger, catalogService client.CatalogService) *Service {
	return &Service{
		log:            log,
		catalogService: catalogService,
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
