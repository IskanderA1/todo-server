package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/IskanderA1/todo-server"
	"github.com/IskanderA1/todo-server/pkg/handler"
	"github.com/IskanderA1/todo-server/pkg/repository"
	"github.com/IskanderA1/todo-server/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error env: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMod:   viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error init db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.NewHandler(service)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error http sever: %s", err.Error())
		}
	}()
	logrus.Print("TodoApp Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	logrus.Print("TodoApp Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occured on sever shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
