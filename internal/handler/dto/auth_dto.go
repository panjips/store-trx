package dto

type ReqisterRequest struct {
	Name		string 		`json:"nama" validate:"required"`		
	Password	string		`json:"kata_sandi" validate:"required"`
	PhoneNumber	string		`json:"no_telp" validate:"required"`
	BirthDate	string		`json:"tanggal_lahir" validate:"required"`
	Work		string		`json:"pekerjaan" validate:"required"`
	Email		string		`json:"email" validate:"required"`
	ProvinceID	string		`json:"id_provinsi" validate:"required"`
	CityID		string		`json:"id_kota" validate:"required"`
}

type LoginRequest struct {
	PhoneNumber	string		`json:"no_telp" validate:"required"`
	Password	string		`json:"kata_sandi" validate:"required"`
}