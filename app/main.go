package main

import (
	"fmt"
	"rakamin-final/pkg/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load();
	err != nil {
		fmt.Println("Error loading .env file")
	}
	config.InitConfig()
	
}