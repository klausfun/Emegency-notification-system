package main

import (
	"EmegencyNotificationSystem/email_service"
	"EmegencyNotificationSystem/email_service/pkg/service"
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

	services := service.NewService()
	// for example
	err := services.SendEmail()
	if err != nil {
		log.Fatalf("error SendEmail(): %s", err.Error())
	}

	srv := new(email_service.Server)
	if err := srv.Run(viper.GetString("ports.email_service")); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
