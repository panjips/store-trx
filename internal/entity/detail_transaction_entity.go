package entity

import (
	"gorm.io/gorm"
)

type DetailTransaction struct {
	gorm.Model
	ID 				uint		`gorm:"primaryKey"`
	Quantity		uint
	TotalPrice		uint
	TransactionID	uint		
	Transaction		Transaction	`gorm:"foreignKey:TransactionID"`
	ProductLogID	uint		
	ProductLog		ProductLog	`gorm:"foreignKey:ProductLogID"`
	StoreID			uint		
	Store			Store		`gorm:"foreignKey:StoreID"`
}