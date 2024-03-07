package handlers

import (
	"residential-registration/backend/internal/services"
	"residential-registration/backend/pkg/logging"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Logger   logging.Logger
	Services services.Services
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.GET("/logout", h.userIdentity, h.logout)
	}
	osbb := router.Group("/osbb")
	{
		osbb.POST("/", h.registerOSBB)
		osbb.GET("/", h.getAllOSBB)
		osbb.GET("/profile", h.userIdentity, h.getOSBBProfile)

		osbb.POST("/:osbbID/inhabitants", h.registerInhabitant)
		osbb.GET("/:osbbID/inhabitants", h.userIdentity, h.getAllInhabitants)
		osbb.GET("/:osbbID/inhabitants/profile", h.userIdentity, h.getInhabitantsProfile)
		osbb.PUT("/:osbbID/inhabitants", h.userIdentity, h.updateInhabitant)

		osbb.POST("/:osbbID/announcements", h.userIdentity, h.addAnnouncement)
		osbb.GET("/:osbbID/announcements", h.userIdentity, h.getAllAnnouncement)
		osbb.PUT("/:osbbID/announcements/:announcementID", h.userIdentity, h.updateAnnouncement)
		osbb.DELETE("/:osbbID/announcements/:announcementID", h.userIdentity, h.deleteAnnouncement)

		osbb.GET("/:osbbID/polls", h.userIdentity, h.getAllPolls)
		osbb.POST("/:osbbID/polls", h.userIdentity, h.addPoll)
		osbb.PUT("/:osbbID/polls/:pollID", h.userIdentity, h.updatePoll)
		osbb.DELETE("/:osbbID/polls/:pollID", h.userIdentity, h.deletePoll)

		osbb.POST("/:osbbID/polls-test", h.userIdentity, h.addPollTest)
		osbb.PUT("/:osbbID/polls-test/:testAnswerID", h.userIdentity, h.updateTestAnswer)
		osbb.DELETE("/:osbbID/polls-test/:testAnswerID", h.userIdentity, h.deleteTestAnswer)

		osbb.POST("/:osbbID/polls/:pollID/answers", h.userIdentity, h.addPollAnswer)
		osbb.POST("/:osbbID/polls/:pollID/answers-test", h.userIdentity, h.addPollAnswerTest)
		osbb.GET("/:osbbID/polls/:pollID/answers", h.userIdentity, h.getAllPollsAnswers)

		osbb.POST("/:osbbID/payments", h.userIdentity, h.addPayment)
		osbb.POST("/:osbbID/payments/:paymentID/make-purchase", h.userIdentity, h.makePurchase)
	}
	router.POST("refresh-token", h.refreshToken)
	return router
}
