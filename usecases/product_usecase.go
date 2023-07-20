package usecases

import (
	"kredit-plus/helper"
	"kredit-plus/repositories"
)

type ProductUsecaseInterface interface {
	GetProducts() helper.Response
}

type productUsecase struct {
	r repositories.ProductRepositoryInterface
}

func InitProductUseCase(r repositories.ProductRepositoryInterface) ProductUsecaseInterface {
	return &productUsecase{r}
}

func (u *productUsecase) GetProducts() helper.Response {
	return <-u.r.GetProducts()
}
