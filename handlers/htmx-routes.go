package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/salimsalimbelkacem/idk-what-i-am-doing-at-this-point/views"
)

func setupHxRoutes(e *echo.Echo) (htmx *echo.Group){
	htmx = e.Group("/hx")
	for path, Component := range views.Routes {
		htmx.GET(path, rendTempl(Component, nil))
	}

	return 
}
