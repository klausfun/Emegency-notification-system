package main

import (
	"EmegencyNotificationSystem/configs"
	"EmegencyNotificationSystem/database"
	"EmegencyNotificationSystem/profile_service/pkg/handler"
	"EmegencyNotificationSystem/profile_service/pkg/repository"
	"EmegencyNotificationSystem/profile_service/pkg/service"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err)
	}

	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := database.InitPostgres()
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	http.Handle("/graphql", handlers.InitGraphQL())
	http.Handle("/playground", playground.Handler("GraphQL", "/graphql"))

	port := viper.GetString("ports.profile_service")
	logrus.Infof("connect to http://localhost:%s/playground for GraphQL playground", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
