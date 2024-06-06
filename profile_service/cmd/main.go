package main

import (
	"EmegencyNotificationSystem/database"
	"EmegencyNotificationSystem/profile_service"
	"EmegencyNotificationSystem/profile_service/pkg/handler"
	"EmegencyNotificationSystem/profile_service/pkg/repository"
	"EmegencyNotificationSystem/profile_service/pkg/service"
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

	srv := new(profile_service.Server)
	if err := srv.Run(viper.GetString("ports.profile_service"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
