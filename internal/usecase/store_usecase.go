package usecase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/responses"
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/repository"
	"store-trx-go/pkg/r2"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gorilla/mux"
)

type StoreHandler struct {
	storeRepo *repository.StoreRepository
}

func NewStoreUsecase (storeRepo *repository.StoreRepository) *StoreHandler {
	return &StoreHandler{storeRepo: storeRepo}
}

func (h *StoreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid limit parameter", nil)
		return
	}

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid page parameter", nil)
		return
	}

	stores, err := h.storeRepo.GetAll(page, limit)
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
	var storeIDStr = mux.Vars(r)["id"]
	storeID, err := strconv.ParseUint(storeIDStr, 10, 32)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid store ID", nil)
		return
	}

	var name = r.FormValue("nama_toko")
	image, _, err := r.FormFile("photo")
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid photo", nil)
		return
	}

	var keyImage string = fmt.Sprintf("%s.png", name)
	uploadFile := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key: aws.String(keyImage),
		Body: image,
	}

	_, err = r2.GetClient().PutObject(context.Background(), uploadFile)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "1 validation error", nil)
		return
	}

	imageEndpoint := os.Getenv("IMAGE_URL_ENDPOINT")
	updateStore := &entity.Store{
		Name: &name,
		ImageURL: aws.String(fmt.Sprintf("%s%s", imageEndpoint, url.PathEscape(keyImage))),
	}

	err = h.storeRepo.Update(uint(storeID), updateStore)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to update store", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "store updated", nil)
}