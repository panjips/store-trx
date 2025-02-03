package main

import (
	"fmt"
	"net/http"
	"store-trx-go/internal/handler/routes"
	"store-trx-go/pkg/config"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load();
	err != nil {
		fmt.Println("Error loading .env file")
	}
	db := config.InitConfig()

	r := mux.NewRouter()

	// r.Use(middleware.AuthenticationMiddleware)
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}).Methods("GET")

	routes.SetupRoutes(r.PathPrefix("/api/v1").Subrouter(), db)

	http.ListenAndServe(":8000", r)
}