package routes

import (
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/usecase"

	"github.com/gorilla/mux"
)

func ProductRoute(r *mux.Router, productHandler *usecase.ProductHandler) {
	product := r.PathPrefix("/product").Subrouter()

	product.Use(middleware.AuthenticationMiddleware)
	product.HandleFunc("", productHandler.GetAll).Methods("GET")
	product.HandleFunc("", productHandler.Create).Methods("POST")
	product.HandleFunc("/{id}", productHandler.GetByID).Methods("GET")
}