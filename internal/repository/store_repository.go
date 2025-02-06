package repository

import (
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"
	"store-trx-go/pkg/database"

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

func (r *StoreRepository) FindByID(ID uint) (*dto.StoreDTO, error) {
	var store dto.StoreDTO
	err := r.db.Model(&entity.Store{}).Where("id = ?", ID).Scan(&store).Error
	if err != nil {
		return nil, err
	}

	return &store, nil
}

func (r *StoreRepository) FindByUserID(userID uint) (*dto.StoreDTO, error) {
	var store dto.StoreDTO
	err := r.db.Model(&entity.Store{}).Where("user_id = ?", userID).Scan(&store).Error
	if err != nil {
		return nil, err
	}

	return &store, nil
}
func (r *StoreRepository) Update(ID uint, store *entity.Store) error {
	err := r.db.Model(&entity.Store{}).Where("id = ?", ID).Updates(store).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *StoreRepository) GetAll(page int, limit int) ([]dto.StoreDTO, error) {
	var stores []dto.StoreDTO
	
	paginate := database.NewPaginate(limit, page).PaginatedResult
	err := r.db.Scopes(paginate).Model(&entity.Store{}).Scan(&stores).Error

	if err != nil {
		return nil, err
	}

	return stores, nil
}
