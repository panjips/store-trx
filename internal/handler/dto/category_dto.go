package dto

type CategoryDTO struct {
	ID			string		`json:"id"`
	Name		string 		`json:"nama_category" validate:"required"`
}