package repositories

import (
	"kredit-plus/helper"
	"kredit-plus/models"
	"net/http"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	GetProducts() chan helper.Response
}

type productRepository struct {
	db *gorm.DB
}

func InitProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &productRepository{db}
}

func (r *productRepository) GetProducts() chan helper.Response {
	result := make(chan helper.Response)
	go func() {
		products := []models.Product{}
		if err := r.db.Find(&products).Error; err != nil {
			result <- helper.Response{
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
				Message:    err.Error(),
			}
			return
		}
		productsResponse := []models.ProductResponse{}
		for _, v := range products {
			productsResponse = append(productsResponse, models.ProductResponse{
				NamaProduk: v.NamaProduk,
			})
		}
		result <- helper.Response{
			StatusCode: http.StatusOK,
			Data:       productsResponse,
			Message:    "Get Products success",
		}
	}()
	return result
}
