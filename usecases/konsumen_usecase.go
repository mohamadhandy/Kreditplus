package usecases

import (
	"kredit-plus/helper"
	"kredit-plus/models"
	"kredit-plus/repositories"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type KonsumenUsecaseInterface interface {
	CreateUser(konsumen models.KonsumenRequest) helper.Response
}

type konsumenUsecase struct {
	konsumenRepository repositories.KonsumenRepositoryInterface
}

func InitKonsumenUsecase(r repositories.KonsumenRepositoryInterface) KonsumenUsecaseInterface {
	return &konsumenUsecase{
		konsumenRepository: r,
	}
}

func (u *konsumenUsecase) CreateUser(konsumen models.KonsumenRequest) helper.Response {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(konsumen.Password), bcrypt.DefaultCost)
	if err != nil {
		return helper.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	konsumen.Password = string(passwordHash)
	return <-u.konsumenRepository.CreateKonsumen(konsumen)
}
