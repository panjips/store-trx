package repository

import (
	"store-trx-go/internal/entity"

	"gorm.io/gorm"
)

type StoreRepository struct{
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{db: db}
}

func (r *StoreRepository) Create(store *entity.Store) error {
	err := r.db.Create(store).Error
	if err != nil {
		return err
	}

	return nil
}