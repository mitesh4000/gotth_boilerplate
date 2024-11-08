package main

import (
	"templeTest/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	UserHandler := handler.UserHandler{}
	app.Static("/static", "view")
	app.GET("/", UserHandler.HandleUserShow)
	app.Start(":3000")
}
