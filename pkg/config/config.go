package config

import (
	"fmt"
	db "store-trx-go/pkg/database"
	"store-trx-go/pkg/r2"

	"gorm.io/gorm"
)


func InitConfig() *gorm.DB {
	DB, err := db.InitializeDB()
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
	}
	db.MigrateEntities(DB)
	fmt.Println("Successfully initialized database")

	r2.InitR2Client()

	return DB
}