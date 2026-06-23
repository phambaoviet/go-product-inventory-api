Go Product Inventory API

A simple CRUD REST API for managing products, built with Go and Gin.

Tech Stack

* Go 1.25.5
* Gin Framework

Features

* CRUD products
* JSON request/response
* Basic validation
* In-memory data storage

API Endpoints

Base URL:

http://localhost:8080

Method	Endpoint	Description
GET	/api/products	Get all products
GET	/api/products/:id	Get product by ID
POST	/api/products	Create a product
PUT	/api/products/:id	Update a product
DELETE	/api/products/:id	Delete a product

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

Notes

This project uses an in-memory slice as a mock database. Data will be reset when the server restarts.