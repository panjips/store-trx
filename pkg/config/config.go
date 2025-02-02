package config

import (
	"fmt"
	db "rakamin-final/pkg/database"
)

func InitConfig() {
	DB, err := db.InitializeDB()
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
	}

	db.MigrateEntities(DB)
	fmt.Println("Successfully initialized database")
}