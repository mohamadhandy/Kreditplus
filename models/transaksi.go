package models

import "time"

type Transaksi struct {
	ID               int       `gorm:"column:id_transaksi" json:"id_transaksi"`
	IDKonsumen       int       `gorm:"column:id_konsumen" json:"id_konsumen"`
	NomorKontrak     string    `gorm:"column:nomor_kontrak" json:"nomor_kontrak"`
	TanggalTransaksi time.Time `gorm:"column:tanggal_transaksi" json:"tanggal_transaksi"`
	OTR              float64   `gorm:"column:otr" json:"otr"`
	AdminFee         float64   `gorm:"column:admin_fee" json:"admin_fee"`
	JumlahCicilan    int       `gorm:"column:jumlah_cicilan" json:"jumlah_cicilan"`
	JumlahBunga      float64   `gorm:"column:jumlah_bunga" json:"jumlah_bunga"`
	NamaAsset        string    `gorm:"column:nama_asset" json:"nama_asset"`
	JenisTransaksi   string    `gorm:"column:jenis_transaksi" json:"jenis_transaksi"`
}

type DetailTransaksi struct {
	IDDetail    int `gorm:"column:id_detail;primarykey" json:"id_detail"`
	IDTransaksi int `gorm:"column:id_transaksi" json:"id_transaksi"`
	IDProduk    int `gorm:"column:id_produk" json:"id_produk"`
	JumlahBeli  int `gorm:"column:jumlah_beli" json:"jumlah_beli"`
}

type TransaksiRequest struct {
	IDKonsumen      int               `json:"id_konsumen"`
	OTR             float64           `json:"otr"`
	JumlahCicilan   int               `json:"jumlah_cicilan"`
	NamaAsset       string            `json:"nama_asset"`
	JenisTransaksi  string            `json:"jenis_transaksi"`
	DetailTransaksi []DetailTransaksi `json:"detail_transaksi"`
}

func (Transaksi) TableName() string {
	return "Transaksi"
}

func (DetailTransaksi) TableName() string {
	return "DetailTransaksi"
}
