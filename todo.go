package goservice

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
	ListId int `json:"listId"`
}

type TodoItem struct {
	Id          int
	Title       string
	Description string
	Done        bool
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
