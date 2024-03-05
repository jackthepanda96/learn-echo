package main

import (
	"21-api/config"
	tControll "21-api/controller/todo"
	uControll "21-api/controller/user"
	"21-api/model/todo"
	"21-api/model/user"
	"21-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // konek DB

	m := user.UserModel{Connection: db}     // bagian yang menghungkan coding kita ke database / bagian dimana kita ngoding untk ke DB
	c := uControll.UserController{Model: m} // bagian yang menghandle segala hal yang berurusan dengan HTTP / echo
	tm := todo.TodoModel{Connection: db}
	tc := tControll.TodoController{Model: tm}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup
	routes.InitRoute(e, c, tc)
	e.Logger.Fatal(e.Start(":1323"))
}
