package routes

import (
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/usecase"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router, userHandler *usecase.UserHandler) {
	user := r.PathPrefix("/user").Subrouter()
	user.Use(middleware.AuthenticationMiddleware)
	user.HandleFunc("", userHandler.Get).Methods("GET")
	user.HandleFunc("", userHandler.Update).Methods("PUT")

	alamat := user.PathPrefix("/alamat").Subrouter()
	alamat.HandleFunc("", userHandler.GetAddress).Methods("GET")
	alamat.HandleFunc("/{id}", userHandler.GetAddressByID).Methods("GET")
	alamat.HandleFunc("", userHandler.CreateAddress).Methods("POST")
	alamat.HandleFunc("/{id}", userHandler.UpdateAddress).Methods("PUT")
	alamat.HandleFunc("/{id}", userHandler.DeleteAddress).Methods("DELETE")

}