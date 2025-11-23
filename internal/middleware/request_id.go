package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID adds a unique request ID to each request for tracking and debugging
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get request ID from header first
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			// Generate new UUID if not provided
			requestID = uuid.New().String()
		}

		// Store in context for later use
		c.Set("RequestID", requestID)

		// Add to response header
		c.Header("X-Request-ID", requestID)

		c.Next()
	}
}
