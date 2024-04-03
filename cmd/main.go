package main

import (
	"github.com/arthurwieb/got/handler"
	"github.com/labstack/echo/v4"
	// https://echo.labstack.com/
)

type DB struct{}

func main() {
	// fmt.Println("Hello, World!")
	app := echo.New()
	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")

}
