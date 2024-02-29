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
		osbb.GET("/:osbbID/inhabitants", h.userIdentity, h.getAllInhabitants)
		osbb.GET("/:osbbID/inhabitants/profile", h.userIdentity, h.getInhabitantsProfile)
		osbb.PUT("/:osbbID/inhabitants", h.userIdentity, h.updateInhabitant)

		osbb.POST("/:osbbID/announcements", h.userIdentity, h.addAnnouncement)
		osbb.GET("/:osbbID/announcements", h.userIdentity, h.getAllAnnouncement)

		osbb.POST("/:osbbID/polls", h.userIdentity, h.addPoll)
		osbb.POST("/:osbbID/polls-test", h.userIdentity, h.addPollTest)
		osbb.GET("/:osbbID/polls", h.userIdentity, h.getAllPolls)

		osbb.POST("/:osbbID/polls/:pollID/answers", h.userIdentity, h.addPollAnswer)
		osbb.POST("/:osbbID/polls/:pollID/answers-test", h.userIdentity, h.addPollAnswerTest)
		osbb.GET("/:osbbID/polls/:pollID/answers", h.userIdentity, h.getAllPollsAnswers)

		osbb.POST("/:osbbID/payments", h.userIdentity, h.addPayment)
		osbb.POST("/:osbbID/payments/:paymentID/make-purchase", h.userIdentity, h.makePurchase)
	}
	return router
}
