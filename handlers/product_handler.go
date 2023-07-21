package handlers

import (
	"kredit-plus/usecases"

	"github.com/gin-gonic/gin"
)

type ProductHandlerInterface interface {
	GetProducts(c *gin.Context)
}

type productHandler struct {
	u usecases.ProductUsecaseInterface
}

func InitProductHandler(u usecases.ProductUsecaseInterface) ProductHandlerInterface {
	return &productHandler{u}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	products := h.u.GetProducts()
	c.JSON(products.StatusCode, products)
}
