package user

import (
	"github.com/labstack/echo/v4"
)

func HealthCheckHTTP(e *echo.Echo) {
	e.GET("/echo", HealthCheck)
}
