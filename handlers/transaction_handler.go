package handlers

import (
	"kredit-plus/models"
	"kredit-plus/usecases"

	"github.com/gin-gonic/gin"
)

type TransactionHandlerInterface interface {
	CreateTransaction(c *gin.Context)
	GetTransactions(c *gin.Context)
}

type transactionHandler struct {
	uc usecases.TransaksiUsecaseInterface
}

func InitTransactionHandler(uc usecases.TransaksiUsecaseInterface) TransactionHandlerInterface {
	return &transactionHandler{uc}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tr := models.TransaksiRequest{}
	c.BindJSON(&tr)
	result := h.uc.CreateTransaction(tokenString, tr)
	c.JSON(result.StatusCode, result)
}

func (h *transactionHandler) GetTransactions(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	result := h.uc.GetTransactions(tokenString)
	c.JSON(result.StatusCode, result)
}
