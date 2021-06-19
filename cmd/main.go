package main

import (
	"cevent"
	"cevent/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	server := new(cevent.Server)
	if err := server.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
