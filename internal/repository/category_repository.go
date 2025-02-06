package repository

import (
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]dto.CategoryDTO, error) {
	var categories []dto.CategoryDTO
	err := r.db.Model(&entity.Category{}).Scan(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(ID uint) (*dto.CategoryDTO, error) {
	var category dto.CategoryDTO
	err := r.db.Model(&entity.Category{}).Where("id = ?", ID).Scan(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Create(category *entity.Category) error {
	err := r.db.Create(category).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) Update(ID uint, category *entity.Category) error {
	err := r.db.Model(&entity.Category{}).Where("id = ?", ID).Updates(category).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) Delete(ID uint) error {
	err := r.db.Delete(&entity.Category{}, ID).Error
	if err != nil {
		return err
	}

	return nil
}