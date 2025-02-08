package repository

import (
	"store-trx-go/internal/entity"

	"gorm.io/gorm"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository (db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db: db,
	}
}

func (r *PhotoRepository) Create(photo *entity.Photo) error {
	err := r.db.Create(photo).Error
	if err != nil {
		return err
	}

	return nil
}