package data

import (
	"21-api/features/user"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

// yang kita butuhkan adalah sebuah model
// tapi kenapa return function kok bukan obyek model?

func New(db *gorm.DB) user.UserModel {
	return &model{
		connection: db,
	}
}

func (m *model) InsertUser(newData user.User) error {
	err := m.connection.Model(&User{}).Create(&newData).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *model) cekUser(hp string) bool {
	var data User
	if err := m.connection.Where("hp = ?", hp).First(&data).Error; err != nil {
		return false
	}
	return true
}

func (m *model) UpdateUser(hp string, data user.User) error {
	if err := m.connection.Model(&data).Where("hp = ?", hp).Update("nama", data.Nama).Update("password", data.Password).Error; err != nil {
		return err
	}
	return nil
}

func (m *model) GetAllUser() ([]user.User, error) {
	var result []user.User

	if err := m.connection.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (m *model) GetUserByHP(hp string) (user.User, error) {
	var result user.User
	if err := m.connection.Model(&User{}).Where("hp = ?", hp).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (m *model) Login(hp string, password string) (user.User, error) {
	var result user.User
	if err := m.connection.Model(&User{}).Where("hp = ? AND password = ?", hp, password).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}
