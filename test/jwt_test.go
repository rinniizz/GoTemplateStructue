package test

import (
	"testing"
	"time"

	"go-template-structure/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	// Test data
	userID := uint(1)
	email := "test@example.com"
	secretKey := "test-secret-key"
	expiration := time.Hour

	// Generate JWT
	token, err := utils.GenerateJWT(userID, email, secretKey, expiration)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateJWT(t *testing.T) {
	// Test data
	userID := uint(1)
	email := "test@example.com"
	secretKey := "test-secret-key"
	expiration := time.Hour

	// Generate JWT
	token, err := utils.GenerateJWT(userID, email, secretKey, expiration)
	assert.NoError(t, err)

	// Validate JWT
	claims, err := utils.ValidateJWT(token, secretKey)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	// Test with invalid token
	secretKey := "test-secret-key"
	invalidToken := "invalid.token.here"

	// Validate JWT
	claims, err := utils.ValidateJWT(invalidToken, secretKey)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateJWT_WrongSecret(t *testing.T) {
	// Test data
	userID := uint(1)
	email := "test@example.com"
	secretKey := "test-secret-key"
	wrongSecret := "wrong-secret-key"
	expiration := time.Hour

	// Generate JWT with correct secret
	token, err := utils.GenerateJWT(userID, email, secretKey, expiration)
	assert.NoError(t, err)

	// Validate JWT with wrong secret
	claims, err := utils.ValidateJWT(token, wrongSecret)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, claims)
}
