package middleware

import (
	"time"

	"go-template-structure/pkg/logger"

	"github.com/gin-gonic/gin"
)

// AuditLog logs all important actions for security auditing
func AuditLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Get user info if authenticated
		userID, _ := c.Get("user_id")
		email, _ := c.Get("email")
		requestID, _ := c.Get("RequestID")

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// Only log important actions
		if shouldAuditLog(method, path, statusCode) {
			logData := map[string]interface{}{
				"request_id": requestID,
				"method":     method,
				"path":       path,
				"status":     statusCode,
				"latency_ms": latency.Milliseconds(),
				"ip":         c.ClientIP(),
				"user_agent": c.Request.UserAgent(),
			}

			// Add user info if available
			if userID != nil {
				logData["user_id"] = userID
			}
			if email != nil {
				logData["email"] = email
			}

			// Log based on status
			if statusCode >= 500 {
				logger.Error("Audit Log - Server Error", logData)
			} else if statusCode >= 400 {
				logger.Warn("Audit Log - Client Error", logData)
			} else {
				logger.Info("Audit Log", logData)
			}
		}
	}
}

// shouldAuditLog determines if the request should be logged
func shouldAuditLog(method, path string, status int) bool {
	// Log all POST, PUT, DELETE requests (data modifications)
	if method == "POST" || method == "PUT" || method == "DELETE" {
		return true
	}

	// Log all errors (4xx, 5xx)
	if status >= 400 {
		return true
	}

	// Log authentication/authorization endpoints
	authPaths := []string{
		"/api/v1/auth/login",
		"/api/v1/auth/register",
		"/api/v1/auth/refresh",
		"/api/v1/auth/logout",
	}
	for _, authPath := range authPaths {
		if path == authPath {
			return true
		}
	}

	// Don't log health checks and swagger
	if path == "/health" || path == "/swagger/index.html" {
		return false
	}

	return false
}
