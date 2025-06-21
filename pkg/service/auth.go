package service

import (
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/Numbone/golang-todo-app/pkg/repository"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user *golang_todo_app.User) (string, error) {
	return s.repo.CreateUser(user)
}
