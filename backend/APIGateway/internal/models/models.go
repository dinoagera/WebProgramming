package models

type Good struct {
	ProductID string  `json:"product_id"`
	Category  string  `json:"category"`
	Sex       string  `json:"sex"`
	Sizes     []int   `json:"sizes"`
	Price     float64 `json:"price"`
	Color     string  `json:"color"`
	Tag       string  `json:"tag"`
	ImageURL  string  `json:"image_url"`
}
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
