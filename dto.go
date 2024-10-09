package main

// DTO Structs
type StudentRequest struct {
	Nama            string  `json:"nama_lengkap"`
	SuratElektronik string  `json:"surat_elektronik"`
	NoHP            string  `json:"no_hp"`
	Alamat          string  `json:"alamat"`
	IPK             float64 `json:"ipk"`
	IsGraduate      bool    `json:"is_graduate"`
}

type StudentRequestByID struct {
	ID string `param:"id"`
}