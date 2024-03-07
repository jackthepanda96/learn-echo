package services

import (
	"21-api/features/user"
	"21-api/helper"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

type service struct {
	model user.UserModel
	pm    helper.PasswordManager
	v     *validator.Validate
}

func NewService(m user.UserModel) user.UserService {
	return &service{
		model: m,
		pm:    helper.NewPasswordManager(),
		v:     validator.New(),
	}
}

func (s *service) Register(newData user.User) error {
	var registerValidate user.Register
	registerValidate.Hp = newData.Hp
	registerValidate.Nama = newData.Nama
	registerValidate.Password = newData.Password
	err := s.v.Struct(&registerValidate)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}

	newPassword, err := s.pm.HashPassword(newData.Password)
	if err != nil {
		return errors.New(helper.ServiceGeneralError)
	}
	newData.Password = newPassword

	err = s.model.InsertUser(newData)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}
	return nil
}
func (s *service) Login(loginData user.User) (user.User, error) {
	return user.User{}, nil
}

func (s *service) Profile(hp string) (user.User, error) {
	return user.User{}, nil
}
