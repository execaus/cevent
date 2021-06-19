package main

import (
	"cevent"
	"cevent/pkg/handler"
	"cevent/pkg/repository"
	"cevent/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(cevent.Server)
	if err := server.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
