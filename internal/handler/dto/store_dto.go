package dto

type StoreDTO struct {
	ID			uint		`json:"id"`
	Name		string 		`json:"name"`
	ImageURL	string		`json:"photo"`
}

type UpdateStoreRequest struct {
	Name		*string 	`json:"nama_toko,omitempty" validate:"omitempty"`
	ImageURL	*string		`json:"photo,omitempty" validate:"omitempty"`
}