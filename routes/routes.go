package routes

import (
	"21-api/config"
	todo "21-api/features/todo"
	user "21-api/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// func InitRoute(c *echo.Echo, ctl user.UserController, tc todo.TodoController) {
// 	c.POST("/users", ctl.Register()) // register -> umum (boleh diakses semua orang)
// 	c.POST("/login", ctl.Login())
// 	c.GET("/users", ctl.ListUser(), echojwt.WithConfig(echojwt.Config{
// 		SigningKey: []byte(config.JWTSECRET),
// 	})) // get all user -> butuh penanda khusus
// 	c.GET("/profile", ctl.Profile(), echojwt.WithConfig(echojwt.Config{
// 		SigningKey: []byte(config.JWTSECRET),
// 	})) // get profile -> butuh penanda khusus
// 	c.PUT("/users/:hp", ctl.Update(), echojwt.WithConfig(echojwt.Config{
// 		SigningKey: []byte(config.JWTSECRET),
// 	})) // update user -> butuh penanda khusus

// 	c.POST("/todos", tc.AddToDo(), echojwt.WithConfig(echojwt.Config{
// 		SigningKey: []byte(config.JWTSECRET),
// 	}))
// 	// c.GET("/todos", tc.AddToDo(), echojwt.WithConfig(echojwt.Config{
// 	// 	SigningKey: []byte(config.JWTSECRET),
// 	// }))
// 	c.PUT("/todos/:todoID", tc.UpdateToDo(), echojwt.WithConfig(echojwt.Config{
// 		SigningKey: []byte(config.JWTSECRET),
// 	}))
// }

func InitRoute(c *echo.Echo, ctl user.UserController, tc todo.TodoController) {
	userRoute(c, ctl)
	todoRoute(c, tc)
}

func userRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/users", ctl.Add()) // register -> umum (boleh diakses semua orang)
	c.POST("/login", ctl.Login())
	// c.GET("/users", ctl.ListUser(), echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte(config.JWTSECRET),
	// })) // get all user -> butuh penanda khusus
	c.GET("/profile", ctl.Profile(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	})) // get profile -> butuh penanda khusus
	// c.PUT("/users/:hp", ctl.Update(), echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte(config.JWTSECRET),
	// })) // update user -> butuh penanda khusus
}

func todoRoute(c *echo.Echo, tc todo.TodoController) {
	// c.POST("/todos", tc.AddToDo(), echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte(config.JWTSECRET),
	// }))
	// // c.GET("/todos", tc.AddToDo(), echojwt.WithConfig(echojwt.Config{
	// // 	SigningKey: []byte(config.JWTSECRET),
	// // }))
	// c.PUT("/todos/:todoID", tc.UpdateToDo(), echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte(config.JWTSECRET),
	// }))
}
