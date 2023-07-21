package repositories

import (
	"kredit-plus/helper"
	"kredit-plus/middleware"
	"kredit-plus/models"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type TransaksiRepositoryInterface interface {
	CreateTransaction(tokenString string, tr models.TransaksiRequest) chan helper.Response
	GetTransactions(tokenString string) chan helper.Response
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
		tx := r.db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				result <- helper.Response{
					Data:       nil,
					Message:    "An unexpected error occurred",
					StatusCode: http.StatusInternalServerError,
				}
			}
		}()
		newTransaction := models.Transaksi{
			JenisTransaksi:   tr.JenisTransaksi,
			OTR:              tr.OTR,
			IDKonsumen:       tr.IDKonsumen,
			JumlahBunga:      helper.HitungJumlahBunga(tr.OTR, 6),
			AdminFee:         helper.HitungAdminFee(tr.OTR),
			TanggalTransaksi: time.Now(),
			JumlahCicilan:    tr.JumlahCicilan,
			NamaAsset:        tr.NamaAsset,
			NomorKontrak:     helper.GenerateNomorKontrak(tr.JenisTransaksi, tr.IDKonsumen, tr.JumlahCicilan),
		}

		if err := tx.Create(&newTransaction).Error; err != nil {
			tx.Rollback()
			result <- helper.Response{
				Data:       nil,
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
			return
		}

		// Membuat data detail transaksi
		for _, dt := range tr.DetailTransaksi {
			newDetailTransaksi := models.DetailTransaksi{
				IDProduk:   dt.IDProduk,
				JumlahBeli: dt.JumlahBeli,
			}
			if newTransaction.ID > 0 {
				newDetailTransaksi.IDTransaksi = newTransaction.ID
			}

			if err := tx.Debug().Create(&newDetailTransaksi).Error; err != nil {
				tx.Rollback()
				result <- helper.Response{
					Data:       nil,
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
				}
				return
			}
		}

		// Commit transaksi jika semuanya berhasil
		tx.Commit()

		result <- helper.Response{
			Data:       newTransaction,
			StatusCode: http.StatusCreated,
			Message:    "Success Create Transaction",
		}
	}()
	return result
}

func (r *transaksiRepository) GetTransactions(tokenString string) chan helper.Response {
	result := make(chan helper.Response)
	go func() {
		transactions := []models.Transaksi{}
		idKonsumen, errToken := middleware.ExtractRoleFromToken(tokenString)
		if errToken != nil {
			result <- helper.Response{
				Data:       nil,
				StatusCode: http.StatusInternalServerError,
				Message:    errToken.Error(),
			}
			return
		}
		if err := r.db.Where("id_konsumen = ?", idKonsumen).Find(&transactions).Error; err != nil {
			result <- helper.Response{
				Data:       nil,
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
			return
		}
		result <- helper.Response{
			Data:       transactions,
			Message:    "Success Get Transactions",
			StatusCode: http.StatusOK,
		}
	}()
	return result
}
