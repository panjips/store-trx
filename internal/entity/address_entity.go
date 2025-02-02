package entity

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID				uint	`gorm:"primaryKey"`
	AddressTitle	string
	RecipientName	string
	PhoneNumber		string
	DetailAddress	string
	UserID			uint		
	User			User	`gorm:"foreignKey:UserID"`
}