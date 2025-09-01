package main

import (
	// "github.com/salimsalimbelkacem/idk-what-i-am-doing-at-this-point/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/salimsalimbelkacem/idk-what-i-am-doing-at-this-point/handlers"
)


func main(){
	e := echo.New()

	e.Use((middleware.Logger()))

	e.Static("/static", "static")

	handlers.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":80"))
}
