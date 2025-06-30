package service

import (
	"errors"
	"fmt"
	"time"

	"go-template-structure/internal/config"
	"go-template-structure/internal/domain"
	"go-template-structure/internal/repository"
	"go-template-structure/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req *domain.CreateUserRequest) (*domain.AuthResponse, error)
	Login(req *domain.LoginRequest) (*domain.AuthResponse, error)
	RefreshToken(refreshToken string) (*domain.AuthResponse, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtConfig config.JWTConfig
}

func NewAuthService(userRepo repository.UserRepository, jwtConfig config.JWTConfig) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtConfig: jwtConfig,
	}
}

func (s *authService) Register(req *domain.CreateUserRequest) (*domain.AuthResponse, error) {
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

	// Generate tokens
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtConfig.Secret, s.jwtConfig.Expiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtConfig.Secret, 7*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &domain.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.jwtConfig.Expiration.Seconds()),
	}, nil
}

func (s *authService) Login(req *domain.LoginRequest) (*domain.AuthResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate tokens
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtConfig.Secret, s.jwtConfig.Expiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtConfig.Secret, 7*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &domain.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.jwtConfig.Expiration.Seconds()),
	}, nil
}

func (s *authService) RefreshToken(refreshToken string) (*domain.AuthResponse, error) {
	// Validate refresh token
	claims, err := utils.ValidateJWT(refreshToken, s.jwtConfig.Secret)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Get user
	user, err := s.userRepo.GetByID(claims.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}

	// Generate new tokens
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtConfig.Secret, s.jwtConfig.Expiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	newRefreshToken, err := utils.GenerateJWT(user.ID, user.Email, s.jwtConfig.Secret, 7*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &domain.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    int64(s.jwtConfig.Expiration.Seconds()),
	}, nil
}
