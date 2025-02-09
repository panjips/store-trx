package routes

import (
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/usecase"

	"github.com/gorilla/mux"
)

func TransactionRoute(r *mux.Router, trxHandler *usecase.TrxHandler){
	trx := r.PathPrefix("/trx").Subrouter()

	trx.Use(middleware.AuthenticationMiddleware)
	trx.HandleFunc("", trxHandler.Create).Methods("POST")
	trx.HandleFunc("", trxHandler.GetAll).Methods("GET")
}