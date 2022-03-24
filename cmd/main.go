package main

import (
	"log"

	"github.com/IskanderA1/todo-server"
	"github.com/IskanderA1/todo-server/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error http sever: %s", err.Error())
	}
}
