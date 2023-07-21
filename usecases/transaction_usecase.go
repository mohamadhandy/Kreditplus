package usecases

import (
	"kredit-plus/helper"
	"kredit-plus/models"
	"kredit-plus/repositories"
)

type TransaksiUsecaseInterface interface {
	CreateTransaction(tokenString string, transaksiRequest models.TransaksiRequest) helper.Response
	GetTransactions(tokenString string) helper.Response
}

type transaksiUsecase struct {
	r repositories.TransaksiRepositoryInterface
}

func InitTransaksiUsecase(r repositories.TransaksiRepositoryInterface) TransaksiUsecaseInterface {
	return &transaksiUsecase{r: r}
}

func (u *transaksiUsecase) CreateTransaction(tokenString string, transaksiRequest models.TransaksiRequest) helper.Response {
	return <-u.r.CreateTransaction(tokenString, transaksiRequest)
}

func (u *transaksiUsecase) GetTransactions(tokenString string) helper.Response {
	return <-u.r.GetTransactions(tokenString)
}
