package repository

import (
	"store-trx-go/internal/entity"
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

func (r *StoreRepository) FindByID(ID uint) (*entity.Store, error) {
	var store entity.Store
	err := r.db.Where("id = ?", ID).First(&store).Error
	if err != nil {
		return nil, err
	}

	return &store, nil
}

func (r *StoreRepository) FindByUserID(userID uint) (*entity.Store, error) {
	var store entity.Store
	err := r.db.Where("user_id = ?", userID).First(&store).Error
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

func (r *StoreRepository) GetAll(page int, limit int) ([]entity.Store, error) {
	var stores []entity.Store
	
	paginate := database.NewPaginate(limit, page).PaginatedResult
	err := r.db.Scopes(paginate).Find(&stores).Error

	if err != nil {
		return nil, err
	}

	return stores, nil
}
