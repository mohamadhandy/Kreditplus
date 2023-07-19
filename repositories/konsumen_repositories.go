package repositories

import (
	"kredit-plus/helper"
	"kredit-plus/models"
	"net/http"

	"gorm.io/gorm"
)

type KonsumenRepositoryInterface interface {
	CreateKonsumen(konsumen models.KonsumenRequest) chan helper.Response
}

type konsumenRepository struct {
	db *gorm.DB
}

func InitKonsumenRepository(db *gorm.DB) KonsumenRepositoryInterface {
	return &konsumenRepository{
		db,
	}
}

func (r *konsumenRepository) CreateKonsumen(konsumen models.KonsumenRequest) chan helper.Response {
	modelKonsumen := models.Konsumen{
		NIK:          konsumen.NIK,
		FullName:     konsumen.FullName,
		FotoKTP:      konsumen.FotoKTP,
		FotoSelfie:   konsumen.FotoSelfie,
		LegalName:    konsumen.LegalName,
		Gaji:         konsumen.Gaji,
		TempatLahir:  konsumen.TempatLahir,
		TanggalLahir: konsumen.TanggalLahir,
		Role:         konsumen.Role,
		Email:        konsumen.Email,
		Password:     konsumen.Password,
	}
	result := make(chan helper.Response)
	go func() {
		if err := r.db.Create(&modelKonsumen).Error; err != nil {
			result <- helper.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
			return
		}
		result <- helper.Response{
			StatusCode: http.StatusCreated,
			Message:    "Register User Success",
			Data:       nil,
		}
	}()
	return result
}
