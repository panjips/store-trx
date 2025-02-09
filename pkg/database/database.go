package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host		string
	Port		string
	User		string
	Password	string
	DBName		string
}

func (config *DBConfig) BuildDSN() string { 
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	);
}

func InitializeDB() (*gorm.DB, error) {
	dbConfig := DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
	}


	dsn := dbConfig.BuildDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{});
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}