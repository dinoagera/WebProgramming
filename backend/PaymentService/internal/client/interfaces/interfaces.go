package client

type PaymentService interface {
	GetTotalPrice(userID string) (float64, error)
	ClearCart(userID string) error
}
