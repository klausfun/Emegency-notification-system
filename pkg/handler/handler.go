package handler

import (
	"EmegencyNotificationSystem/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}

	api := router.Group("/api")
	{
		notifications := api.Group("/notifications")
		{
			notifications.POST("/", h.createNotification)
			notifications.GET("/", h.getNotifications)
			notifications.GET("/:id", h.getNotificationById)
			notifications.PUT("/", h.updateNotification)
			notifications.DELETE("/", h.deleteNotification)
			notifications.POST("/send", h.sendNotification)
		}
	}

	return router
}
