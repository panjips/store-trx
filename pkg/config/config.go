package config

import (
	"fmt"
	db "store-trx-go/pkg/database"

	"gorm.io/gorm"
)

func InitConfig() *gorm.DB {
	DB, err := db.InitializeDB()
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
	}

	db.MigrateEntities(DB)
	fmt.Println("Successfully initialized database")

	return DB
}