package dto

type CategoryDTO struct {
	ID			uint		`json:"id"`
	Name		string 		`json:"nama_category" validate:"required"`
}