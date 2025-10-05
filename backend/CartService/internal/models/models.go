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
	Name      string  `json:"name"`
}

type AddItemRequest struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Name      string  `json:"name"`
}
type UpdateItemRequest struct {
	ProductID     string `json:"product_id"`
	TypeOperation int    `json:"type_operation" `
}
