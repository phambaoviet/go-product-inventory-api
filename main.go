package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Apple", Price: 100},
	{ID: 2, Name: "Banana", Price: 200},
	{ID: 3, Name: "Orange", Price: 300},
}

func main() {
	r := gin.Default()

	r.GET("/api/products", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": products})
	})
	r.POST("/api/products", func(c *gin.Context) {
		var newProduct Product

		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(404, gin.H{"error": "invalid json"})
			return
		}
		newProduct.ID = len(products) + 1
		// Add new product to products new
		products = append(products, newProduct)
		c.JSON(201, newProduct)
	})
	r.GET("/api/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		// convert ID from string to int
		id, _ := strconv.Atoi(idStr)
		for _, product := range products {
			if product.ID == id {
				c.JSON(200, gin.H{"message": product})
				return
			}
		}
		c.JSON(404, gin.H{"message": "Product not found"})
	})
	r.DELETE("/api/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		for index, product := range products {
			if product.ID == id {
				// Remove the product at the found index from the slice
				products = append(products[:index], products[index+1:]...)
				c.JSON(200, gin.H{"message": product})
				return
			}
		}
		c.JSON(404, gin.H{"message": "Product not found"})
	})
	r.PUT("/api/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		var updateProduct Product
		for index, product := range products {
			// Find and overwrite product in slice
			if product.ID == id {
				updateProduct = product
				products[index] = updateProduct
				c.JSON(200, gin.H{"message": product})
				return
			}
		}
		c.JSON(404, gin.H{"message": "Product not found"})
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server cannot start...: ", err)
	}
}
