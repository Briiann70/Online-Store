package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-online-store/database"
	"test-online-store/models"
	"test-online-store/services"
)

func AddToCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.AddItemToCart(cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Added to cart"})
}

func GetCartProducts(c *gin.Context) {
	customerID := c.Param("customer_id")

	// Ambil semua item di keranjang untuk customer
	var cartItems []models.CartItem
	if err := database.DB.
		Joins("JOIN carts ON carts.id = cart_items.cart_id").
		Where("carts.customer_id = ?", customerID).
		Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items", "details": err.Error()})
		return
	}

	if len(cartItems) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No items found in cart"})
		return
	}

	// Ambil detail produk berdasarkan product_id di cartItems
	var products []models.Product
	for _, item := range cartItems {
		var product models.Product
		if err := database.DB.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product", "product_id": item.ProductID, "details": err.Error()})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{"products": products, "cartItems": cartItems})
}

func RemoveFromCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.RemoveItemFromCart(cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
