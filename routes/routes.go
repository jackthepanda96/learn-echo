package routes

import (
	"21-api/config"
	todo "21-api/features/todo"
	user "21-api/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController, tc todo.TodoController) {
	userRoute(c, ctl)
	todoRoute(c, tc)
}

func userRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/users", ctl.Add())
	c.POST("/login", ctl.Login())
	c.GET("/profile", ctl.Profile(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}

func todoRoute(c *echo.Echo, tc todo.TodoController) {
	c.POST("/todos", tc.Add(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
