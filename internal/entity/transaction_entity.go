package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID				uint		`gorm:"primaryKey"`
	TotalPrice		uint
	InvoiceCode		string
	PaymentMethod	string
	AddressID		uint		
	Address 		Address		`gorm:"foreignKey:AddressID"`
	UserID			uint
	User  			User		`gorm:"foreignKey:UserID"`
} 