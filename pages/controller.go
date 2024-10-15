package pages

import (
	"blog/markdown"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func GetAboutPage(c echo.Context) error {
	dir := "public/markdown/"
	slug := "about"
	err, page := loadPage(dir, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.Render(http.StatusOK, "about", PageSchema{Page: page})
}

func loadPage(dir, slug string) (error, Page) {
	f := markdown.MDFile{}
	err := f.Load(filepath.Join(dir, fmt.Sprintf("%s.md", slug)))
	if err != nil {
		return err, Page{}
	}

	p := Page{}
	err = p.WithMarkdown(f)
	if err != nil {
		return err, Page{}
	}

	return nil, p
}
