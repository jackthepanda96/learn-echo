package controller

import (
	"21-api/middlewares"
	"21-api/model"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
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
				return c.JSON(http.StatusUnsupportedMediaType, map[string]any{"code": http.StatusUnsupportedMediaType, "message": "format data tidak didukung"})
			}
			return c.JSON(http.StatusBadRequest, map[string]any{"code": http.StatusBadRequest, "message": "data yang dikirmkan tidak sesuai"})
		}
		err = us.Model.AddUser(input) // ini adalah fungsi yang kita buat sendiri
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "terjadi kesalahan pada sistem")
		}
		return c.JSON(http.StatusCreated, "selamat data sudah terdaftar")
	}
}

func (us *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType, map[string]any{"code": http.StatusUnsupportedMediaType, "message": "format data tidak didukung"})
			}
			return c.JSON(http.StatusBadRequest, map[string]any{"code": http.StatusBadRequest, "message": "data yang dikirmkan tidak sesuai"})
		}
		result, err := us.Model.Login(input.Hp, input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "terjadi kesalahan pada sistem")
		}
		token, err := middlewares.GenerateJWT(result.Hp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"code":    http.StatusInternalServerError,
				"message": "terjadi kesalahan pada sistem, gagal memproses data",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{"code": http.StatusOK, "message": "selama anda berhasil login", "data": result, "token": token})

	}
}

func (us *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hp = c.Param("hp")
		var input model.User
		err := c.Bind(&input)
		if err != nil {
			log.Println("masalah baca input:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType, err.Error())
			}
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		isFound := us.Model.CekUser(hp)

		if !isFound {
			return c.JSON(http.StatusNotFound, "data tidak ditemukan")
		}

		err = us.Model.Update(hp, input)

		if err != nil {
			log.Println("masalah database :", err.Error())
			return c.JSON(http.StatusInternalServerError, "terjadi kesalahan saat update data")
		}

		return c.JSON(http.StatusOK, "data berhasil di update")
	}
}

func (us *UserController) ListUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		listUser, err := us.Model.GetAllUser()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"code":    http.StatusInternalServerError,
				"message": "terjadi kesalahan pada sistem",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"code":    http.StatusOK,
			"message": "berhasil mendapatkan data",
			"data":    listUser,
		})
	}
}

func (us *UserController) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hpFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))

		result, err := us.Model.GetProfile(hpFromToken)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]any{
					"code":    http.StatusNotFound,
					"message": "data tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"code":    http.StatusInternalServerError,
				"message": "terjadi kesalahan pada sistem",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"code":    http.StatusOK,
			"message": "berhasil mendapatkan data",
			"data":    result,
		})
	}
}
