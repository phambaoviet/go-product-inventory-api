package routes

import (
	"go-product-inventory-api/handlers"
	"go-product-inventory-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	productRoutes := r.Group("/api/products")
	productRoutes.Use(middlewares.ApikeyMiddleware())
	productRoutes.GET("/", handlers.GetProducts)
	productRoutes.POST("/", handlers.CreateProduct)

	productRoutes.GET("/slug/:slug", handlers.GetProductBySlug)
	productRoutes.GET("/:id", handlers.GetProductByID)
	productRoutes.PUT("/:id", handlers.UpdateProduct)
	productRoutes.DELETE("/:id", handlers.DeleteProduct)
}
