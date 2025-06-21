package repository

import (
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{}
}

func (r *AuthPostgres) CreateUser(user *golang_todo_app.User) {
	return 0, int
}
