package services

import (
	"errors"
	"test-online-store/database"
	"test-online-store/models"
)

func RegisterCustomer(customer models.Customer) error {
	result := database.DB.Create(&customer)
	return result.Error
}

// Updated LoginCustomer function
func LoginCustomer(email string, password string) (models.Customer, error) {
	var customer models.Customer
	result := database.DB.Where("email = ? AND password = ?", email, password).First(&customer)
	if result.Error != nil {
		return customer, errors.New("invalid email or password")
	}
	return customer, nil
}
