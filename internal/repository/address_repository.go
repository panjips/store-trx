package repository

import (
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"

	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (r *AddressRepository) Create(address *entity.Address) error {
	err := r.db.Create(address).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *AddressRepository) FindByID(ID uint) (*dto.AddressDTO, error){
	var address dto.AddressDTO
	err := r.db.Model(&entity.Address{}).First(&address, ID).Error
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *AddressRepository) FindByUserID(userID uint) ([]dto.AddressDTO, error) {
	var addresses []dto.AddressDTO
	err := r.db.Model(&entity.Address{}).Find(&addresses, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (r *AddressRepository) Update(ID uint, address *entity.Address) error {
	err := r.db.Model(&entity.Address{}).Where("id = ?", ID).Updates(address).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *AddressRepository) Delete(ID uint) error {
	err := r.db.Delete(&entity.Address{}, ID).Error
	if err != nil {
		return err
	}

	return nil
}