package service

import (
	"catalogservice/internal/models"
	storage "catalogservice/internal/storage/interfaces"
	"catalogservice/lib"
	"log/slog"
)

type Service struct {
	log        *slog.Logger
	getCatalog storage.GetCatalog
	getImage   storage.GetImage
}

func New(log *slog.Logger, getCatalog storage.GetCatalog, getImage storage.GetImage) *Service {
	return &Service{
		log:        log,
		getCatalog: getCatalog,
		getImage:   getImage,
	}
}
func (s *Service) GetCatalog() ([]models.Good, error) {
	goods, err := s.getCatalog.GetCatalog()
	if err != nil {
		s.log.Info("failed to get catalog", "err", err)
		return nil, err
	}
	return goods, nil
}
func (s *Service) GetImage(productID string) ([]byte, error) {
	imageData, err := s.getImage.GetImage(productID)
	if err != nil {
		if err == lib.ErrImageNotFound {
			s.log.Info("failed to get image", "err", err)
			return nil, err
		}
		s.log.Info("failed to get image", "err", err)
		return nil, err
	}
	return imageData, nil
}
