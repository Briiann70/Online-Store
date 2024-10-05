package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test-online-store/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:Aska2908_03@tcp(127.0.0.1:3306)/teststore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	db.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Cart{}, &models.CartItem{}, &models.Transaction{})
}
