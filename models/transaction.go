package models

type Transaction struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	CustomerID  uint    `json:"customer_id"`
	TotalAmount float64 `json:"total_amount"`
}
