package usecase

import (
	"encoding/json"
	"net/http"
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"
	"store-trx-go/internal/handler/responses"
	"store-trx-go/internal/repository"
	"strconv"

	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo *repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{
		categoryRepo: categoryRepo,
	}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CategoryDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid request body", nil)
		return
	}

	category := &entity.Category{
		Name: req.Name,
	}

	err = h.categoryRepo.Create(category)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to create new category", nil)
		return
	}
	
	responses.HTTPResponse(w, "success", http.StatusCreated, "success to create new category", nil)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.CategoryDTO
	categoryIDStr := mux.Vars(r)["id"]
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid address ID", nil)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid request body", nil)
		return
	}

	category := &entity.Category{
		Name: req.Name,
	}

	err = h.categoryRepo.Update(uint(categoryID),category)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to update category", nil)
		return
	}
	
	responses.HTTPResponse(w, "success", http.StatusCreated, "success to update category", nil)
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories , err := h.categoryRepo.GetAll()
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "failed to retrive categories", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "success retrive categories", categories)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	categoryIDStr := mux.Vars(r)["id"]
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid address ID", nil)
		return
	}

	category , err := h.categoryRepo.GetByID(uint(categoryID))
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "failed to retrive category", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "success retrive category", category)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	categoryIDStr := mux.Vars(r)["id"]
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid address ID", nil)
		return
	}
	err = h.categoryRepo.Delete(uint(categoryID))
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "failed to delete category", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "success delete category", nil)
}