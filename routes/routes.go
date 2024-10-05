package routes

import (
	"github.com/gin-gonic/gin"
	"test-online-store/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Auth routes
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

	// Other routes (products, cart, etc.)
	r.POST("/products/search", controllers.GetProductsByCategory)
	r.POST("/products", controllers.CreateProduct)
	r.POST("/cart", controllers.AddToCart)
	r.GET("/cart/products/:customer_id", controllers.GetCartProducts)
	r.DELETE("/cart", controllers.RemoveFromCart)
	r.POST("/checkout", controllers.Checkout)

	return r
}
