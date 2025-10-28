package client

type PaymentService interface {
	Purchase(userID string) (float64, error)
}
