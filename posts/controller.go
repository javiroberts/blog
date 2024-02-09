package posts

import (
	"backend/markdown"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path/filepath"
)

//func GetByID(c echo.Context) error {
//	slug := c.Param("slug")
//	return c.Render(http.StatusOK, "post", PostSchema{Article: markdown.LoadMarkdownPost(fmt.Sprintf("posts/%s", slug))})
//}

//func GetAboutPage(c echo.Context) error {
//	return c.Render(http.StatusOK, "post", PostSchema{Article: markdown.LoadMarkdownPost("about")})
//}

func GetPost(c echo.Context) error {
	dir := "public/markdown/posts/"
	slug := c.Param("slug")
	post := loadPost(dir, slug)
	return c.Render(http.StatusOK, "post", PostSchema{Post: post})
}

func GetPostList(c echo.Context) error {
	postList := loadPostList("public/markdown/posts/")
	return c.Render(http.StatusOK, "list", PostListSchema{List: postList})
}

func loadPost(dir, slug string) Post {
	f := markdown.MDFile{}
	f.Load(filepath.Join(dir, fmt.Sprintf("%s.md", slug)))

	p := Post{}
	p.WithMarkdown(f)

	return p
}

func loadPostList(dir string) []Post {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	ps := []Post{}
	for _, slug := range dirEntries {
		f := markdown.MDFile{}
		f.Load(filepath.Join(dir, slug.Name()))

		p := Post{}
		p.WithMarkdown(f)
		ps = append(ps, p)
	}
	return ps
}
