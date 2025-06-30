package utils

import (
	"net/http"

	"go-template-structure/internal/domain"

	"github.com/gin-gonic/gin"
)

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, domain.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message string, error interface{}) {
	c.JSON(statusCode, domain.APIResponse{
		Success: false,
		Message: message,
		Error:   error,
	})
}

// GetUserIDFromContext extracts user ID from gin context
func GetUserIDFromContext(c *gin.Context) uint {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}

	if id, ok := userID.(uint); ok {
		return id
	}

	return 0
}

// GetUserEmailFromContext extracts user email from gin context
func GetUserEmailFromContext(c *gin.Context) string {
	email, exists := c.Get("user_email")
	if !exists {
		return ""
	}

	if emailStr, ok := email.(string); ok {
		return emailStr
	}

	return ""
}
