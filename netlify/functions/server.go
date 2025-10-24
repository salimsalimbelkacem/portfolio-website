package main

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

var app *echo.Echo

func init(){
	app = echo.New()

	app.Use((middleware.Logger()))

	app.Static("/public/static", "static")

	app.GET("/", rendTempl( views.Layout(), views.Home(),))

	for path, Component := range views.Routes {
		println(path)
		app.GET("/hx/"+path, rendTempl(Component, nil))
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
