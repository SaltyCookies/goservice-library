package main

import (
	"GoProjects/goservice-library"
	"GoProjects/goservice-library/pkg/handler"
	"GoProjects/goservice-library/pkg/repository"
	"GoProjects/goservice-library/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBname:   "postgres",
		SSLmode:  "disable",
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatal("failed to connect to database", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(goservice.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Fatal(err.Error())
	}
}
