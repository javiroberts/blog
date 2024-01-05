package main

import (
	"backend/markdown"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
	"os"
)

type postBody struct {
	Body template.HTML
}

func registerStaticRoutes(e *echo.Echo) {
	e.Static("/static", "public/static")
}

func registerRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})
	e.GET("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about", postBody{Body: loadMarkdownPost()})
	})
}

func loadMarkdownPost() template.HTML {
	content, err := os.ReadFile("public/markdown/about.md")
	if err != nil {
		panic(err)
	}
	return template.HTML(markdown.MDToHTML(content))
}
