package handler

import (
	"cevent/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("api")
	{
		event := api.Group("/events")
		{
			event.POST("/", h.createEvent)
			event.PUT("/:id", h.updateEvent)
			event.DELETE("/:id", h.deleteEvent)
		}
	}

	return router
}
