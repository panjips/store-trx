package dto


type PhotoDTO struct {
	ID			uint		`json:"id"`
	URL			string		`json:"url"`
}

type ProductDTO struct {
	ID				uint		`json:"id"`
	Name			string		`json:"nama_produk"`
	Slug			string		`json:"slug"`
	ResellerPrice	string		`json:"harga_reseller"`
	CustomerPrice	string		`json:"harga_konsumen"`
	Stock			uint		`json:"stok"`
	Description		string		`json:"deskripsi"`
	Store			StoreDTO	`json:"toko"`	
	Category		CategoryDTO	`json:"category"`
	Photos			[]PhotoDTO	`json:"photos"`		
}

type PostProductRequest struct {
	Name			string		`json:"nama_produk" validate:"required"`
	CategoryID		uint		`json:"category_id" validate:"required"`
	ResellerPrice	string		`json:"harga_reseller" validate:"required"`
	CustomerPrice	string		`json:"harga_konsumen" validate:"required"`
	Stock			uint		`json:"stok" validate:"required"`
	Description		string		`json:"deskripsi" validate:"required"`
	Photos			[]string	`json:"photos" validate:"required"`
}


type ParamsGetAllProduct struct {
	Limit			int
	Page			int
	NamaProduk 		string
	CategoryID		int
	TokoID			int
	MinHarga		uint
	MaxHarga		uint
}