package service

import (
	"GoProjects/goservice-library"
	"GoProjects/goservice-library/pkg/repository"
	"errors"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list goservice.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]goservice.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, listId int) (goservice.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId int, listId int, input goservice.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return errors.New("Failed to validate input")
	}
	return s.repo.Update(userId, listId, input)
}
