package service

import (
	"crypto/sha1"
	"fmt"
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/Numbone/golang-todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signKey = "asdasdqwe"
)

type AuthService struct {
	repo repository.Authorization
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user golang_todo_app.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})
	return token.SignedString([]byte(signKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(signKey)))
}

func (s *AuthService) ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signKey), nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("token is not valid")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, fmt.Errorf("token claims are not of type *TokenClaims")
	}

	return claims.UserId, nil
}
