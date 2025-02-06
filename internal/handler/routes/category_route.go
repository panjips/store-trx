package routes

import (
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/usecase"

	"github.com/gorilla/mux"
)

func CategoryRoute(r *mux.Router, categoryHandler *usecase.CategoryHandler) {
	category := r.PathPrefix("/category").Subrouter()

	category.Use(middleware.AuthenticationMiddleware)
	category.HandleFunc("", categoryHandler.GetAll).Methods("GET")
	category.HandleFunc("/{id}", categoryHandler.GetByID).Methods("GET")

	admin := category.NewRoute().Subrouter()
    admin.Use(middleware.AuthenticationAdmin)
	admin.HandleFunc("", categoryHandler.Create).Methods("POST")
	admin.HandleFunc("/{id}", categoryHandler.Update).Methods("PUT")
	admin.HandleFunc("/{id}", categoryHandler.Delete).Methods("DELETE")
}