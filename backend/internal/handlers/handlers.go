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
		osbb.POST("/:osbbID/inhabitants", h.registerInhabitant)
		osbb.POST("/:osbbID/announcements", h.userIdentity, h.addAnnouncement)
		osbb.POST("/:osbbID/polls", h.addPoll)
		osbb.POST("/:osbbID/polls-test", h.addPollTest)
		osbb.POST("/:osbbID/polls/:pollID/answers", h.addPollAnswer)
		osbb.POST("/:osbbID/polls/:pollID/answers-test", h.addPollAnswerTest)
	}
	return router
}
