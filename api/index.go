package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/a-h/templ"

	"github.com/salimsalimbelkacem/portfolio-website/views"
)

func rendTempl(layout templ.Component, component templ.Component) func(c echo.Context) error {
	return func(c echo.Context) error {
		if layout != nil {
			return layout.Render(templ.WithChildren(c.Request().Context(), component), c.Response().Writer)
		}

		return component.Render(c.Request().Context(), c.Response().Writer)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	e.Use((middleware.Logger()))

	e.Static("/static", "static")

	e.GET("/", rendTempl( views.Layout(), views.Home(),))

	htmx := e.Group("/hx")

	for path, Component := range views.Routes {
		htmx.GET("/api/"+path, rendTempl(Component, nil))
	}

	e.ServeHTTP(w, r)
}
