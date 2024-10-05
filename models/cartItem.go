package models

type CartItem struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
