package service

import (
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/Numbone/golang-todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user golang_todo_app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list golang_todo_app.TodoList) (int, error)
	GetAll(userId int) ([]golang_todo_app.TodoList, error)
	GetById(userId int, listId int) (golang_todo_app.TodoList, error)
	Delete(userId int, listId int) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}

}
