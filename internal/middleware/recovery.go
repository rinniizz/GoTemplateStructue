package middleware

import (
	"net/http"

	"go-template-structure/internal/domain"
	"go-template-structure/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Recovery middleware for panic recovery
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.WithFields(map[string]interface{}{
			"error":  recovered,
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"ip":     c.ClientIP(),
		}).Error("Panic recovered")

		c.JSON(http.StatusInternalServerError, domain.APIResponse{
			Success: false,
			Message: "Internal server error",
			Error:   "An unexpected error occurred",
		})
	})
}
