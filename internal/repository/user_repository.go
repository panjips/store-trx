package repository

import (
	"store-trx-go/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}