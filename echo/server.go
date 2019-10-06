package main

import (
	"github.com/ookayama-playgrounds/go-playgrounds/echo/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/hello", handler.Hello)
	e.GET("/hey/:email", handler.GetHello)
	e.POST("/hey", handler.InsertHello)
	e.Logger.Fatal(e.Start(":8080"))
}

