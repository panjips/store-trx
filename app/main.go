package main

import (
	"fmt"
	"net/http"
	"store-trx-go/internal/handler/routes"
	"store-trx-go/pkg/config"

	_ "store-trx-go/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	if err := godotenv.Load();
	err != nil {
		fmt.Println("Error loading .env file")
	}
	db := config.InitConfig()

	r := mux.NewRouter()
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	routes.SetupRoutes(r.PathPrefix("/api/v1").Subrouter(), db)
	http.ListenAndServe(":8000", r)
}