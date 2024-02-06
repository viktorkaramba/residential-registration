package handlers

import (
	"residential-registration/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services services.Services
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ping pong",
		})
	})
	return router
}
