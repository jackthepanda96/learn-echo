package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var daftarUser = make([]User, 0)

type User struct {
	Nama     string `json:"nama" form:"nama"`
	Hp       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

func LihatListUser(c echo.Context) error {
	if len(daftarUser) == 0 {
		return c.JSON(http.StatusOK, "data empty")
	}
	return c.JSON(http.StatusOK, daftarUser)
}

func Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType, err.Error())
			}
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		daftarUser = append(daftarUser, input)
		return c.JSON(http.StatusCreated, "selamat data sudah terdaftar")
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", LihatListUser)
	e.POST("/users", Register())
	//users/089
	//path parameter -> biasa digunakan untuk mencari SEBUAH data secara spesifik
	e.DELETE("/users/:userID", func(c echo.Context) error {
		var userID = c.Param("userID")
		return c.JSON(http.StatusOK, userID)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
