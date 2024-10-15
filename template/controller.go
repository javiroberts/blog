package template

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(e *echo.Echo, path string) {
	t := &Template{
		templates: template.Must(template.ParseGlob(path)),
	}
	e.Renderer = t
}
