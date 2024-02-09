package main

import (
	"backend/pages"
	"backend/posts"
	"github.com/labstack/echo/v4"
	"net/http"
)

type errorTemplate struct {
	ErrorCode   int
	ErrorReason string
}

func registerStaticRoutes(e *echo.Echo) {
	e.Static("/static", "public/static")
}

func registerRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})
	e.GET("/about", pages.GetAboutPage)
	e.GET("/posts", posts.GetPostList)
	e.GET("/posts/:slug", posts.GetPost)
	e.GET("/error/404", func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "404", errorTemplate{
			ErrorCode:   http.StatusNotFound,
			ErrorReason: "Article not found in database",
		})
	})
}
