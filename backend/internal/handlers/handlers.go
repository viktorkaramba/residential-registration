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

	osbb := router.Group("/osbb")
	{
		osbb.POST("/", h.registerOSBB)
		osbb.POST("/:id/inhabitant", h.registerInhabitant)
	}
	return router
}
