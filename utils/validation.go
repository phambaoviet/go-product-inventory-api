package utils

import (
	"regexp"
	"strings"

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
func GenerateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.TrimSpace(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
func IsValidSlug(slug string) bool {
	if slug == "" {
		return true
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]+$", slug)
	if !matched {
		return false
	}
	onlyNumber, _ := regexp.MatchString("^[0-9]+$", slug)
	return !onlyNumber
}
func IsValidSearchName(name string) bool {
	return len(name) <= 100
}
