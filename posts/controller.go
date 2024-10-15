package posts

import (
	"blog/markdown"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/labstack/echo/v4"
)

func GetPost(c echo.Context) error {
	dir := "public/markdown/posts/"
	slug := c.Param("slug")
	err, post := loadPost(dir, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.Render(http.StatusOK, "post", PostSchema{Post: post})
}

func GetPostList(c echo.Context) error {
	err, postList := loadPostList("public/markdown/posts/")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.Render(http.StatusOK, "list", PostListSchema{List: postList})
}

// TODO: refactor for better error handling
func loadPost(dir, slug string) (error, Post) {
	f := markdown.MDFile{}
	err := f.Load(filepath.Join(dir, fmt.Sprintf("%s.md", slug)))
	if err != nil {
		return err, Post{}
	}

	p := Post{}
	p.WithMarkdown(f)

	return nil, p
}

// TODO: refactor for better error handling
func loadPostList(dir string) (error, []Post) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return err, []Post{}
	}

	ps := []Post{}
	for _, slug := range dirEntries {
		f := markdown.MDFile{}
		f.Load(filepath.Join(dir, slug.Name()))

		p := Post{}
		p.WithMarkdown(f)
		ps = append(ps, p)
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Date.After(ps[j].Date)
	})

	return nil, ps
}
