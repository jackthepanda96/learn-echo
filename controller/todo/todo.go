package todo

import (
	"21-api/helper"
	"21-api/middlewares"
	"21-api/model/todo"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	Model todo.TodoModel
}

func (tc *TodoController) AddToDo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ToDoRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		// Cek middleware (extract token)
		// c.Get("user").(*jwt.Token) -> notasi PASTI kalo mau mengambil jwt token pada echo

		hp := middlewares.DecodeToken(c.Get("user").(*jwt.Token))

		if hp == "" {
			log.Println("error decode token:", "hp tidak ditemukan")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "tidak dapat mengakses fitur ini", nil))
		}

		var inputProcess todo.Todo
		inputProcess.Kegiatan = input.Kegiatan
		inputProcess.Pemilik = hp

		result, err := tc.Model.Insert(inputProcess)

		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada proses server", nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil menambahkan kegiatan", result))
	}
}

func (tc *TodoController) UpdateToDo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ToDoRequest

		readID := c.Param("todoID")
		cnv, err := strconv.Atoi(readID)
		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}
		err = c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		hp := middlewares.DecodeToken(c.Get("user").(*jwt.Token))

		if hp == "" {
			log.Println("error decode token:", "hp tidak ditemukan")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "tidak dapat mengakses fitur ini", nil))
		}

		var inputProcess todo.Todo
		inputProcess.Kegiatan = input.Kegiatan

		result, err := tc.Model.UpdateKegiatan(hp, uint(cnv), inputProcess)

		if err != nil {
			log.Println("error update db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada proses server", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "berhasil mengubah data kegiatan", result))
	}

}
