package models

type Konsumen struct {
	ID           int     `gorm:"column:id_konsumen" json:"id_konsumen"`
	NIK          string  `gorm:"column:nik" json:"nik"`
	FullName     string  `gorm:"column:full_name" json:"full_name"`
	LegalName    string  `gorm:"column:legal_name" json:"legal_name"`
	Gaji         float64 `gorm:"column:gaji" json:"gaji"`
	TempatLahir  string  `gorm:"column:tempat_lahir" json:"tempat_lahir"`
	TanggalLahir string  `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	FotoKTP      string  `gorm:"column:foto_ktp" json:"foto_ktp"`
	FotoSelfie   string  `gorm:"column:foto_selfie" json:"foto_selfie"`
	Role         string  `gorm:"column:role" json:"role"`
	Email        string  `gorm:"column:email" json:"email"`
	Password     string  `gorm:"column:password" json:"password"`
}

type KonsumenRequest struct {
	NIK          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	Gaji         float64 `json:"gaji"`
	TempatLahir  string  `json:"tempat_lahir"`
	TanggalLahir string  `json:"tanggal_lahir"`
	FotoKTP      string  `json:"foto_ktp"`
	FotoSelfie   string  `json:"foto_selfie"`
	Role         string  `json:"role"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
}

func (Konsumen) TableName() string {
	return "Konsumen"
}
