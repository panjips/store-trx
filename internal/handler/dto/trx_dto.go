package dto

type PostDetailTrxRequest struct {
	ProductID	uint	`json:"product_id" validate:"required"`
	Quantity	uint	`json:"kuantitas" validate:"required"`
}

type PostTrxRequest struct {
	PaymentMethod		string					`json:"method_bayar" validate:"required"`
	AddressID			uint					`json:"alamat_kirim" validate:"required"`
	DetailTrx			[]PostDetailTrxRequest	`json:"detail_trx" validate:"required"`
}


type DetailTrxDTO struct {
	ProductID	uint	`json:"product_id"`
	Quantity	uint	`json:"kuantitas"`
	TotalPrice	uint	`json:"total_harga"`
}

type TrxDTO struct {
	ID					uint			`json:"id"`
	PaymentMethod		string			`json:"method_bayar"`
	AddressID			uint			`json:"alamat_kirim"`
	UserID				uint			`json:"id_user"`
	TotalPrice			uint			`json:"total_harga"`
	DetailTrx			[]DetailTrxDTO	`json:"detail_trx"`
}