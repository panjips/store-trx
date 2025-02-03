package dto

type UpdateStoreRequest struct {
	Name		*string 	`json:"nama_toko,omitempty" validate:"omitempty"`
	ImageURL	*string		`json:"photo,omitempty" validate:"omitempty"`
}