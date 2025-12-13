package service

import (
	"catalogservice/internal/models"
	storage "catalogservice/internal/storage/interfaces"
	"catalogservice/lib"
	"log/slog"
	"strconv"
)

type Service struct {
	log             *slog.Logger
	getCatalog      storage.GetCatalog
	getImage        storage.GetImage
	getFavourites   storage.GetFavourites
	addFavourite    storage.AddFavourite
	removeFavourite storage.RemoveFavourite
	getpol          storage.GetPol
}

func New(log *slog.Logger, getCatalog storage.GetCatalog, getImage storage.GetImage, getFavourites storage.GetFavourites, addFavourite storage.AddFavourite, removeFavourite storage.RemoveFavourite, getpol storage.GetPol) *Service {
	return &Service{
		log:             log,
		getCatalog:      getCatalog,
		getImage:        getImage,
		getFavourites:   getFavourites,
		addFavourite:    addFavourite,
		removeFavourite: removeFavourite,
		getpol:          getpol,
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
func (s *Service) GetProduct(id string) (models.Good, error) {
	good, err := s.getCatalog.GetProduct(id)
	if err != nil {
		s.log.Info("failed to get product", "err", err)
		return models.Good{}, err
	}
	return good, nil
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
		return nil, err
	}
	favourites, err := s.getFavourites.GetFavourites(userIDInt)
	if err != nil {
		s.log.Info("failed to get favourites", "err", err)
		return nil, err
	}
	return favourites, nil
}
func (s *Service) AddFavourite(userID, productID string) error {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		s.log.Info("failed to convert string to int", "err", err)
		return err
	}
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		s.log.Info("failed to convert string to int", "err", err)
		return err
	}
	err = s.addFavourite.AddFavourite(userIDInt, productIDInt)
	if err != nil {
		if err == lib.ErrAlreadyInFavourites {
			s.log.Info("failed to add, goods is already added", "err", err)
			return err
		}
		s.log.Info("failed to add favourite", "err", err)
		return err
	}
	s.log.Info("favourite added", "userID", userID, "productID", productID)
	return nil
}
func (s *Service) RemoveFavourite(userID, productID string) error {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		s.log.Info("failed to convert string to int", "err", err)
		return err
	}
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		s.log.Info("failed to convert string to int", "err", err)
		return err
	}
	err = s.removeFavourite.RemoveFavourite(userIDInt, productIDInt)
	if err != nil {
		if err == lib.ErrAlreadyDeleted {
			s.log.Info("failed to remove, product have been already deleted", "err", err)
			return err
		}
		s.log.Info("failed to remove favourite", "err", err)
		return err
	}
	return nil
}
func (s *Service) GetMale() ([]models.Good, error) {
	goods, err := s.getpol.GetMale()
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
func (s *Service) GetFemale() ([]models.Good, error) {
	goods, err := s.getpol.GetFemale()
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
