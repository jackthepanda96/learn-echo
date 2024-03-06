package handler

import (
	"21-api/features/todo"
	"21-api/helper"
	"21-api/middlewares"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	model todo.TodoModel
}

func NewHandler(m todo.TodoModel) todo.TodoController {
	return &controller{
		model: m,
	}
}

func (ct *controller) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ToDoRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		hp := middlewares.DecodeToken(c.Get("user").(*jwt.Token))

		if hp == "" {
			log.Println("error decode token:", "hp tidak ditemukan")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "tidak dapat mengakses fitur ini", nil))
		}

		var inputProcess todo.Todo
		inputProcess.Kegiatan = input.Kegiatan
		inputProcess.Pemilik = hp

		result, err := ct.model.InsertTodo(inputProcess)

		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil menambahkan kegiatan", result))
	}
}

func (ct *controller) Update() echo.HandlerFunc {
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

		result, err := ct.model.UpdateTodo(hp, uint(cnv), inputProcess)

		if err != nil {
			log.Println("error update db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada proses server", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "berhasil mengubah data kegiatan", result))
	}

}
