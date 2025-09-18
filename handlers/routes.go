package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/salimsalimbelkacem/portfolio-website/views"
	"github.com/a-h/templ"
)

func rendTempl(layout templ.Component, component templ.Component) func(c echo.Context) error {
	return func(c echo.Context) error {
		if layout != nil {
			return layout.Render(templ.WithChildren(c.Request().Context(), component), c.Response().Writer)
		}

		return component.Render(c.Request().Context(), c.Response().Writer)
	}
}



func SetupRoutes(e *echo.Echo){
	// e.GET("/", rendTempl(views.Layout("Home"), views.Home()))

	e.GET("/", rendTempl( views.Layout(), views.Home(),))
	setupHxRoutes(e)
}
