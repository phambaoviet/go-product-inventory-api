package main

import (
	"go-product-inventory-api/models"
	"go-product-inventory-api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Products)
	})
	r.POST("/api/products", func(c *gin.Context) {
		var newProduct models.Product

		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
			return
		}
		newProduct.ID = len(models.Products) + 1
		// Add new product to products new
		models.Products = append(models.Products, newProduct)
		c.JSON(http.StatusCreated, newProduct)
	})
	r.GET("/api/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		// convert ID from string to int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
			return
		}
		for _, product := range models.Products {
			if product.ID == id {
				c.JSON(http.StatusOK, product)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
	})
	r.DELETE("/api/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
			return
		}
		for index, product := range models.Products {
			if product.ID == id {
				// Remove the product at the found index from the slice
				models.Products = append(models.Products[:index], models.Products[index+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "delete success"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
	})
	r.PUT("/api/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
			return
		}
		var updateProduct models.Product
		if err := c.ShouldBindJSON(&updateProduct); err != nil {
			c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
			return
		}
		for index, product := range models.Products {
			// Find and overwrite product in slice
			if product.ID == id {
				updateProduct.ID = product.ID
				models.Products[index] = updateProduct
				c.JSON(http.StatusOK, updateProduct)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server cannot start...: ", err)
	}
}
