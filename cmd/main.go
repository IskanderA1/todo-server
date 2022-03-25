package main

import (
	"log"

	"github.com/IskanderA1/todo-server"
	"github.com/IskanderA1/todo-server/pkg/handler"
	"github.com/IskanderA1/todo-server/pkg/repository"
	"github.com/IskanderA1/todo-server/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handlers := handler.NewHandler(service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error http sever: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
