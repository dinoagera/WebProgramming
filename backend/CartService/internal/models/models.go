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

// type AddItemRequest struct {
// 	ProductID string  `json:"product_id" binding:"required"`
// 	Quantity  int     `json:"quantity" binding:"min=1"`
// 	Price     float64 `json:"price" binding:"min=0"`
// 	Name      string  `json:"name" binding:"required"`
// }
