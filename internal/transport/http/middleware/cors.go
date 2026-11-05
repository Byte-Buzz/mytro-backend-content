package middleware

import (
	"mytro-backend-content/internal/infrastructure/config"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS applies CORS middleware to the Gin router.
func CORS(cors config.CORSConfig) gin.HandlerFunc {
	// This function returns a Gin handler function that processes incoming HTTP requests and sets CORS headers based on the provided configuration.
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Determine if the origin is allowed
		var allowedOrigin string
		if slices.Contains(cors.AllowedOrigins, "*") {
			allowedOrigin = "*"
		} else if origin != "" && slices.Contains(cors.AllowedOrigins, origin) {
			allowedOrigin = origin
		}

		if allowedOrigin == "" && len(cors.AllowedOrigins) > 0 {
			// Origin not allowed
			c.AbortWithStatus(403)
			return
		}

		// Set CORS headers
		if allowedOrigin != "" {
			c.Header("Access-Control-Allow-Origin", allowedOrigin)
		}
		c.Header("Access-Control-Allow-Methods", strings.Join(cors.AllowedMethods, ", "))
		c.Header("Access-Control-Allow-Headers", strings.Join(cors.AllowedHeaders, ", "))
		if cors.AllowCredentials {
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No Content
			return
		}

		c.Next()
	}
}
