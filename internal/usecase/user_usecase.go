package usecase

import (
	"encoding/json"
	"net/http"
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"
	"store-trx-go/internal/handler/responses"
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/repository"
	"store-trx-go/pkg/utils"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct{
	userRepo	*repository.UserRepository
	addressRepo *repository.AddressRepository
}

func NewUserUsecase(userRepo *repository.UserRepository, addressRepo *repository.AddressRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
		addressRepo: addressRepo,
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(uint)

	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "user mot found", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "retrive user profile", user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateProfileRequest
	userID := r.Context().Value(middleware.UserIDKey).(uint)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid request body", nil)
		return 
	}

	validationErr := utils.ValidateRequest(req)
	if validationErr != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "validation error", validationErr)
		return
	}

	var hashedPassword []byte
	var birthDate time.Time
	var err error

	if req.Password != nil {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to hash password", nil)
			return
		}
	}

	if req.BirthDate != nil {
		birthDate, err = time.Parse("02/01/2006", *req.BirthDate)
		if err != nil {
			responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid birth date format", nil)
			return
		}
	}

	updateUser := &entity.User{
		Name: *req.Name,		
		Password: string(hashedPassword),
		PhoneNumber: *req.PhoneNumber,
		BirthDate: &birthDate,
		Work: *req.Work,
		Email: *req.Email,
		ProvinceID: *req.ProvinceID,		
		CityID: *req.CityID,
	}

	if err := h.userRepo.UpdateProfile(userID, updateUser); err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed update profile", nil)
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "success update profile", nil)
}

func (h *UserHandler) GetAddress(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(uint)

	addresses, err := h.addressRepo.FindByUserID(userID)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "address not found", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "retrive user address", addresses)
}

func (h *UserHandler) GetAddressByID(w http.ResponseWriter, r *http.Request) {
	addressIDStr := mux.Vars(r)["id"]
	addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid address ID", nil)
		return
	}

	address, err := h.addressRepo.FindByID(uint(addressID))
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "address not found", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "retrive user address", address)

}

func (h *UserHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	var req dto.PostAddressRequest
	userID := r.Context().Value(middleware.UserIDKey).(uint)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid request body", nil)
		return
	}

	validationErr := utils.ValidateRequest(req)
	if validationErr != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "validation error", validationErr)
		return
	}

	address := &entity.Address{
		UserID: userID,
		AddressTitle: req.AddressTitle,
		RecipientName: req.RecipientName,
		PhoneNumber: req.PhoneNumber,
		DetailAddress: req.DetailAddress,
	}

	if err := h.addressRepo.Create(address); err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to create address", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "address created successfully", nil)
}

func (h *UserHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateAddressRequest
	addressIDStr := mux.Vars(r)["id"]
	addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid address ID", nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid request body", nil)
		return
	}

	validationErr := utils.ValidateRequest(req)
	if validationErr != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "validation error", validationErr)
		return
	}

	updateAddress := &entity.Address{
		RecipientName: *req.RecipientName,
		PhoneNumber: *req.PhoneNumber,
		DetailAddress: *req.DetailAddress,
	}

	err = h.addressRepo.Update(uint(addressID), updateAddress)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "address not found", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "address updated successfully", nil)
}

func (h *UserHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	addressIDStr := mux.Vars(r)["id"]
	addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid address ID", nil)
		return
	}

	err = h.addressRepo.Delete(uint(addressID))
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to delete address", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "address deleted successfully", nil)
}