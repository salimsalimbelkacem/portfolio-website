package main

import (
	// "github.com/salimsalimbelkacem/gotth-contact-app/views"
	"github.com/labstack/echo/v4"
	"github.com/salimsalimbelkacem/gotth-contact-app/handlers"
)


func main(){
	e := echo.New()

	e.Static("/static", "static")

	handlers.SetupRoutes(e)

	e.Start(":5000")
}
