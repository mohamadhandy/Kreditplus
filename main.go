package main

import (
	"context"
	"kredit-plus/config"
	"kredit-plus/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RouteAPI(router, context.Background(), config.NewConnection())
	router.Run("localhost:9000")
}
