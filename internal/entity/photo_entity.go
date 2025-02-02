package entity

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID			uint		`gorm:"primaryKey"`
	URL			string
	ProductID	uint		
	Product		Product		`gorm:"foreignKey:ProductID"`
}