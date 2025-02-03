package routes

import (
	"store-trx-go/internal/usecase"

	"github.com/gorilla/mux"
)

func AuthRoute(r *mux.Router, authHandler *usecase.AuthHandler) {
	auth := r.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")
}