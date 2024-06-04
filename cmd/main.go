package main

import (
	"EmegencyNotificationSystem"
	"EmegencyNotificationSystem/pkg/handler"
	"EmegencyNotificationSystem/pkg/repository"
	"EmegencyNotificationSystem/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(EmegencyNotificationSystem.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
