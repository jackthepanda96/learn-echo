package user

import "github.com/labstack/echo/v4"

type UserController interface {
	Add() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, error)
}

type UserModel interface {
	InsertUser(newData User) error
	UpdateUser(hp string, data User) error
	Login(hp string, password string) (User, error)
	GetUserByHP(hp string) (User, error)
}

type User struct {
	Nama     string
	HP       string
	Password string
}
