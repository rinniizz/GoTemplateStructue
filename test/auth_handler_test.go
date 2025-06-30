package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-template-structure/internal/domain"
	"go-template-structure/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthService is a mock implementation of AuthService
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(req *domain.CreateUserRequest) (*domain.AuthResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*domain.AuthResponse), args.Error(1)
}

func (m *MockAuthService) Login(req *domain.LoginRequest) (*domain.AuthResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*domain.AuthResponse), args.Error(1)
}

func (m *MockAuthService) RefreshToken(refreshToken string) (*domain.AuthResponse, error) {
	args := m.Called(refreshToken)
	return args.Get(0).(*domain.AuthResponse), args.Error(1)
}

func TestAuthHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup
	mockService := new(MockAuthService)
	authHandler := handler.NewAuthHandler(mockService)

	// Test data
	request := &domain.CreateUserRequest{
		Email:     "test@example.com",
		Username:  "testuser",
		Password:  "password123",
		FirstName: "Test",
		LastName:  "User",
	}

	expectedResponse := &domain.AuthResponse{
		User: &domain.User{
			ID:        1,
			Email:     "test@example.com",
			Username:  "testuser",
			FirstName: "Test",
			LastName:  "User",
		},
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiresIn:    86400,
	}

	// Mock service call
	mockService.On("Register", request).Return(expectedResponse, nil)

	// Create request
	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Setup router
	router := gin.New()
	router.POST("/auth/register", authHandler.Register)

	// Perform request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var response domain.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response.Success)
	assert.Equal(t, "User registered successfully", response.Message)

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup
	mockService := new(MockAuthService)
	authHandler := handler.NewAuthHandler(mockService)

	// Test data
	request := &domain.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	expectedResponse := &domain.AuthResponse{
		User: &domain.User{
			ID:       1,
			Email:    "test@example.com",
			Username: "testuser",
		},
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiresIn:    86400,
	}

	// Mock service call
	mockService.On("Login", request).Return(expectedResponse, nil)

	// Create request
	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Setup router
	router := gin.New()
	router.POST("/auth/login", authHandler.Login)

	// Perform request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response domain.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response.Success)
	assert.Equal(t, "Login successful", response.Message)

	// Verify mock was called
	mockService.AssertExpectations(t)
}
