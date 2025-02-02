package entity

import (
	"gorm.io/gorm"
)

type ProductLog struct {
	gorm.Model
	ID				uint		`gorm:"primaryKey"`
	Name			string
	Slug			string
	ResellerPrice	string
	CustomerPrice	string
	Description		string
	ProductID		uint		
	Product			Product		`gorm:"foreignKey:ProductID"`
	StoreID			uint		
	Store			Store		`gorm:"foreignKey:StoreID"`
	CategoryID		uint		
	Category		Category	`gorm:"foreignKey:CategoryID"`
}