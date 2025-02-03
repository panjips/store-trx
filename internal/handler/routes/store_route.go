package routes

import (
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/usecase"

	"github.com/gorilla/mux"
)

func StoreRoute(r *mux.Router, storeHandler *usecase.StoreHandler){
	store := r.PathPrefix("/toko").Subrouter()
	
	store.Use(middleware.AuthenticationMiddleware)
	store.HandleFunc("/my", storeHandler.GetByUserID).Methods("GET")
	store.HandleFunc("/{id}", storeHandler.GetByID).Methods("GET")
	store.HandleFunc("/{id}", storeHandler.Update).Methods("PUT")
}