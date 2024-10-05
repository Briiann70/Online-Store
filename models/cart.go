package models

type Cart struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	CustomerID uint       `json:"customer_id"`
	Items      []CartItem `json:"items" gorm:"foreignKey:CartID"`
}
