package handlers

import (
	"go-product-inventory-api/models"
	"go-product-inventory-api/services"
	"go-product-inventory-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	slug := c.Query("slug")
	name := c.Query("name")
	if !utils.IsValidSlug(slug) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid slug"})
		return
	}
	if !utils.IsValidSearchName(name) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid search name"})
		return
	}
	product := services.GetProducts(services.ProductFilter{
		Slug: slug,
		Name: name,
	})
	c.JSON(http.StatusOK, product)
}

func GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	product, found := services.GetProductByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)

}
func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	createdProduct := services.CreateProduct(newProduct)

	c.JSON(http.StatusCreated, createdProduct)
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
	updateProduct, found := services.UpdateProduct(id, updateProduct)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, updateProduct)
}
func DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	deleted := services.DeleteProduct(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "delete success"})
}
func GetProductBySlug(c *gin.Context) {
	slug := c.Param("slug")
	c.JSON(http.StatusOK, gin.H{"slug": slug})
}
