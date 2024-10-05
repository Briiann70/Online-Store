package services

import (
	"test-online-store/database"
	"test-online-store/models"
)

func GetProductsByCategory(category string) ([]models.Product, error) {
	var products []models.Product

	if category != "" {
		// Jika ada kategori, filter berdasarkan kategori
		result := database.DB.Where("category = ?", category).Find(&products)
		return products, result.Error
	}

	// Jika tidak ada kategori, ambil semua produk
	result := database.DB.Find(&products)
	return products, result.Error
}

func CreateProduct(product models.Product) error {
	result := database.DB.Create(&product)
	return result.Error
}
