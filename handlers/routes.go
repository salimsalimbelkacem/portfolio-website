package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/salimsalimbelkacem/gotth-contact-app/views"
	"github.com/a-h/templ"
)

func render_templ(Layout templ.Component, component templ.Component) func(c echo.Context) error {
	return func(c echo.Context) error {
		return Layout.Render( templ.WithChildren(c.Request().Context(), component), c.Response().Writer)
	}
}

func SetupRoutes(e *echo.Echo){
	e.GET("/", render_templ(views.Layout("Home"), views.Home()))

	// htmx := e.Group("/hx")
	//
	// htmx.GET("/poupe")
}
