package routes

import (
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/usecase"

	"github.com/gorilla/mux"
)

func CategoryRoute(r *mux.Router, categoryHandler *usecase.CategoryHandler) {
	category := r.PathPrefix("/category").Subrouter()

	category.Use(middleware.AuthenticationMiddleware)
	category.Use(middleware.AuthenticationAdmin)
	category.HandleFunc("", categoryHandler.GetAll).Methods("GET")
}