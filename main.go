package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/api/products", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Get all products"})
	})

	r.Run(":8080")
}
