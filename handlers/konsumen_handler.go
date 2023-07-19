package handlers

import (
	"kredit-plus/models"
	"kredit-plus/usecases"

	"github.com/gin-gonic/gin"
)

type KonsumenHandlerInterface interface {
	CreateUser(c *gin.Context)
}

type konsumenHandler struct {
	KonsumenUseCase usecases.KonsumenUsecaseInterface
}

func InitKonsumenHandler(u usecases.KonsumenUsecaseInterface) KonsumenHandlerInterface {
	return &konsumenHandler{
		KonsumenUseCase: u,
	}
}

func (h *konsumenHandler) CreateUser(c *gin.Context) {
	konsumen := models.KonsumenRequest{}
	c.BindJSON(&konsumen)
	konsumenResult := h.KonsumenUseCase.CreateUser(konsumen)
	c.JSON(konsumenResult.StatusCode, konsumenResult)
}
