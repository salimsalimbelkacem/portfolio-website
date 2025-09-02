package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/salimsalimbelkacem/idk-what-i-am-doing-at-this-point/views"
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

	e.GET("/", func(c echo.Context) error {
		title := c.QueryParam("title")

		var component templ.Component = views.Routes["/"+title]
		if component == nil {component = views.Home()}

		return rendTempl(
			views.Layout(title),
			component,
			)(c)
	})
	setupHxRoutes(e)
}
