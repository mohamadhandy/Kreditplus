package repositories

import (
	"kredit-plus/helper"
	"kredit-plus/models"

	"gorm.io/gorm"
)

type TransaksiRepositoryInterface interface {
	CreateTransaction(tokenString string, tr models.TransaksiRequest) chan helper.Response
}

type transaksiRepository struct {
	db *gorm.DB
}

func InitTransaksiRepository(db *gorm.DB) TransaksiRepositoryInterface {
	return &transaksiRepository{db}
}

func (r *transaksiRepository) CreateTransaction(tokenString string, tr models.TransaksiRequest) chan helper.Response {
	result := make(chan helper.Response)
	go func() {

	}()
	return result
}
