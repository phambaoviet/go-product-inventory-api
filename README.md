Go Product Inventory API

A lightweight RESTful API built with Golang and the Gin Framework to manage a simple product inventory.

This project demonstrates core backend concepts such as routing, HTTP methods, JSON request/response handling, path parameters, validation, error handling, and in-memory CRUD operations.

Features

* GET /api/products - Retrieve all products.
* GET /api/products/:id - Get a product by ID.
* POST /api/products - Create a new product.
* PUT /api/products/:id - Update an existing product.
* DELETE /api/products/:id - Delete a product by ID.

Core Concepts Applied

* RESTful API Design: Used HTTP methods such as GET, POST, PUT, and DELETE to represent CRUD operations.
* Gin Routing: Built API routes using the Gin framework.
* JSON Handling: Bound incoming JSON request bodies to Go structs.
* Path Parameters: Retrieved product IDs from URL parameters.
* Type Conversion: Used strconv.Atoi to convert path parameters from string to integer.
* Error Handling: Returned appropriate error responses for invalid input, missing resources, and unsupported operations.
* Slice Manipulation: Used Go slices as an in-memory data store, including deleting items with append(products[:i], products[i+1:]...).
* Struct Tags: Used JSON struct tags such as json:"name" to control JSON field names.

Tech Stack

* Go
* Gin Framework
* In-memory data storage using slices

Project Structure

go-product-inventory-api/
├── main.go
├── go.mod
├── go.sum
└── README.md

Prerequisites

Make sure you have Go installed:

go version

Installation

Clone the repository:

git clone https://github.com/phambaoviet/go-product-inventory-api.git
cd go-product-inventory-api

Install dependencies:

go mod tidy

Run the application:

go run main.go

The server will run at:

http://localhost:8080

API Endpoints

Get all products

curl http://localhost:8080/api/products

Get product by ID

curl http://localhost:8080/api/products/1

Create product

curl -X POST http://localhost:8080/api/products \
-H "Content-Type: application/json" \
-d '{"name":"Keyboard","price":25.99,"quantity":10}'

Update product

curl -X PUT http://localhost:8080/api/products/1 \
-H "Content-Type: application/json" \
-d '{"name":"Mechanical Keyboard","price":59.99,"quantity":5}'

Delete product

curl -X DELETE http://localhost:8080/api/products/1

Notes

This project uses an in-memory slice as a mock database. All data will be reset when the server restarts.

The next improvement for this project could be connecting it to a real database such as PostgreSQL.