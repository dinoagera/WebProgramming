package service

import (
	"catalogservice/internal/models"
	storage "catalogservice/internal/storage/interfaces"
	"log/slog"
)

type Service struct {
	log        *slog.Logger
	getCatalog storage.GetCatalog
}

func New(log *slog.Logger) *Service {
	return &Service{log: log}
}
func (s *Service) GetCatalog() (models.Goods, error) {
	goods, err := s.getCatalog.GetCatalog()
	if err != nil {
		s.log.Info("failed to get catalog", "err", err)
		return goods, err
	}
	return goods, nil
}
