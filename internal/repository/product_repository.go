package repository

import (
	"store-trx-go/internal/entity"
	"store-trx-go/internal/handler/dto"
	"store-trx-go/pkg/database"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *entity.Product) error {
	err := r.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Update(ID uint, product *dto.PostProductRequest) error {
	err := r.db.Model(&entity.Product{}).Where("id = ?", ID).Updates(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) GetAll(params *dto.ParamsGetAllProduct) ([]*dto.ProductDTO, error) {
	var products []entity.Product
	paginate := database.NewPaginate(params.Limit, params.Page).PaginatedResult

	query := r.db.Model(products).Preload("Store").Preload("Category").Preload("Photos")
	if params.NamaProduk != "" {
		query = query.Where("name = ?" , params.NamaProduk)
	}

	if params.CategoryID > 0 {
		query = query.Where("category_id = ?" , params.CategoryID)
	}
	
	if params.TokoID > 0 {
		query = query.Where("store_id = ?" , params.TokoID)
	}

	if params.MinHarga != 0 {
		query = query.Where("customer_price >= ?" , params.MinHarga)
	}

	if params.MaxHarga != 0 {
		query = query.Where("customer_price <= ?" , params.MaxHarga)
	}

	err := query.Scopes(paginate).Find(&products).Error
	if err != nil {
		return nil, err
	}

	var productsDTO []*dto.ProductDTO
	for _ ,product := range products {
		var photoDTO []dto.PhotoDTO
		for _, photo := range product.Photos {
			photoDTO = append(photoDTO, dto.PhotoDTO{
				ID:  photo.ID,
				URL: photo.URL,
			})
		}

		productsDTO = append(productsDTO, &dto.ProductDTO{
			ID: product.ID,
			Name: product.Name,
			Slug: product.Slug,
			ResellerPrice: product.ResellerPrice,
			CustomerPrice: product.CustomerPrice,
			Stock: product.Stock,
			Description: product.Description,
			Store: dto.StoreDTO{
				ID: product.Store.ID,
				Name: *product.Store.Name,
				ImageURL: *product.Store.ImageURL,
			},
			Category: dto.CategoryDTO{
				ID: product.Category.ID,
				Name: product.Category.Name,
			},
			Photos: photoDTO,
		})
	}

	return productsDTO, nil
}

func (r *ProductRepository) FindByID(ID uint) (*dto.ProductDTO, error) {
	var product entity.Product
	err := r.db.Preload("Store").Preload("Category").Preload("Photos").Where("id = ?", ID).First(&product).Error
	if err != nil {
		return nil, err
	}

	var photoDTO []dto.PhotoDTO
	for _, photo := range product.Photos {
		photoDTO = append(photoDTO, dto.PhotoDTO{
			ID:  photo.ID,
			URL: photo.URL,
		})
	}

	productDTO := &dto.ProductDTO{
		ID: product.ID,
		Name: product.Name,
		Slug: product.Slug,
		ResellerPrice: product.ResellerPrice,
		CustomerPrice: product.CustomerPrice,
		Stock: product.Stock,
		Description: product.Description,
		Store: dto.StoreDTO{
			ID: product.Store.ID,
			Name: *product.Store.Name,
			ImageURL: *product.Store.ImageURL,
		},
		Category: dto.CategoryDTO{
			ID: product.Category.ID,
			Name: product.Category.Name,
		},
		Photos: photoDTO,
	}

	return productDTO, nil
}