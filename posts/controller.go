package posts

import (
	"backend/markdown"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func GetByID(c echo.Context) error {
	slug := c.Param("slug")
	return c.Render(http.StatusOK, "post", PostSchema{Article: markdown.LoadMarkdownPost(fmt.Sprintf("posts/%s", slug))})
}

func GetAboutPage(c echo.Context) error {
	return c.Render(http.StatusOK, "post", PostSchema{Article: markdown.LoadMarkdownPost("about")})
}

func GetPostList(c echo.Context) error {
	postList := buildPostList("public/markdown/posts/")
	return c.Render(http.StatusOK, "list", PostListSchema{List: postList})
}

func buildPostList(dir string) []Post {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	p := Post{}
	ps := []Post{}
	for _, slug := range dirEntries {
		p.Load(slug.Name())
		ps = append(ps, p)
	}
	return ps
}
