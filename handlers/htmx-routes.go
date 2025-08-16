package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/salimsalimbelkacem/gotth-contact-app/views"
)

func setupHxRoutes(e *echo.Echo) (htmx *echo.Group){
	htmx = e.Group("/hx")
	htmx.GET("/home", rendTempl(views.Home(), nil))
	htmx.GET("/about", rendTempl(views.About(), nil))

	return 
}
