package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApikeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-Api-Key")
		expectedAPIKey := os.Getenv("API_KEY")
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No X-api-Key"})
			return
		}
		if expectedAPIKey == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "API key is not configured"})
			return
		}
		if apiKey != expectedAPIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid X-Api-Key"})
			return
		}
		c.Next()
	}
}
