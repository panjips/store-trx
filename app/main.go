package main

import (
	"fmt"
	"net/http"
	"os"
	"store-trx-go/internal/handler/routes"
	"store-trx-go/pkg/config"

	_ "store-trx-go/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

	handler := cors.AllowAll().Handler(r)
	port := os.Getenv("PORT")
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}