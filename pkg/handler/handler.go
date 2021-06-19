package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	api := router.Group("api")
	{
		event := api.Group("/events")
		{
			event.POST("/")
			event.PUT("/:id")
			event.DELETE("/:id")
		}
	}

	return router
}
