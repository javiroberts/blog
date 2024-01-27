package posts

import "time"

type Post struct {
	Author      string
	Title       string
	Excerpt     string
	Article     string
	Slug        string
	Published   time.Time
	LastUpdated time.Time
}
