package main

import (
	"github.com/labstack/echo/v4"
	user "github.com/thitiphongD/api-echo/modules/healthcheck"
)

func main() {
	e := echo.New()
	user.HealthCheckHTTP(e)
	e.Logger.Fatal(e.Start(":8080"))
}
