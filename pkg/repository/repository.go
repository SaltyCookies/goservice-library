package repository

import (
	"GoProjects/goservice-library"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user goservice.User) (int, error)
	GetUser(username, password string) (goservice.User, error)
}

type TodoList interface {
	Create(userId int, list goservice.TodoList) (int, error)
	GetAll(userId int) ([]goservice.TodoList, error)
	GetById(userId int, listId int) (goservice.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input goservice.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item goservice.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]goservice.TodoItem, error)
	GetById(userId, itemId int) (goservice.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input goservice.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
