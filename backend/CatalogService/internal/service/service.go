package service

import (
	"catalogservice/internal/models"
	storage "catalogservice/internal/storage/interfaces"
	"catalogservice/lib"
	"log/slog"
	"strconv"
)

type Service struct {
	log           *slog.Logger
	getCatalog    storage.GetCatalog
	getImage      storage.GetImage
	getFavourites storage.GetFavourites
}

func New(log *slog.Logger, getCatalog storage.GetCatalog, getImage storage.GetImage, getFavourites storage.GetFavourites) *Service {
	return &Service{
		log:           log,
		getCatalog:    getCatalog,
		getImage:      getImage,
		getFavourites: getFavourites,
	}
}
func (s *Service) GetCatalog() ([]models.Good, error) {
	goods, err := s.getCatalog.GetCatalog()
	if err != nil {
		if err == lib.ErrCatalogIsEmpty {
			s.log.Info("catalog is empty")
			return nil, err
		}
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
func (s *Service) GetFavourites(userID string) ([]models.Favourites, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		s.log.Info("failed to convert string to int", "err", err)
		return []models.Favourites{}, err
	}
	favourites, err := s.getFavourites.GetFavourites(userIDInt)
	if err != nil {
		s.log.Info("failed get favourites", "err", err)
		return []models.Favourites{}, err
	}
	return favourites, nil
}
