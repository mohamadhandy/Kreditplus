package models

type Product struct {
	ID         int    `json:"id_product" gorm:"column:id_produk"`
	NamaProduk string `json:"name" gorm:"column:nama_produk"`
}

type ProductResponse struct {
	NamaProduk string `json:"nama_produk" gorm:"column:nama_produk"`
}

func (Product) TableName() string {
	return "Produk"
}
