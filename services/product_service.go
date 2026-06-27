package services

import (
	"go-product-inventory-api/models"
	"go-product-inventory-api/repositories"
	"strings"
)

type ProductFilter struct {
	Slug string
	Name string
}

func GetProducts(filter ProductFilter) []models.Product {
	products := repositories.GetAllProducts()
	filteredProduct := []models.Product{}
	for _, product := range products {
		if filter.Slug != "" && product.Slug != filter.Slug {
			continue
		}
		if filter.Name != "" && !strings.Contains(strings.ToLower(product.Name), strings.ToLower(filter.Name)) {
			continue
		}
		filteredProduct = append(filteredProduct, product)
	}
	return filteredProduct
}
func GetProductByID(id int) (models.Product, bool) {
	return repositories.GetProductByID(id)
}
func CreateProduct(product models.Product) models.Product {
	return repositories.CreateProduct(product)
}
func UpdateProduct(id int, product models.Product) (models.Product, bool) {
	return repositories.UpdateProduct(id, product)
}
func DeleteProduct(id int) bool {
	return repositories.DeleteProduct(id)
}
