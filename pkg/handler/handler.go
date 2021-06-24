package handler

import (
	"cevent/pkg/service"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"time"
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

	api := router.Group("api", h.userIdentity)
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

func saveImage(file multipart.File, header *multipart.FileHeader, err error) (string, error) {
	if err != nil {
		return "", err
	}
	filename := generateRandomString(12) + ".png"

	out, err := os.Create("./upload/" + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	s := make([]byte, length)
	for i := range s {
		s[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(s)
}
