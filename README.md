Go Product Inventory API

A simple CRUD REST API for managing products, built with Go and Gin.

Tech Stack

* Go 1.25.5
* Gin Framework

Features

* CRUD products
* JSON request/response
* Product validation
* Auto-generate product slug from name
* Query filter by name and slug
* Service layer for product business logic
* In-memory data storage

API Endpoints

Base URL:

http://localhost:8080

Method	Endpoint	Description
GET	/api/products	Get all products
GET	/api/products?name=keyboard	Filter products by name
GET	/api/products?slug=mechanical-keyboard	Filter products by slug
GET	/api/products/:id	Get product by ID
POST	/api/products	Create a product
PUT	/api/products/:id	Update a product
DELETE	/api/products/:id	Delete a product

Project Structure

go-product-inventory-api/
├── main.go
├── routes/
│   └── product_routes.go
├── handlers/
│   └── product_handler.go
├── services/
│   └── product_service.go
├── models/
│   └── product.go
├── utils/
│   └── validation.go
├── go.mod
└── README.md

Folder Responsibilities

main.go      -> start the Gin server
routes/      -> define API routes
handlers/    -> handle HTTP requests and responses
services/    -> handle product business logic
models/      -> define product model and mock data
utils/       -> helper functions for validation and slug generation

Getting Started

Install dependencies:

go mod tidy

Run the project:

go run .

The server will start at:

http://localhost:8080

Sample Product JSON

{
"name": "Keyboard",
"price": 25.99,
"quantity": 10
}

After creating a product, the API will automatically generate an id and slug.

Example response:

{
"id": 1,
"name": "Keyboard",
"slug": "keyboard",
"price": 25.99,
"quantity": 10
}

Query Examples

Filter products by name:

GET /api/products?name=keyboard

Filter products by slug:

GET /api/products?slug=mechanical-keyboard

Notes

This project uses an in-memory slice as a mock database. Data will be reset when the server restarts.