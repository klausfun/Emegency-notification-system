package handler

import (
	"EmegencyNotificationSystem/notification_service/pkg/service"
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

	api := router.Group("/api")
	{
		notifications := api.Group("/notifications")
		{
			notifications.POST("/", h.createNotification)
			notifications.GET("/", h.getNotifications)
		}
	}

	return router
}
