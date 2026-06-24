package models

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Slug     string `json:"slug"`
	Price    int    `json:"price" binding:"required,gt=0"`
	Quantity int    `json:"quantity" binding:"required,gte=0"`
}

var Products = []Product{
	{ID: 1, Name: "Apple", Price: 100, Quantity: 1},
	{ID: 2, Name: "Banana", Price: 200, Quantity: 2},
	{ID: 3, Name: "Orange", Price: 300, Quantity: 3},
}
