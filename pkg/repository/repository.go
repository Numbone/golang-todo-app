package repository

import (
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user *golang_todo_app.User) (int, err)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
