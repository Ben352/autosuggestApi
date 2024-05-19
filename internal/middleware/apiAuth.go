package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateAPIKeyMiddleware(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestAPIKey := c.GetHeader("X-API-KEY")
		if requestAPIKey != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
