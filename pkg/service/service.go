package service

import (
	"GoProjects/goservice-library"
	"GoProjects/goservice-library/pkg/repository"
)

type Authorization interface {
	CreateUser(user goservice.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list goservice.TodoList) (int, error)
	GetAll(userId int) ([]goservice.TodoList, error)
	GetById(userId int, listId int) (goservice.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input goservice.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int, listId int, item goservice.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]goservice.TodoItem, error)
	GetById(userId, itemId int) (goservice.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input goservice.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
