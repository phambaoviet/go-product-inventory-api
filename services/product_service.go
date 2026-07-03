package services

import (
	"go-product-inventory-api/models"
	"go-product-inventory-api/repositories"
	"go-product-inventory-api/utils"
	"strings"
)

type ProductFilter struct {
	Slug string
	Name string
}

func GetProducts(filter ProductFilter) ([]models.Product, error) {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return nil, err
	}
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
	return filteredProduct, nil
}
func GetProductByID(id int) (models.Product, error) {
	return repositories.GetProductByID(id)
}
func CreateProduct(product models.Product) (models.Product, error) {
	product.Slug = utils.GenerateSlug(product.Name)
	return repositories.CreateProduct(product)
}
func UpdateProduct(id int, product models.Product) (models.Product, error) {
	product.Slug = utils.GenerateSlug(product.Name)
	return repositories.UpdateProduct(id, product)
}
func DeleteProduct(id int) (bool, error) {
	return repositories.DeleteProduct(id)
}
