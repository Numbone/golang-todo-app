package service

import (
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/Numbone/golang-todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo}
}

func (s *TodoListService) Create(userId int, list golang_todo_app.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]golang_todo_app.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, listId int) (golang_todo_app.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Update(userId, id int, input golang_todo_app.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}

func (s *TodoListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}
