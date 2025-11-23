package test

import (
	"context"
	"testing"
	"time"

	"go-template-structure/internal/config"
	"go-template-structure/internal/domain"
	"go-template-structure/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) List(offset, limit int) ([]domain.User, int64, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]domain.User), args.Get(1).(int64), args.Error(2)
}

func (m *MockUserRepository) Exists(email, username string) (bool, error) {
	args := m.Called(email, username)
	return args.Bool(0), args.Error(1)
}

// MockRedisInterface is a mock implementation of RedisInterface
type MockRedisInterface struct {
	mock.Mock
}

func (m *MockRedisInterface) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockRedisInterface) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockRedisInterface) Del(ctx context.Context, keys ...string) error {
	args := m.Called(ctx, keys)
	return args.Error(0)
}

// TestUserService_GetProfile tests the GetProfile method
func TestUserService_GetProfile(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRedis := new(MockRedisInterface)
	jwtConfig := config.JWTConfig{Secret: "test-secret", Expiration: 24 * time.Hour}

	userService := service.NewUserService(mockRepo, mockRedis, jwtConfig)

	testUser := &domain.User{
		ID:        1,
		Email:     "test@example.com",
		Username:  "testuser",
		FirstName: "Test",
		LastName:  "User",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetByID", uint(1)).Return(testUser, nil).Once()

		result, err := userService.GetProfile(1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, testUser.Email, result.Email)
		assert.Equal(t, testUser.Username, result.Username)
		mockRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockRepo.On("GetByID", uint(999)).Return(nil, assert.AnError).Once()

		result, err := userService.GetProfile(999)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

// TestUserService_GetUsers tests the GetUsers method
func TestUserService_GetUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRedis := new(MockRedisInterface)
	jwtConfig := config.JWTConfig{Secret: "test-secret", Expiration: 24 * time.Hour}

	userService := service.NewUserService(mockRepo, mockRedis, jwtConfig)

	testUsers := []domain.User{
		{ID: 1, Email: "user1@example.com", Username: "user1"},
		{ID: 2, Email: "user2@example.com", Username: "user2"},
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("List", 0, 10).Return(testUsers, int64(2), nil).Once()

		result, pagination, err := userService.GetUsers(1, 10)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 2)
		assert.NotNil(t, pagination)
		assert.Equal(t, int64(2), pagination.Total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo.On("List", 0, 10).Return([]domain.User{}, int64(0), assert.AnError).Once()

		result, pagination, err := userService.GetUsers(1, 10)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, pagination)
		mockRepo.AssertExpectations(t)
	})
}

// TestUserService_DeleteUser tests the DeleteUser method
func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRedis := new(MockRedisInterface)
	jwtConfig := config.JWTConfig{Secret: "test-secret", Expiration: 24 * time.Hour}

	userService := service.NewUserService(mockRepo, mockRedis, jwtConfig)

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Delete", uint(1)).Return(nil).Once()

		err := userService.DeleteUser(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo.On("Delete", uint(999)).Return(assert.AnError).Once()

		err := userService.DeleteUser(999)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
