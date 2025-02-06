package routes

import (
	"store-trx-go/internal/repository"
	"store-trx-go/internal/usecase"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(r *mux.Router, db *gorm.DB) {

	storeRepo := repository.NewStoreRepository(db)
	userRepo := repository.NewUserRepository(db)
	addressRepo := repository.NewAddressRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, storeRepo)
	userUsecase := usecase.NewUserUsecase(userRepo, addressRepo)
	storeUsecase := usecase.NewStoreUsecase(storeRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	AuthRoute(r, authUsecase)
	StoreRoute(r, storeUsecase)
	UserRoute(r, userUsecase)
	CategoryRoute(r, categoryUsecase)
}