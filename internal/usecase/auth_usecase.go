package usecase

import (
	"encoding/json"
	"net/http"
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"
	"store-trx-go/internal/handler/responses"
	"store-trx-go/internal/repository"
	"store-trx-go/pkg/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo *repository.UserRepository
	storeRepo *repository.StoreRepository
}

func NewAuthUsecase(userRepo *repository.UserRepository, storeRepo *repository.StoreRepository) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
		storeRepo: storeRepo,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.ReqisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid request body", nil)
		return
	}

	validationErr := utils.ValidateRequest(req)
	if validationErr != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "validation error", validationErr)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to hash password", nil)
		return
	}

	birthDate, err := time.Parse("02/01/2006", req.BirthDate)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid birth date format", nil)
		return
	}

	user := &entity.User{
		Name: req.Name,		
		Password: string(hashedPassword),
		PhoneNumber: req.PhoneNumber,
		BirthDate: &birthDate,
		Work: req.Work,
		Email: req.Email,
		ProvinceID: req.ProvinceID,		
		CityID: req.CityID,
	}

	err = h.userRepo.Create(user)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to create user", nil)
		return
	}

	store := &entity.Store{
		UserID: user.ID,
	}

	err = h.storeRepo.Create(store)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to create store", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "user registered successfully", nil)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid request body", req)
		return
	}

	validationErr := utils.ValidateRequest(req)
	if validationErr != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "validation error", validationErr)
		return
	}

	user, err := h.userRepo.Login(req.PhoneNumber)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "invalid credential", nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusUnauthorized, "invalid password", nil)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to generate token", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader( http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",	
		"message": "login success",
		"token": token,
	})
}