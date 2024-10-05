package services

import (
	"test-online-store/database"
	"test-online-store/models"
)

// Implement AddItemToCart
func AddItemToCart(cartItem models.CartItem) error {
	result := database.DB.Create(&cartItem)
	return result.Error
}

// Existing GetCartItems function
func GetCartItems(customerID string) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	result := database.DB.Where("customer_id = ?", customerID).Find(&cartItems)
	return cartItems, result.Error
}

// Existing RemoveItemFromCart function
func RemoveItemFromCart(cartItem models.CartItem) error {
	result := database.DB.Where("id = ?", cartItem.ID).Delete(&cartItem)
	return result.Error
}
