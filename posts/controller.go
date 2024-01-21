package posts

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

func loadMarkdownPost(slug string) template.HTML {
	content, err := os.ReadFile(fmt.Sprintf("public/markdown/%s.md", slug))
	if err != nil {
		panic(err)
	}
	return template.HTML(markdown.MDToHTML(content))
}

func GetByID(c echo.Context) error {
	slug := c.Param("slug")
	return c.Render(http.StatusOK, "post", postBody{Body: loadMarkdownPost(fmt.Sprintf("posts/%s", slug))})
}

func GetAboutPage(c echo.Context) error {
	return c.Render(http.StatusOK, "post", postBody{Body: loadMarkdownPost("about")})
}
