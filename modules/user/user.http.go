package user

import "github.com/labstack/echo/v4"

func UserHTTP(e *echo.Echo) {
	e.GET("/users", GetAllUsers)
	e.GET("/user/:id", GetUser)
	e.POST("/register", RegisterUser)
	e.POST("/login", Login)
}
