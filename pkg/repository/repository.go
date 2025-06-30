package repository

import (
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user golang_todo_app.User) (int, error)
	GetUser(username, password string) (golang_todo_app.User, error)
}

type TodoList interface {
	Create(userId int, list golang_todo_app.TodoList) (int, error)
	GetAll(userId int) ([]golang_todo_app.TodoList, error)
	GetById(userId int, listId int) (golang_todo_app.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, id int, list golang_todo_app.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item golang_todo_app.TodoItem) (int, error)
	GetAll(userId, listId int) ([]golang_todo_app.TodoItem, error)
	GetById(userId, itemId int) (golang_todo_app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input golang_todo_app.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
