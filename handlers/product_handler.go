package handlers

import (
	"go-product-inventory-api/models"
	"go-product-inventory-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, models.Products)
}
func GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
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
}
func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	newProduct.ID = len(models.Products) + 1
	models.Products = append(models.Products, newProduct)
	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}
func UpdateProduct(c *gin.Context) {
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
		if product.ID == id {
			updateProduct.ID = product.ID
			models.Products[index] = updateProduct
			c.JSON(http.StatusOK, updateProduct)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
}
func DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	for index, product := range models.Products {
		if product.ID == id {
			models.Products = append(models.Products[:index], models.Products[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
}
