package mock

import (
	"errors"

	"go-template-structure/internal/domain"
)

type MockUserRepository struct {
	users []domain.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: []domain.User{
			{
				ID:        1,
				Email:     "admin@example.com",
				Username:  "admin",
				Password:  "$2a$10$example.hash.for.password123", // bcrypt hash for "password123"
				FirstName: "Admin",
				LastName:  "User",
				IsActive:  true,
			},
		},
	}
}

func (r *MockUserRepository) Create(user *domain.User) error {
	user.ID = uint(len(r.users) + 1)
	r.users = append(r.users, *user)
	return nil
}

func (r *MockUserRepository) GetByID(id uint) (*domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *MockUserRepository) GetByEmail(email string) (*domain.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *MockUserRepository) GetByUsername(username string) (*domain.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *MockUserRepository) Update(user *domain.User) error {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = *user
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *MockUserRepository) Delete(id uint) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *MockUserRepository) List(offset, limit int) ([]domain.User, int64, error) {
	total := int64(len(r.users))

	if offset >= len(r.users) {
		return []domain.User{}, total, nil
	}

	end := offset + limit
	if end > len(r.users) {
		end = len(r.users)
	}

	return r.users[offset:end], total, nil
}

func (r *MockUserRepository) Exists(email, username string) (bool, error) {
	for _, user := range r.users {
		if user.Email == email || user.Username == username {
			return true, nil
		}
	}
	return false, nil
}
