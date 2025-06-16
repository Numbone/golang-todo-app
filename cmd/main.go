package main

import (
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/Numbone/golang-todo-app/pkg/handler"
	"github.com/Numbone/golang-todo-app/pkg/repository"
	"github.com/Numbone/golang-todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(golang_todo_app.Server)
	if err := srv.Run(":8081", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
