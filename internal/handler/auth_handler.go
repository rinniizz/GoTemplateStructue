package handler

import (
	"net/http"

	"go-template-structure/internal/domain"
	"go-template-structure/internal/service"
	"go-template-structure/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param user body domain.CreateUserRequest true "User registration data"
// @Success 201 {object} domain.APIResponse{data=domain.AuthResponse}
// @Failure 400 {object} domain.APIResponse
// @Failure 409 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	authResponse, err := h.authService.Register(&req)
	if err != nil {
		if err.Error() == "user with this email or username already exists" {
			utils.ErrorResponse(c, http.StatusConflict, "User already exists", err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user", err.Error())
		return
	}

	c.JSON(http.StatusCreated, domain.APIResponse{
		Success: true,
		Message: "User registered successfully",
		Data:    authResponse,
	})
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body domain.LoginRequest true "User login credentials"
// @Success 200 {object} domain.APIResponse{data=domain.AuthResponse}
// @Failure 400 {object} domain.APIResponse
// @Failure 401 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	authResponse, err := h.authService.Login(&req)
	if err != nil {
		if err.Error() == "invalid email or password" || err.Error() == "user account is inactive" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Authentication failed", err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to login", err.Error())
		return
	}

	utils.SuccessResponse(c, "Login successful", authResponse)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Get new access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param token body domain.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} domain.APIResponse{data=domain.AuthResponse}
// @Failure 400 {object} domain.APIResponse
// @Failure 401 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req domain.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	authResponse, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		if err.Error() == "invalid refresh token" || err.Error() == "user not found" || err.Error() == "user account is inactive" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Token refresh failed", err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to refresh token", err.Error())
		return
	}

	utils.SuccessResponse(c, "Token refreshed successfully", authResponse)
}
