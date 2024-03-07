package todo

import "github.com/labstack/echo/v4"

type TodoController interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// ShowMyTodo() echo.HandlerFunc
}

type TodoModel interface {
	InsertTodo(pemilik string, kegiatanBaru Todo) (Todo, error)
	UpdateTodo(pemilik string, todoID uint, data Todo) (Todo, error)
	// DeleteTodo()
	GetTodoByOwner(pemilik string) ([]Todo, error)
}

type TodoService interface {
}

type Todo struct {
	Kegiatan string
}
