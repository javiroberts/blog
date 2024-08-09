package main

import (
	"blog/template"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

func main() {
	logger := zerolog.New(os.Stdout)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.Info().
					Str("URI", v.URI).
					Int("status", v.Status).
					Send()
			} else {
				logger.Error().
					Str("URI", v.URI).
					Int("status", v.Status).
					Msg(v.Error.Error())
			}
			return nil
		},
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {

	}

	template.NewTemplateRenderer(e, "public/templates/**/*.gohtml")

	registerStaticRoutes(e)
	registerRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
