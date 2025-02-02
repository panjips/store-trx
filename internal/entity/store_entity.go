package entity

import (
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	ID			uint		`gorm:"primaryKey"`
	Name		string
	ImageURL	string
	UserID		uint		
	User		User		`gorm:"foreignKey:UserID"`
}