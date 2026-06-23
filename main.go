package main

import (
	"go-product-inventory-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterProductRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server cannot start...: ", err)
	}
}
