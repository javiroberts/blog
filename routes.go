package main

import (
	"backend/markdown"
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
	"os"
)

type postBody struct {
	Body template.HTML
}

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
	e.GET("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about", postBody{Body: loadMarkdownPost("about")})
	})
	e.GET("/post/:slug", func(c echo.Context) error {
		slug := c.Param("slug")
		return c.Render(http.StatusOK, "about", postBody{Body: loadMarkdownPost(fmt.Sprintf("posts/%s", slug))})
	})
	e.GET("/error/404", func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "404", errorTemplate{
			ErrorCode:   http.StatusNotFound,
			ErrorReason: "Article not found in database",
		})
	})
}

func loadMarkdownPost(slug string) template.HTML {
	content, err := os.ReadFile(fmt.Sprintf("public/markdown/%s.md", slug))
	if err != nil {
		panic(err)
	}
	return template.HTML(markdown.MDToHTML(content))
}
