package handlers

import (
	"kredit-plus/helper"
	"kredit-plus/middleware"
	"kredit-plus/models"
	"kredit-plus/usecases"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationHandlerInterface interface {
	Login(c *gin.Context)
}

type authHandler struct {
	auth usecases.AuthUseCaseInterface
}

func InitAuthenticationHandler(auth usecases.AuthUseCaseInterface) AuthenticationHandlerInterface {
	return &authHandler{
		auth,
	}
}

func (h *authHandler) Login(c *gin.Context) {
	email := c.Request.FormValue("email")
	pass := c.Request.FormValue("password")
	reqModel := models.AuthRequest{
		Email:    email,
		Password: pass,
	}
	// kalau error ganti ke pointer
	result := h.auth.BeginSession(reqModel)
	if result.StatusCode != 200 {
		c.JSON(result.StatusCode, helper.Response{
			StatusCode: result.StatusCode,
			Message:    result.Message,
			Data:       nil,
		})
		return
	}
	konsumenResponse, _ := result.Data.(models.KonsumenResponse)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.MyCustomClaims{
		IDKonsumen: konsumenResponse.ID,
		Email:      konsumenResponse.Email,
		Name:       konsumenResponse.FullName,
		Role:       konsumenResponse.Role,
		Gaji:       konsumenResponse.Gaji,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusUnprocessableEntity, helper.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		StatusCode: http.StatusOK,
		Message:    "Login Success",
		Data:       tokenString,
	})
}
