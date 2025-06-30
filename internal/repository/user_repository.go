package repository

import (
	"go-template-structure/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id uint) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
	List(offset, limit int) ([]domain.User, int64, error)
	Exists(email, username string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *userRepository) List(offset, limit int) ([]domain.User, int64, error) {
	var users []domain.User
	var total int64

	// Count total records
	if err := r.db.Model(&domain.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) Exists(email, username string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.User{}).
		Where("email = ? OR username = ?", email, username).
		Count(&count).Error

	return count > 0, err
}
