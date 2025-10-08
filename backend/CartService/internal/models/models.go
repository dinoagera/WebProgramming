package models

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
