package service

import (
	"log/slog"
	client "paymentservice/internal/client/interfaces"
)

type Service struct {
	log      *slog.Logger
	purchase client.PaymentService
}

func New(log *slog.Logger, purchase client.PaymentService) *Service {
	return &Service{
		log:      log,
		purchase: purchase,
	}
}
func (s *Service) Purchase(userID string) (float64, error) {
	price, err := s.purchase.Purchase(userID)
	if err != nil {
		s.log.Info("failed to purchase", "err", err)
		return 0.0, err
	}
	return price, nil
}
