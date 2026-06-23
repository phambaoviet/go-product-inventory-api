package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationError {
			switch e.Tag() {
			case "required":
				errors[e.Field()] = e.Field() + " is required"
			case "gt":
				errors[e.Field()] = e.Field() + " must be greater than " + e.Param()
			case "gte":
				errors[e.Field()] = e.Field() + " must be greater than or equal le " + e.Param()
			}
		}
		return gin.H{"errors": errors}
	}
	return gin.H{"error": "invalid request body"}
}
