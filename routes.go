package main

import (
	"blog/pages"
	"blog/posts"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
}
