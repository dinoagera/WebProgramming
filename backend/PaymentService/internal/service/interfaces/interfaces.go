package service

type Paymentservice interface {
	Purchase(userID string) (float64, error)
}
