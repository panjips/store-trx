package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID				uint		`gorm:"primaryKey"`
	Name			string
	Slug			string
	ResellerPrice	string
	CustomerPrice	string
	Stock			uint
	Description		string
	StoreID			uint		
	Store			Store		`gorm:"foreignKey:StoreID"`
	CategoryID		uint		
	Category		Category	`gorm:"foreignKey:CategoryID"`
	Photos			[]Photo		`gorm:"foreignKey:ProductID"`
}