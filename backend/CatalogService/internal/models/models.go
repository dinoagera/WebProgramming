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
type Favourites struct {
	UserID    string  `json:"user_id"`
	ProductID string  `json:"product_id"`
	Category  string  `json:"category"`
	Sex       string  `json:"sex"`
	Sizes     []int   `json:"sizes"`
	Price     float64 `json:"price"`
	Color     string  `json:"color"`
	Tag       string  `json:"tag"`
	ImageURL  string  `json:"image_url"`
}
type AddRequest struct {
	ProductID string `json:"product_id"`
}
type RemoveRequest struct {
	ProductID string `json:"product_id"`
}
