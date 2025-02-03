package dto

type UpdateProfileRequest struct {
	Name		*string 	`json:"nama,omitempty" validate:"omitempty"`		
	Password	*string		`json:"kata_sandi,omitempty" validate:"omitempty"`
	PhoneNumber	*string		`json:"no_telp,omitempty" validate:"omitempty"`
	BirthDate	*string		`json:"tanggal_lahir,omitempty" validate:"omitempty"`
	Work		*string		`json:"pekerjaan,omitempty" validate:"omitempty"`
	Email		*string		`json:"email,omitempty" validate:"omitempty"`
	ProvinceID	*string		`json:"id_provinsi,omitempty" validate:"omitempty"`
	CityID		*string		`json:"id_kota,omitempty" validate:"omitempty"`
}


type PostAddressRequest struct {
	AddressTitle	string 	`json:"judul_alamat" validate:"required"`	
	RecipientName	string 	`json:"nama_penerima" validate:"required"`
	PhoneNumber		string 	`json:"no_telp" validate:"required"`
	DetailAddress	string 	`json:"detail_alamat" validate:"required"`
}

type UpdateAddressRequest struct {
	RecipientName	*string 	`json:"nama_penerima,omitempty" validate:"omitempty"`
	PhoneNumber		*string 	`json:"no_telp,omitempty" validate:"omitempty"`
	DetailAddress	*string 	`json:"detail_alamat,omitempty" validate:"omitempty"`
}