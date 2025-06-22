package service

import (
	"crypto/sha1"
	"fmt"
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/Numbone/golang-todo-app/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user golang_todo_app.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte("asdasdqwe")))
}
