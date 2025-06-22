package repository

import (
	"fmt"
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) CreateUser(user golang_todo_app.User) (int, error) {
	if r.db == nil {
		return 0, fmt.Errorf("DB connection is nil")
	}

	query := fmt.Sprintf(`INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id`, userTable)

	var id int
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
