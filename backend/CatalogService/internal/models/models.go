package models

// type Goods struct {
// 	Goods []Good `json:"goods"`
// }
type Good struct {
	ProductID string  `json:"product_id"`
	Category  string  `json:"category"`
	Sex       string  `json:"sex"`
	Size      int     `json:"size"`
	Price     float64 `json:"price"`
	Color     string  `json:"color"`
	Tag       string  `json:"tag"`
	ImageURL  string  `json:"image_url"`
}
