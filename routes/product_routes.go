package routes

import (
	"go-product-inventory-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	productRoutes := r.Group("/api/products")
	productRoutes.GET("/", handlers.GetProducts)
	productRoutes.POST("/", handlers.CreateProduct)

	productRoutes.GET("/:id", handlers.GetProductByID)
	productRoutes.PUT("/:id", handlers.UpdateProduct)
	productRoutes.DELETE("/:id", handlers.DeleteProduct)
}
