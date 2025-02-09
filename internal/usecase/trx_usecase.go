package usecase

import (
	"encoding/json"
	"net/http"
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"
	"store-trx-go/internal/handler/responses"
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/repository"
	"strconv"
)

type TrxHandler struct {
	trxRepo		*repository.TrxRepository
	productRepo	*repository.ProductRepository
}

func NewTrxUsecase(trxRepo *repository.TrxRepository, productRepo *repository.ProductRepository) *TrxHandler {
	return &TrxHandler{
		trxRepo: trxRepo,
		productRepo: productRepo,
	}
}

func (h *TrxHandler) Create(w http.ResponseWriter, r *http.Request) {
	var transaction dto.PostTrxRequest
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid body request", nil)
		return
	}

	userIDStr := r.Context().Value(middleware.UserIDKey)
	userID, ok := userIDStr.(uint)
	if !ok {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "invalid user ID", nil)
		return
	}

	var trx = entity.Transaction{
		PaymentMethod: transaction.PaymentMethod,
		UserID: userID,
		AddressID: transaction.AddressID,
	}

	err = h.trxRepo.Create(&trx)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed create new transaction", nil)
	}

	var totalTrx uint = 0

	for _, detailTrx := range transaction.DetailTrx  {
		product, err := h.productRepo.FindByID(uint(detailTrx.ProductID))
		if err != nil {
			responses.HTTPResponse(w, "error", http.StatusInternalServerError, "invalid product", err)
			return
		}
		price, err := strconv.ParseUint(product.CustomerPrice, 10, 64)
		if err != nil {
			responses.HTTPResponse(w, "error", http.StatusInternalServerError, "invalid product price", nil)
			return
		}
		var detailTrx = entity.DetailTransaction{
			TransactionID: trx.ID,
			Quantity: detailTrx.Quantity,
			TotalPrice: (detailTrx.Quantity * uint(price)),
			ProductID: product.ID,
		}
		err = h.trxRepo.CreateDetailTrx(&detailTrx)
		if err != nil {
			responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed create new detail trx", nil)
			return
		}
		totalTrx += (detailTrx.Quantity * uint(price))
	}

	trx.TotalPrice = totalTrx
	err = h.trxRepo.Update(&trx)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to update trx", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusCreated, "success create new transaction", nil)

}

func (h *TrxHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tranaction, err := h.trxRepo.GetAll()
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "failed retrive tranaction", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "transaction retrieved", tranaction)
}