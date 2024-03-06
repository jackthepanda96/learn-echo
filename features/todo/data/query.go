package data

import (
	"21-api/features/todo"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) todo.TodoModel {
	return &model{
		connection: db,
	}
}

func (tm *model) InsertTodo(kegiatanBaru todo.Todo) (todo.Todo, error) {
	if err := tm.connection.Create(&kegiatanBaru).Error; err != nil {
		return todo.Todo{}, err
	}

	return kegiatanBaru, nil
}

func (tm *model) UpdateTodo(pemilik string, todoID uint, data todo.Todo) (todo.Todo, error) {
	var qry = tm.connection.Where("pemilik = ? AND id = ?", pemilik, todoID).Updates(data)
	if err := qry.Error; err != nil {
		return todo.Todo{}, err
	}

	if qry.RowsAffected < 1 {
		return todo.Todo{}, errors.New("no data affected")
	}

	return data, nil
}

func (tm *model) GetTodoByOwner(pemilik string) ([]todo.Todo, error) {
	var result []todo.Todo
	if err := tm.connection.Where("pemilik = ?", pemilik).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
