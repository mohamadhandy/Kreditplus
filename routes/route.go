package routes

import (
	"context"
	"kredit-plus/handlers"
	"kredit-plus/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteAPI(g *gin.Engine, parentCtx context.Context, db *gorm.DB) {
	repository := usecases.InitRepository(db)
	konsumen := handlers.InitVersionOneKonsumenHandler(repository)
	auth := handlers.InitVersionOneAuthHandler(repository)

	g.POST("/api/login", auth.Login)

	g.POST("/api/users", konsumen.CreateUser)
	g.POST("/api/upload-image", konsumen.UploadImageKonsumen)
}
