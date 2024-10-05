package services

import (
	"test-online-store/database"
	"test-online-store/models"
)

func CreateTransaction(transaction models.Transaction) error {
	// Simpan transaksi ke database
	result := database.DB.Create(&transaction)
	return result.Error
}
