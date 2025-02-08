package usecase

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"
	"store-trx-go/internal/handler/responses"
	"store-trx-go/internal/middleware"
	"store-trx-go/internal/repository"
	"store-trx-go/pkg/r2"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

type ProductHandler struct {
	productRepo *repository.ProductRepository
	photoRepo	*repository.PhotoRepository
	storeRepo 	*repository.StoreRepository
}

func NewProductUsecase(productRepo *repository.ProductRepository, photoRepo	*repository.PhotoRepository, storeRepo *repository.StoreRepository) *ProductHandler {
	return &ProductHandler{
		productRepo: productRepo,
		photoRepo: photoRepo,
		storeRepo: storeRepo,
	}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product dto.PostProductRequest
	product.Name = r.FormValue("nama_produk")
	categoryID, err := strconv.ParseUint(r.FormValue("category_id"), 10, 64)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid category_id", nil)
		return
	}
	product.CategoryID = uint(categoryID)
	product.ResellerPrice = r.FormValue("harga_reseller")
	product.CustomerPrice = r.FormValue("harga_konsumen")
	stock, err := strconv.ParseUint(r.FormValue("stok"), 10, 64)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid category_id", nil)
		return
	}
	product.Stock = uint(stock)

	product.Description = r.FormValue("deskripsi")
	fileHeaders, ok := r.MultipartForm.File["photos"]
	if !ok || len(fileHeaders) == 0 {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid file photos", nil)
		return
	}

	var wg sync.WaitGroup
	photoURLs := make([]string, len(fileHeaders))
	var mu sync.Mutex
	errCh := make(chan error, len(fileHeaders))

	for i, fh := range fileHeaders {
		wg.Add(1)
		go func(i int, fh *multipart.FileHeader) {
			defer wg.Done()
			file, err := fh.Open()
			if err != nil {
				errCh <- fmt.Errorf("failed to open file: %w", err)
				return
			}
			defer file.Close()
			result, err := r2.UploadFile(fmt.Sprintf("%s-%d.png", product.Name, i), file)
			if err != nil {
				errCh <- fmt.Errorf("failed to upload file: %w", err)
				return
			}
			mu.Lock()
			photoURLs[i] = result
			mu.Unlock()
		}(i, fh)
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			responses.HTTPResponse(w, "error", http.StatusInternalServerError, err.Error(), nil)
			return
		}
	}

	product.Photos = photoURLs
	userID := r.Context().Value(middleware.UserIDKey).(uint)
	store, err := h.storeRepo.FindByUserID(userID)
	if err != nil {
		responses.HTTPResponse(w, "error",  http.StatusInternalServerError, "user doesnt have store", nil)
		return
	}


	newProduct := &entity.Product{
		Name: product.Name,
		Slug: slug.Make(product.Name),
		ResellerPrice: product.ResellerPrice,
		CustomerPrice: product.CustomerPrice,
		Stock: product.Stock,
		Description: product.Description,
		CategoryID: product.CategoryID,
		StoreID: store.ID,
	}

	err = h.productRepo.Create(newProduct)
	if err != nil {
		responses.HTTPResponse(w, "error",  http.StatusInternalServerError, "failed to store product", nil)
		return
	}

	errCh = make(chan error, len(product.Photos))
	for _, photoURL := range product.Photos {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			newPhoto := &entity.Photo{
				ProductID: newProduct.ID,
				URL: url,
			}
			if err := h.photoRepo.Create(newPhoto); err != nil {
				errCh <- err
			}
		}(photoURL)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			responses.HTTPResponse(w, "error", http.StatusInternalServerError, "failed to store photo product", nil)
			return
		}
	}

	responses.HTTPResponse(w, "success",  http.StatusCreated, "success create new product", nil)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request)  {
	params := &dto.ParamsGetAllProduct{
		Limit: 10,
		Page: 1,
		NamaProduk: "",
		CategoryID: 0,
		TokoID: 0,
		MinHarga: 0,
		MaxHarga: 0,
	}

	params.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	params.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	params.NamaProduk = r.URL.Query().Get("nama_produk")
	params.CategoryID, _ = strconv.Atoi(r.URL.Query().Get("category_id"))
	params.TokoID, _ = strconv.Atoi(r.URL.Query().Get("toko_id"))
	minHarga, _ := strconv.ParseUint(r.URL.Query().Get("min_harga"), 10, 64)
	params.MinHarga = uint(minHarga)
	maxHarga, _ := strconv.ParseUint(r.URL.Query().Get("max_harga"), 10, 64)
	params.MaxHarga = uint(maxHarga)

	products, err := h.productRepo.GetAll(params)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "failed retrive products", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "product retrieved", products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request)  {
	productIDStr := mux.Vars(r)["id"]
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid product ID", nil)
		return
	}

	product, err := h.productRepo.FindByID(uint(productID))
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "product not found", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "product retrieved", product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	userIDStr := mux.Vars(r)["id"]
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid product ID", nil)
		return
	}

	var product dto.PostProductRequest
	product.Name = r.FormValue("nama_produk")
	categoryID, err := strconv.ParseUint(r.FormValue("category_id"), 10, 64)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid category_id", nil)
		return
	}
	product.CategoryID = uint(categoryID)
	product.ResellerPrice = r.FormValue("harga_reseller")
	product.CustomerPrice = r.FormValue("harga_konsumen")
	stock, err := strconv.ParseUint(r.FormValue("stok"), 10, 64)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid category_id", nil)
		return
	}
	product.Stock = uint(stock)

	product.Description = r.FormValue("deskripsi")
	fileHeaders, ok := r.MultipartForm.File["photos"]
	if !ok || len(fileHeaders) == 0 {
		responses.HTTPResponse(w, "error", http.StatusBadRequest, "invalid file photos", nil)
		return
	}

	err = h.productRepo.Update(uint(userID), &product)
	if err != nil {
		responses.HTTPResponse(w, "error", http.StatusNotFound, "product not found", nil)
		return
	}

	responses.HTTPResponse(w, "success", http.StatusOK, "product retrieved", product)
}