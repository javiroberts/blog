package posts

import "html/template"

type PostSchema struct {
	Article template.HTML
}

type PostListSchema struct {
	List []Post
}
