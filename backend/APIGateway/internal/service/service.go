package service

import (
	"apigateway/internal/client"
	"apigateway/internal/models"
	"log/slog"
	"strconv"
)

type Service struct {
	log            *slog.Logger
	catalogService client.CatalogService
	authServce     client.AuthService
	cartService    client.CartService
}

func New(log *slog.Logger, catalogService client.CatalogService, authService client.AuthService, cartService client.CartService) *Service {
	return &Service{
		log:            log,
		catalogService: catalogService,
		authServce:     authService,
		cartService:    cartService,
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

func (s *Service) GetCart(userID int64) (models.Cart, error) {
	userIDStr := strconv.Itoa(int(userID))
	cart, err := s.cartService.GetCart(userIDStr)
	if err != nil {
		s.log.Info("failed to get cart", "err", err)
		return models.Cart{}, err
	}
	return cart, nil
}
func (s *Service) AddItem(userID int64, productID string, quantity int, price float64, category string) error {
	userIDStr := strconv.Itoa(int(userID))
	err := s.cartService.AddItem(userIDStr, productID, quantity, price, category)
	if err != nil {
		s.log.Info("failed to add item", "err", err)
		return err
	}
	return nil
}
func (s *Service) RemoveItem(userID int64, productID string) error {
	userIDStr := strconv.Itoa(int(userID))
	err := s.cartService.RemoveItem(userIDStr, productID)
	if err != nil {
		s.log.Info("failed to remove item", "err", err)
		return err
	}
	return nil
}
