package entity

import (
	"gorm.io/gorm"
)

type DetailTransaction struct {
	gorm.Model
	ID 				uint			`gorm:"primaryKey"`
	Quantity		uint
	TotalPrice		uint
	ProductID		uint		
	Product			Product		`gorm:"foreignKey:ProductID"`
	TransactionID	uint		
	Transaction		Transaction		`gorm:"foreignKey:TransactionID"`
}