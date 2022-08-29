package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"quanlyhoso/config"
	"quanlyhoso/database"
	"quanlyhoso/route"
)

func main() {
	e := echo.New()
	config.InitDotEnv()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	database.Connect()

	route.Route(e)
	e.Logger.Fatal(e.Start(":" + config.GetEnv().AppPort))
}
