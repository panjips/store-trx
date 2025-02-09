package repository

import (
	"fmt"
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"
	"time"

	"gorm.io/gorm"
)

type TrxRepository struct {
	db *gorm.DB
}

func NewTrxRepository(db *gorm.DB) *TrxRepository {
	return &TrxRepository{
		db: db,
	}
}

func (r *TrxRepository) Create(trx *entity.Transaction) error{
	err := r.db.Create(&trx).Error
	if err != nil {
		return err
	}

	datePart := time.Now().Format("20060102")
	trx.InvoiceCode = fmt.Sprintf("INV-%s-%04d", datePart, trx.ID)

	return nil
}

func (r *TrxRepository) Update(trx *entity.Transaction) error{
	err := r.db.Updates(trx).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TrxRepository) CreateDetailTrx(trx *entity.DetailTransaction) error{
	err := r.db.Create(&trx).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TrxRepository) GetAll() ([]*dto.TrxDTO, error){
	var transaction []entity.Transaction
	err := r.db.Preload("User").Preload("Address").Preload("DetailTransaction").Find(&transaction).Error
	if err != nil {
		return nil, err
	}

	var trx []*dto.TrxDTO
	for _, data := range transaction {
		var detailTrx []dto.DetailTrxDTO
		for _, detail := range data.DetailTransaction {
			detailTrx = append(detailTrx, dto.DetailTrxDTO{
				ProductID: detail.ProductID,
				Quantity: detail.Quantity,
				TotalPrice: detail.TotalPrice,
			})
		}

		
		trx = append(trx, &dto.TrxDTO{
			ID: data.ID,
			PaymentMethod: data.PaymentMethod,
			AddressID: data.AddressID,
			TotalPrice: data.TotalPrice,
			DetailTrx: detailTrx,
		})
	}
	

	return trx, nil
}