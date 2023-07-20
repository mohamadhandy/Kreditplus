package repositories

import (
	"kredit-plus/helper"
	"kredit-plus/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	BeginSession(authRequest models.AuthRequest) chan helper.Response
}

type authRepository struct {
	db *gorm.DB
}

func InitAuthRepository(db *gorm.DB) AuthRepositoryInterface {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) BeginSession(authRequest models.AuthRequest) chan helper.Response {
	result := make(chan helper.Response)
	go func() {
		konsumen := models.Konsumen{}
		email, password := authRequest.Email, authRequest.Password
		err := r.db.Where(&models.Konsumen{Email: email}).First(&konsumen).Error
		if err == gorm.ErrRecordNotFound {
			// isi data nanti
			result <- helper.Response{
				Data:       nil,
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
			return
		} else if err != nil {
			// isi data nanti
			result <- helper.Response{
				Data:       nil,
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(konsumen.Password), []byte(password)); err != nil {
			result <- helper.Response{
				Data:       nil,
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
			return
		}
		konsumenResponse := models.KonsumenResponse{
			ID:           konsumen.ID,
			NIK:          konsumen.NIK,
			FullName:     konsumen.FullName,
			LegalName:    konsumen.LegalName,
			Gaji:         konsumen.Gaji,
			TempatLahir:  konsumen.TempatLahir,
			TanggalLahir: konsumen.TanggalLahir,
			FotoKTP:      konsumen.FotoKTP,
			FotoSelfie:   konsumen.FotoSelfie,
			Role:         konsumen.Role,
			Email:        konsumen.Email,
		}
		result <- helper.Response{
			Data:       konsumenResponse,
			Message:    "Login success",
			StatusCode: http.StatusOK,
		}
	}()
	return result
}
