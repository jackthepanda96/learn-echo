package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Nama string
	// Nama     string `json:"nama" form:"nama"`
	Hp       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", func(c echo.Context) error {
		var daftarUser = make([]User, 0)
		daftarUser = append(daftarUser, User{Nama: "Jerry"})
		daftarUser = append(daftarUser, User{Nama: "Malik"})

		return c.JSON(http.StatusOK, daftarUser)
	})
	e.POST("/users", func(c echo.Context) error {
		var input User
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, input)
	})
	//users/089
	//path parameter -> biasa digunakan untuk mencari SEBUAH data secara spesifik
	e.DELETE("/users/:userID", func(c echo.Context) error {
		var userID = c.Param("userID")
		return c.JSON(http.StatusOK, userID)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
