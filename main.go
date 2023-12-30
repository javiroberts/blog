package main

import (
	"backend/template"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	template.NewTemplateRenderer(e, "public/templates/**/*.gohtml")

	registerStaticRoutes(e)
	registerRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
