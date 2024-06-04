package main

import (
	"EmegencyNotificationSystem"
	"EmegencyNotificationSystem/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(EmegencyNotificationSystem.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
