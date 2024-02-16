package handlers

import (
	"residential-registration/backend/internal/services"
	"residential-registration/backend/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Logger   logging.Logger
	Services services.Services
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.GET("/logout", h.userIdentity, h.logout)
	}
	osbb := router.Group("/osbb")
	{
		osbb.POST("/", h.registerOSBB)
		osbb.POST("/:osbbID/inhabitants", h.registerInhabitant)
		osbb.POST("/:osbbID/announcements", h.userIdentity, h.addAnnouncement)
		osbb.POST("/:osbbID/polls", h.userIdentity, h.addPoll)
		osbb.POST("/:osbbID/polls-test", h.userIdentity, h.addPollTest)
		osbb.POST("/:osbbID/polls/:pollID/answers", h.userIdentity, h.addPollAnswer)
		osbb.POST("/:osbbID/polls/:pollID/answers-test", h.userIdentity, h.addPollAnswerTest)
	}
	return router
}
