package posts

import (
	"backend/markdown"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

func GetByID(c echo.Context) error {
	slug := c.Param("slug")
	return c.Render(http.StatusOK, "post", PostSchema{Article: markdown.LoadMarkdownPost(fmt.Sprintf("posts/%s", slug))})
}

func GetAboutPage(c echo.Context) error {
	return c.Render(http.StatusOK, "post", PostSchema{Article: markdown.LoadMarkdownPost("about")})
}

func GetPostList(c echo.Context) error {
	testList := buildPostList("public/markdown/posts/")
	return c.Render(http.StatusOK, "list", PostListSchema{List: testList})
}

func buildPostList(dir string) []Post {
	fmt.Println(os.ReadDir(dir))

	return []Post{
		{
			Author:      "Javier Roberts",
			Title:       "Golang in 2024",
			Excerpt:     "The panorama to come for go in 2024",
			Article:     "",
			Slug:        "deploying-an-aks-cluster-using-terraform",
			Published:   time.Now(),
			LastUpdated: time.Now(),
		},
		{
			Author:      "Javier Roberts",
			Title:       "Future JS",
			Excerpt:     "How Javascript is supposed to evolve in the near future",
			Article:     "",
			Slug:        "lessons-learned-from-scaling-a-startup",
			Published:   time.Now(),
			LastUpdated: time.Now(),
		},
	}
}
