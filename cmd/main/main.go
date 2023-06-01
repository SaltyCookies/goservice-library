package main

import (
	"GoProjects/goservice-library"
	"GoProjects/goservice-library/pkg/handler"
	"GoProjects/goservice-library/pkg/repository"
	"GoProjects/goservice-library/pkg/service"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title Todo App API
// @version 1.0
// @description API server for todo list app

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	curDir, err := os.Getwd()
	if err != nil {
		logrus.Println(err)
	}
	logrus.Println(curDir)
	if err := godotenv.Load(curDir + "/.env"); err != nil {
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

	go func() {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			logrus.Fatal(err.Error())
		}
	}()

	logrus.Print("Todo app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server GRACEFUL SHUTDOWN SUCCESSFUL")

	if err := srv.Stop(context.Background()); err != nil {
		logrus.Error("Error shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Error("Error shutting down: %s", err.Error())
	}

}
