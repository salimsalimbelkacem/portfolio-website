package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/salimsalimbelkacem/gotth-contact-app/views"
	"github.com/a-h/templ"
)

func render_templ(component templ.Component, title string) func(c echo.Context) error {
	return func(c echo.Context) error {
		return views.Layout(title).Render( templ.WithChildren(c.Request().Context(), component), c.Response().Writer)
	}
}

func SetupRoutes(e *echo.Echo){
	e.GET("/", render_templ(views.Content_one(), "pipo"))
	e.GET("/oh", render_templ(views.Content_two(), "pipoza"))

	// htmx := e.Group("/hx")
	//
	// htmx.GET("/poupe")
}
