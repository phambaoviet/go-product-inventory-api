package models

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Slug     string  `json:"slug"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Quantity int     `json:"quantity" binding:"required,gte=0"`
}
