package todo

import (
	"errors"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Kegiatan string
	Pemilik  string `gorm:"type:varchar(13);"`
}

type TodoModel struct {
	Connection *gorm.DB
}

func (tm *TodoModel) Insert(kegiatanBaru Todo) (Todo, error) {
	if err := tm.Connection.Create(&kegiatanBaru).Error; err != nil {
		return Todo{}, err
	}

	return kegiatanBaru, nil
}

func (tm *TodoModel) ListKegiatan(pemilik string) ([]Todo, error) {
	var result []Todo
	if err := tm.Connection.Where("pemilik = ?", pemilik).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (tm *TodoModel) UpdateKegiatan(pemilik string, todoID uint, data Todo) (Todo, error) {
	var qry = tm.Connection.Where("pemilik = ? AND id = ?", pemilik, todoID).Updates(data)
	if err := qry.Error; err != nil {
		return Todo{}, err
	}

	if qry.RowsAffected < 1 {
		return Todo{}, errors.New("no data affected")
	}

	return data, nil
}
