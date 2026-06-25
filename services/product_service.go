package services

import (
	"go-product-inventory-api/models"
	"go-product-inventory-api/utils"
	"strings"
)

type ProductFilter struct {
	Slug string
	Name string
}

func GetProducts(filter ProductFilter) []models.Product {
	filteredProduct := []models.Product{}
	for _, product := range models.Products {
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
func GetAllProducts() []models.Product {
	return models.Products
}
func GetProductByID(id int) (models.Product, bool) {
	for _, product := range models.Products {
		if product.ID == id {
			return product, true
		}
	}
	return models.Product{}, false
}
func CreateProduct(product models.Product) models.Product {
	product.ID = models.NextProductID
	models.NextProductID++
	product.Slug = utils.GenerateSlug(product.Name)
	models.Products = append(models.Products, product)
	return product
}
func UpdateProduct(id int, product models.Product) (models.Product, bool) {
	for index, existingProduct := range models.Products {
		if existingProduct.ID == id {
			product.ID = existingProduct.ID
			product.Slug = utils.GenerateSlug(product.Name)

			models.Products[index] = product
			return product, true
		}
	}
	return models.Product{}, false
}
func DeleteProduct(id int) bool {
	for index, product := range models.Products {
		if product.ID == id {
			models.Products = append(models.Products[:index], models.Products[index+1:]...)
			return true
		}
	}
	return false
}
