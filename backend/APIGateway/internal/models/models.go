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
type Cart struct {
	UserId string     `json:"user_id"`
	Items  []CartItem `json:"items"`
	Total  float64    `json:"total"`
}
type CartItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Category  string  `json:"category"`
}

type AddItemRequest struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Category  string  `json:"category"`
}
type UpdateItemRequest struct {
	ProductID     string `json:"product_id"`
	TypeOperation int    `json:"type_operation" `
}
type RemoveItemRequest struct {
	ProductID string `json:"product_id"`
}
type Favourites struct {
	ProductID string  `json:"product_id"`
	Category  string  `json:"category"`
	Sex       string  `json:"sex"`
	Sizes     []int   `json:"sizes"`
	Price     float64 `json:"price"`
	Color     string  `json:"color"`
	Tag       string  `json:"tag"`
	ImageURL  string  `json:"image_url"`
}
type AddFavouriteRequest struct {
	ProductID string `json:"product_id"`
}
type RemoveFavouriteRequest struct {
	ProductID string `json:"product_id"`
}
