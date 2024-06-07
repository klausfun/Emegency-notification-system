package main

import (
	"EmegencyNotificationSystem/database"
	"EmegencyNotificationSystem/notification_service"
	"EmegencyNotificationSystem/notification_service/pkg/handler"
	"EmegencyNotificationSystem/notification_service/pkg/repository"
	"EmegencyNotificationSystem/notification_service/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	viper.SetConfigFile("configs/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	db, err := database.InitPostgres()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(notification_service.Server)
	if err := srv.Run(viper.GetString("ports.notification_service"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
