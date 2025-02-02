package database

import (
	"fmt"
	"rakamin-final/internal/entity"

	"gorm.io/gorm"
)

func MigrateEntities(DB *gorm.DB) {
	err := DB.AutoMigrate(
		&entity.User{},
		&entity.Store{},
		&entity.Address{},
		&entity.Category{},
		&entity.Product{},
		&entity.Photo{},
		&entity.ProductLog{},
		&entity.Transaction{},
		&entity.DetailTransaction{},
	)

	if err != nil {
		fmt.Printf("Failed to auto-migrate entities: %v\n", err)
	}

	fmt.Println("Successfully migrated entities")
}