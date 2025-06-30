package domain

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"` // Never return password in JSON
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Avatar    string         `json:"avatar"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Username  string `json:"username" binding:"required,min=3,max=50"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Email     string `json:"email" binding:"omitempty,email"`
	Username  string `json:"username" binding:"omitempty,min=3,max=50"`
	FirstName string `json:"first_name" binding:"omitempty"`
	LastName  string `json:"last_name" binding:"omitempty"`
	Avatar    string `json:"avatar" binding:"omitempty"`
}

// LoginRequest represents the request payload for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse represents the response for authentication
type AuthResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// RefreshTokenRequest represents the request payload for token refresh
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// PaginationResponse represents pagination metadata
type PaginationResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// UserListResponse represents the response for user list
type UserListResponse struct {
	Users      []User              `json:"users"`
	Pagination *PaginationResponse `json:"pagination"`
}
