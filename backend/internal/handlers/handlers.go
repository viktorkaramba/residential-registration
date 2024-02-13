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
	auth := router.Group("/auth")
	{
		auth.POST("/register-inhabitant/:id", h.registerInhabitant)
	}
	return router
}
