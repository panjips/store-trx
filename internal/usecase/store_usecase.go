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

	"github.com/gorilla/mux"
)

type StoreHandler struct {
	storeRepo *repository.StoreRepository
}

func NewStoreUsecase (storeRepo *repository.StoreRepository) *StoreHandler {
	return &StoreHandler{storeRepo: storeRepo}
}

func (h *StoreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	stores, err := h.storeRepo.GetAll()
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to get stores", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "stores retrieved", stores)
}

func (h *StoreHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	storeIDStr := mux.Vars(r)["id"]
	storeID, err := strconv.ParseUint(storeIDStr, 10, 32)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid store ID", nil)
		return
	}

	store, err := h.storeRepo.FindByID(uint(storeID))
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "store not found", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "store retrieved", store)
}

func (h *StoreHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(uint)
	store, err := h.storeRepo.FindByUserID(userID)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "store not found", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "store retrieved", store)
}

func (h *StoreHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateStoreRequest
	var storeIDStr = mux.Vars(r)["id"]
	storeID, err := strconv.ParseUint(storeIDStr, 10, 32)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid store ID", nil)
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

	updateStore := &entity.Store{
		Name: req.Name,
		ImageURL: req.ImageURL,
	}

	err = h.storeRepo.Update(uint(storeID), updateStore)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to update store", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "store updated", nil)
}