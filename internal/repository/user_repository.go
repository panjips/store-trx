package repository

import (
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"

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

func (r *UserRepository) Login(phoneNumber string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByID(ID uint) (*dto.UserDTO, error) {
	var user dto.UserDTO
	err := r.db.Model(&entity.User{}).First(&user, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateProfile(ID uint, user *entity.User) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", ID).Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}