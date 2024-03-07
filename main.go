package main

import (
	"21-api/config"
	td "21-api/features/todo/data"
	th "21-api/features/todo/handler"
	ts "21-api/features/todo/services"
	"21-api/features/user/data"
	"21-api/features/user/handler"
	"21-api/features/user/services"
	"21-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // konek DB

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	todoData := td.New(db)
	todoService := ts.NewTodoService(todoData)
	todoHandler := th.NewHandler(todoService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup
	routes.InitRoute(e, userHandler, todoHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
