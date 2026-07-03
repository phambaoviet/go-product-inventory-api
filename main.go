package main

import (
	"go-product-inventory-api/config"
	"go-product-inventory-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
	}
	config.Init()
	r := gin.Default()

	routes.RegisterProductRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server cannot start...: ", err)
	}
}
