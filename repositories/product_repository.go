package repositories

import (
	"go-product-inventory-api/config"
	"go-product-inventory-api/models"
)

func GetAllProducts() ([]models.Product, error) {
	rows, err := config.DB.Query("SELECT id, name, slug, price, quantity FROM products ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []models.Product{}

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Slug,
			&product.Price,
			&product.Quantity,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
func GetProductByID(id int) (models.Product, error) {
	var product models.Product
	row := config.DB.QueryRow("SELECT id, name, slug, price, quantity FROM products WHERE id = $1", id)
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Slug,
		&product.Price,
		&product.Quantity,
	)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	row := config.DB.QueryRow("INSERT INTO products (name, slug, price, quantity) VALUES ($1,$2,$3,$4) RETURNING id",
		product.Name, product.Slug, product.Price, product.Quantity)
	err := row.Scan(&product.ID)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
func UpdateProduct(id int, product models.Product) (models.Product, error) {
	row := config.DB.QueryRow("UPDATE products SET name = $1, slug = $2, price = $3, quantity = $4"+
		" WHERE id = $5 RETURNING id, name, slug, price, quantity",
		product.Name, product.Slug, product.Price, product.Quantity, id)
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Slug,
		&product.Price,
		&product.Quantity)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
func DeleteProduct(id int) (bool, error) {
	result, err := config.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
