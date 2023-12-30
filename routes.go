package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func registerStaticRoutes(e *echo.Echo) {
	e.Static("/static", "public/static")
}

func registerRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})
}
