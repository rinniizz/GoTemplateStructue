package handler

import (
	"net/http"
	"strconv"

	"go-template-structure/internal/domain"
	"go-template-structure/internal/service"
	"go-template-structure/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get current user profile
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.APIResponse{data=domain.User}
// @Failure 401 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /users/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get profile", err.Error())
		return
	}

	utils.SuccessResponse(c, "Profile retrieved successfully", user)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update current user profile
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body domain.UpdateUserRequest true "User update data"
// @Success 200 {object} domain.APIResponse{data=domain.User}
// @Failure 400 {object} domain.APIResponse
// @Failure 401 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /users/profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)

	var req domain.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	user, err := h.userService.UpdateProfile(userID, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update profile", err.Error())
		return
	}

	utils.SuccessResponse(c, "Profile updated successfully", user)
}

// GetUsers godoc
// @Summary Get users list
// @Description Get paginated list of users
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} domain.APIResponse{data=domain.UserListResponse}
// @Failure 401 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	users, pagination, err := h.userService.GetUsers(page, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get users", err.Error())
		return
	}

	response := domain.UserListResponse{
		Users:      users,
		Pagination: pagination,
	}

	utils.SuccessResponse(c, "Users retrieved successfully", response)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} domain.APIResponse{data=domain.User}
// @Failure 400 {object} domain.APIResponse
// @Failure 401 {object} domain.APIResponse
// @Failure 404 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	user, err := h.userService.GetUser(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			utils.ErrorResponse(c, http.StatusNotFound, "User not found", err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get user", err.Error())
		return
	}

	utils.SuccessResponse(c, "User retrieved successfully", user)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param user body domain.UpdateUserRequest true "User update data"
// @Success 200 {object} domain.APIResponse{data=domain.User}
// @Failure 400 {object} domain.APIResponse
// @Failure 401 {object} domain.APIResponse
// @Failure 404 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	var req domain.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	user, err := h.userService.UpdateUser(uint(id), &req)
	if err != nil {
		if err.Error() == "user not found" {
			utils.ErrorResponse(c, http.StatusNotFound, "User not found", err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	utils.SuccessResponse(c, "User updated successfully", user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} domain.APIResponse
// @Failure 400 {object} domain.APIResponse
// @Failure 401 {object} domain.APIResponse
// @Failure 404 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	err = h.userService.DeleteUser(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			utils.ErrorResponse(c, http.StatusNotFound, "User not found", err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}

	utils.SuccessResponse(c, "User deleted successfully", nil)
}
