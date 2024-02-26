package routes

import (
	"21-api/controller"

	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl controller.UserController) {
	c.POST("/users", ctl.Register())
	c.PUT("/users/:hp", ctl.Update())
}
