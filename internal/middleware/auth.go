package middleware

import (
	"net/http"
	"strings"

	"go-template-structure/pkg/utils"

	"github.com/gin-gonic/gin"
)

// JWTAuth middleware for JWT authentication
func JWTAuth(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization header required", "missing authorization header")
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header format", "expected format: Bearer <token>")
			c.Abort()
			return
		}

		token := tokenParts[1]

		// Validate JWT token
		claims, err := utils.ValidateJWT(token, secretKey)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token", err.Error())
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}
