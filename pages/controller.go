package pages

import (
	"backend/markdown"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"path/filepath"
)

func GetAboutPage(c echo.Context) error {
	dir := "public/markdown/"
	slug := "about"
	page := loadPage(dir, slug)
	return c.Render(http.StatusOK, "about", PageSchema{Page: page})
}

func loadPage(dir, slug string) Page {
	f := markdown.MDFile{}
	f.Load(filepath.Join(dir, fmt.Sprintf("%s.md", slug)))

	p := Page{}
	p.WithMarkdown(f)

	return p
}
