package handlers

import (
	"go-product-inventory-api/models"
	"go-product-inventory-api/utils"
	"net/http"
	"strconv"
	"strings"

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
	filteredProduct := []models.Product{}
	for _, product := range models.Products {
		if slug != "" && product.Slug != slug {
			continue
		}
		if name != "" && !strings.Contains(strings.ToLower(product.Name), strings.ToLower(name)) {
			continue
		}
		filteredProduct = append(filteredProduct, product)
	}
	c.JSON(http.StatusOK, filteredProduct)
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
	newProduct.ID = models.NextProductID
	models.NextProductID++
	newProduct.Slug = utils.GenerateSlug(newProduct.Name)
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
			updateProduct.Slug = utils.GenerateSlug(updateProduct.Name)
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
func GetProductBySlug(c *gin.Context) {
	slug := c.Param("slug")
	c.JSON(http.StatusOK, gin.H{"slug": slug})
}
