package controller

import (
	"21-api/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model model.UserModel
}

func (us *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType, err.Error())
			}
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err = us.Model.AddUser(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "terjadi kesalahan pada sistem")
		}
		return c.JSON(http.StatusCreated, "selamat data sudah terdaftar")
	}
}
