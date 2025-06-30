package service

import (
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/Numbone/golang-todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo, listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item golang_todo_app.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]golang_todo_app.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (golang_todo_app.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input golang_todo_app.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
