package model

import "gorm.io/gorm"

type User struct {
	Nama     string `json:"nama" form:"nama"`
	Hp       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

type UserModel struct {
	Connection *gorm.DB
}

func (um *UserModel) AddUser(newData User) error {
	err := um.Connection.Create(&newData).Error
	if err != nil {
		return err
	}

	return nil
}
