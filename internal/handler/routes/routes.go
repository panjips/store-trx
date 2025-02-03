package routes

import (
	"store-trx-go/internal/repository"
	"store-trx-go/internal/usecase"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(r *mux.Router, db *gorm.DB) {

	tokoRepo := repository.NewStoreRepository(db)
	userRepo := repository.NewUserRepository(db)
	addressRepo := repository.NewAddressRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, tokoRepo)
	userUsecase := usecase.NewUserUsecase(userRepo, addressRepo)

	AuthRoute(r, authUsecase)
	UserRoute(r, userUsecase)
}