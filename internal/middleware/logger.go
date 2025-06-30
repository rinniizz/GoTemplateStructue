package middleware

import (
	"github.com/gin-gonic/gin"
)

// Logger middleware for request logging
func Logger() gin.HandlerFunc {
	return gin.Logger()
}
