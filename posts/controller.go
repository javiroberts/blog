package posts

import (
	"backend/markdown"
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
)

type post struct {
	Article template.HTML
}

func GetByID(c echo.Context) error {
	slug := c.Param("slug")
	return c.Render(http.StatusOK, "post", post{Article: markdown.LoadMarkdownPost(fmt.Sprintf("posts/%s", slug))})
}

func GetAboutPage(c echo.Context) error {
	return c.Render(http.StatusOK, "post", post{Article: markdown.LoadMarkdownPost("about")})
}
