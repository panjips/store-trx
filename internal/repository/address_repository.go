package repository

import (
	"store-trx-go/internal/entity"

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

func (r *AddressRepository) FindByID(ID uint) (*entity.Address, error){
	var address entity.Address
	err := r.db.Where("id = ?", ID).First(&address).Error
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *AddressRepository) FindByUserID(userID uint) ([]entity.Address, error) {
	var addresses []entity.Address
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
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

func (r *AddressRepository) Delete(address *entity.Address) error {
	err := r.db.Delete(address).Error
	if err != nil {
		return err
	}

	return nil
}