package main

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/salimsalimbelkacem/portfolio-website/views"
)

// --- Templ render helper ---
func rendTempl(layout templ.Component, component templ.Component) func(c echo.Context) error {
	return func(c echo.Context) error {
		if layout != nil {
			return layout.Render(templ.WithChildren(c.Request().Context(), component), c.Response().Writer)
		}
		return component.Render(c.Request().Context(), c.Response().Writer)
	}
}

// --- Echo app setup ---
var app *echo.Echo

func init() {
	app = echo.New()
	app.Use(middleware.Logger())
	app.Static("/public/static", "static")

	app.GET("/", rendTempl(views.Layout(), views.Home()))

	for path, Component := range views.Routes {
		println(path)
		app.GET("/hx/"+path, rendTempl(Component, nil))
	}
}

// --- Lambda adapter ---
type responseCapture struct {
	header     http.Header
	bodyBuffer bytes.Buffer
	statusCode int
}

func newResponseCapture() *responseCapture {
	return &responseCapture{
		header:     make(http.Header),
		statusCode: 200,
	}
}

func (r *responseCapture) Header() http.Header         { return r.header }
func (r *responseCapture) Write(b []byte) (int, error) { return r.bodyBuffer.Write(b) }
func (r *responseCapture) WriteHeader(code int)        { r.statusCode = code }

// --- Lambda handler ---
func handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// Create an HTTP request Echo can understand
	body := io.NopCloser(bytes.NewBufferString(req.Body))
	r, err := http.NewRequestWithContext(ctx, req.HTTPMethod, req.Path, body)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, "failed to create request"), nil
	}

	// Copy headers from API Gateway
	for k, v := range req.Headers {
		r.Header.Set(k, v)
	}

	// Capture Echo’s output
	w := newResponseCapture()
	app.ServeHTTP(w, r)

	// Prepare response
	return &events.APIGatewayProxyResponse{
		StatusCode: w.statusCode,
		Headers: map[string]string{
			"Content-Type": w.Header().Get("Content-Type"),
		},
		Body:            w.bodyBuffer.String(),
		IsBase64Encoded: false,
	}, nil
}

func errorResponse(status int, msg string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       msg,
	}
}

// --- Entrypoint ---
func main() {
	lambda.Start(handler)
}
