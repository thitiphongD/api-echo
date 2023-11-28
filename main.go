package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thitiphongD/api-echo/db"
	check "github.com/thitiphongD/api-echo/modules/healthcheck"
	user "github.com/thitiphongD/api-echo/modules/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}

	defer func() {
		sqlDB, _ := database.DB()
		_ = sqlDB.Close()
	}()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	check.HealthCheckHTTP(e)
	user.UserHTTP(e)
	e.Logger.Fatal(e.Start(":8080"))
}
