package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-template-structure/internal/config"
	"go-template-structure/internal/domain"
	"go-template-structure/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RedisInterface defines methods for Redis operations
type RedisInterface interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
}

type UserService interface {
	CreateUser(req *domain.CreateUserRequest) (*domain.User, error)
	GetUser(id uint) (*domain.User, error)
	GetUsers(page, limit int) ([]domain.User, *domain.PaginationResponse, error)
	UpdateUser(id uint, req *domain.UpdateUserRequest) (*domain.User, error)
	DeleteUser(id uint) error
	GetProfile(userID uint) (*domain.User, error)
	UpdateProfile(userID uint, req *domain.UpdateUserRequest) (*domain.User, error)
}

type userService struct {
	userRepo    repository.UserRepository
	redisClient RedisInterface // ใช้ interface แทน
	jwtConfig   config.JWTConfig
}

func NewUserService(userRepo repository.UserRepository, redisClient interface{}, jwtConfig config.JWTConfig) UserService {
	// Type assertion เพื่อให้ใช้ interface ได้
	var redisInterface RedisInterface
	if client, ok := redisClient.(RedisInterface); ok {
		redisInterface = client
	}

	return &userService{
		userRepo:    userRepo,
		redisClient: redisInterface,
		jwtConfig:   jwtConfig,
	}
}

func (s *userService) CreateUser(req *domain.CreateUserRequest) (*domain.User, error) {
	// Check if user already exists
	exists, err := s.userRepo.Exists(req.Email, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}

	if exists {
		return nil, errors.New("user with this email or username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &domain.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsActive:  true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Cache user
	s.cacheUser(user)

	return user, nil
}

func (s *userService) GetUser(id uint) (*domain.User, error) {
	// Try to get from cache first
	if user := s.getUserFromCache(id); user != nil {
		return user, nil
	}

	// Get from database
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Cache user
	s.cacheUser(user)

	return user, nil
}

func (s *userService) GetUsers(page, limit int) ([]domain.User, *domain.PaginationResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	users, total, err := s.userRepo.List(offset, limit)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get users: %w", err)
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	pagination := &domain.PaginationResponse{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return users, pagination, nil
}

func (s *userService) UpdateUser(id uint, req *domain.UpdateUserRequest) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Update fields if provided
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Update cache
	s.cacheUser(user)

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	if err := s.userRepo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	// Remove from cache
	s.removeUserFromCache(id)

	return nil
}

func (s *userService) GetProfile(userID uint) (*domain.User, error) {
	return s.GetUser(userID)
}

func (s *userService) UpdateProfile(userID uint, req *domain.UpdateUserRequest) (*domain.User, error) {
	return s.UpdateUser(userID, req)
}

// Cache operations
func (s *userService) cacheUser(user *domain.User) {
	if s.redisClient == nil {
		return
	}

	ctx := context.Background()
	key := fmt.Sprintf("user:%d", user.ID)

	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	s.redisClient.Set(ctx, key, string(data), 30*time.Minute)
}

func (s *userService) getUserFromCache(id uint) *domain.User {
	if s.redisClient == nil {
		return nil
	}

	ctx := context.Background()
	key := fmt.Sprintf("user:%d", id)

	data, err := s.redisClient.Get(ctx, key)
	if err != nil {
		return nil
	}

	var user domain.User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil
	}

	return &user
}

func (s *userService) removeUserFromCache(id uint) {
	if s.redisClient == nil {
		return
	}

	ctx := context.Background()
	key := fmt.Sprintf("user:%d", id)
	s.redisClient.Del(ctx, key)
}
