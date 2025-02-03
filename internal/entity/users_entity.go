package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	ID			uint		`gorm:"primaryKey"`
	Name		string		
	Password	string
	PhoneNumber	string		`gorm:"unique"`
	BirthDate	*time.Time
	Gender		*string
	About		*string		
	Work		string
	Email		string		`gorm:"unique"`
	ProvinceID	string		
	CityID		string
	IsAdmin		bool		`gorm:"default:false"`
}